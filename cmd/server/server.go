package main

import (
	"context"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/accesslist"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/app/grpc/generated/service"
	grpcHandler "github.com/aligang/YandexPracticumGoAdvanced/lib/app/grpc/handler"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/app/http/handler"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/compress"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/config"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/encrypt"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/logging"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/storage"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	printBuildInfo()
	conf := config.GetServerConfig()
	logging.Configure(os.Stdout, zerolog.DebugLevel)
	logging.Debug("Starting Server with config : %+v\n", *conf)
	Storage, Type := storage.New(conf)

	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	httpShutdownSignal := make(chan any, 1)
	grpcShutdownSignal := make(chan any, 1)

	wg := sync.WaitGroup{}
	wg.Add(1)

	//Exit signal fan out
	go func() {
		<-exitSignal
		httpShutdownSignal <- struct{}{}
		grpcShutdownSignal <- struct{}{}
		close(exitSignal)
		wg.Done()
	}()

	//HTTP Server
	wg.Add(1)
	go func() {
		encryption := encrypt.GetServerPlugin(conf)
		mux := handler.New(Storage, conf.Key, Type)
		mux.Use(middleware.RequestID)
		mux.Use(middleware.RealIP)
		mux.Use(middleware.Recoverer)

		mux.With(
			accesslist.IPValidationInterceptor(conf.TrustedSubnet),
		).Post("/update/{metricType}/{metricName}/{metricValue}", mux.Update)
		mux.With(
			accesslist.IPValidationInterceptor(conf.TrustedSubnet),
			compress.GzipHandle,
			encryption.DecryptWithPrivateKey,
		).Post("/update/", mux.UpdateWithJSON)
		mux.With(
			accesslist.IPValidationInterceptor(conf.TrustedSubnet),
			compress.GzipHandle,
			encryption.DecryptWithPrivateKey,
		).Post("/updates/", mux.BulkUpdate)

		mux.With(compress.GzipHandle).Get("/", mux.FetchAll)
		mux.Get("/value/{metricType}/{metricName}", mux.Fetch)
		mux.With(compress.GzipHandle, encryption.DecryptWithPrivateKey).Post("/value/", mux.FetchWithJSON)
		mux.Get("/ping", mux.Ping)

		srv := http.Server{Addr: conf.Address, Handler: mux}

		wg.Add(1)
		go func() {
			<-httpShutdownSignal
			logging.Debug("HTTP Server Stops...")
			if err := srv.Shutdown(context.Background()); err != nil {
				log.Printf("HTTP Server Shutdown: %v", err)
			}
			close(httpShutdownSignal)
			wg.Done()
		}()
		logging.Debug("HTTP Server Starts...")
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
		wg.Done()
	}()

	//GRPC Server
	wg.Add(1)
	go func() {
		listen, err := net.Listen("tcp", ":3200")
		if err != nil {
			log.Fatal(err)
		}
		s := grpc.NewServer()
		handler := grpcHandler.New(Storage, conf.Key, Type)
		service.RegisterMetricsServiceServer(s, handler)
		logging.Debug("GRPC Server Starts...")

		wg.Add(1)
		go func() {
			<-grpcShutdownSignal
			logging.Debug("GRPC Server Stops...")
			s.GracefulStop()
			close(grpcShutdownSignal)
			wg.Done()
		}()

		if err := s.Serve(listen); err != nil {
			log.Fatal(err)
		}
		wg.Done()
	}()

	wg.Wait()
}
