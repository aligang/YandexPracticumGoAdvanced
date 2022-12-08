package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/aligang/YandexPracticumGoAdvanced/lib/collector"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/config"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/logging"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/metric"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/reporter"
	"github.com/rs/zerolog"
)

func main() {
	printBuildInfo()
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
	os.Exit(0)
}
