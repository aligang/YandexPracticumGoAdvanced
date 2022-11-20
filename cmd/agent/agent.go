package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/aligang/YandexPracticumGoAdvanced/iinternal/collector"
	"github.com/aligang/YandexPracticumGoAdvanced/iinternal/config"
	"github.com/aligang/YandexPracticumGoAdvanced/iinternal/logging"
	"github.com/aligang/YandexPracticumGoAdvanced/iinternal/metric"
	"github.com/aligang/YandexPracticumGoAdvanced/iinternal/reporter"
	"github.com/rs/zerolog"
)

func main() {
	conf := config.NewAgent()
	config.GetAgentCLIConfig(conf)
	config.GetAgentENVConfig(conf)
	logging.Configure(os.Stdout, zerolog.DebugLevel)
	logging.Logger.Printf("Starting Agent with config: %+v", conf)
	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)

	bus := make(chan metric.Stats)
	go collector.CollectMetrics(conf, bus)
	go reporter.SendMetrics(conf, bus)
	go reporter.BulkSendMetrics(conf, bus)

	<-exitSignal

}
