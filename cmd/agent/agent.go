package main

import (
	"context"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/agent"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/config"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/logging"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/metric"
	"github.com/rs/zerolog"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	printBuildInfo()
	conf := config.GetAgentConfig()
	logging.Configure(os.Stdout, zerolog.DebugLevel)
	logging.Logger.Printf("Starting Agent with config: %+v", conf)
	a := agent.New(conf)
	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	ctx := context.Background()
	defer ctx.Done()
	ctx, cancel := context.WithCancel(ctx)
	bus := make(chan metric.Stats)
	stopWorkers := make(chan interface{})
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		a.CollectMetrics(ctx, conf, bus, stopWorkers)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		a.SendMetrics(conf, bus, stopWorkers)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		a.BulkSendMetrics(conf, bus, stopWorkers)
		wg.Done()
	}()

	<-exitSignal
	cancel()
	wg.Wait()
	os.Exit(0)
}
