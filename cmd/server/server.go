package main

import (
	"context"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/compress"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/config"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/encrypt"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/handler"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/logging"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/storage"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	printBuildInfo()
	conf := config.GetServerConfig()
	logging.Configure(os.Stdout, zerolog.DebugLevel)
	logging.Debug("Starting Server with config : %+v\n", *conf)
	Storage, Type := storage.New(conf)

	encryption := encrypt.GetServerPlugin(conf)
	mux := handler.New(Storage, conf.Key, Type)
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Recoverer)

	mux.Post("/update/{metricType}/{metricName}/{metricValue}", mux.Update)
	mux.With(compress.GzipHandle, encryption.DecryptWithPrivateKey).Post("/update/", mux.UpdateWithJSON)
	mux.With(compress.GzipHandle, encryption.DecryptWithPrivateKey).Post("/updates/", mux.BulkUpdate)

	mux.With(compress.GzipHandle).Get("/", mux.FetchAll)
	mux.Get("/value/{metricType}/{metricName}", mux.Fetch)
	mux.With(compress.GzipHandle, encryption.DecryptWithPrivateKey).Post("/value/", mux.FetchWithJSON)
	mux.Get("/ping", mux.Ping)

	srv := http.Server{Addr: conf.Address, Handler: mux}

	idleConnsClosed := make(chan struct{})
	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-exitSignal
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
