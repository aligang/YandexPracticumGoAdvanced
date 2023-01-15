package main

import (
	"context"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/agent"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/config"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/logging"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/metric"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	printBuildInfo()
	conf := config.GetAgentConfig()
	logging.Configure(os.Stdout, zerolog.DebugLevel)
	logging.Debug("Starting Agent with config: %+v", conf)

	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	ctx := context.Background()
	defer ctx.Done()
	ctx, cancel := context.WithCancel(ctx)
	conn, err := grpc.Dial(conf.GrpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	a := agent.New(conf, conn)

	inBus := make(chan metric.Stats)
	unaryJSONBus := make(chan metric.Stats)
	bulkJSONBus := make(chan metric.Stats)
	unaryGRPCBus := make(chan metric.Stats)
	bulkGRPCBus := make(chan metric.Stats)

	stopReporters := make(chan interface{})
	unaryJSONExit := make(chan interface{})
	bulkJSONExit := make(chan interface{})
	unaryGRPCExit := make(chan interface{})
	bulkGRPCExit := make(chan interface{})

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		var m metric.Stats
	loop:
		for {
			select {
			case m = <-inBus:
				unaryJSONBus <- m
				bulkJSONBus <- m
				unaryGRPCBus <- m
				bulkGRPCBus <- m
			case <-stopReporters:
				unaryJSONExit <- struct{}{}
				bulkJSONExit <- struct{}{}
				unaryGRPCExit <- struct{}{}
				bulkGRPCExit <- struct{}{}
				close(inBus)
				break loop
			}
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		logging.Debug("Starting metrics collector")
		a.CollectMetrics(ctx, conf, inBus, stopReporters)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		logging.Debug("Starting unary HTTP/JSON reporter")
		a.SendMetrics(conf, unaryJSONBus, unaryJSONExit)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		logging.Debug("Starting bulk HTTP/JSON reporter")
		a.BulkSendMetrics(conf, bulkJSONBus, bulkJSONExit)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		logging.Debug("Starting unary GRPC reporter")
		a.SendGrpcMetrics(conf, unaryGRPCBus, unaryGRPCExit)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		logging.Debug("Starting bulk GRPC reporter")
		a.BulkSendGrpcMetrics(conf, bulkGRPCBus, bulkGRPCExit)
		wg.Done()
	}()

	<-exitSignal
	cancel()
	wg.Wait()
	os.Exit(0)
}
