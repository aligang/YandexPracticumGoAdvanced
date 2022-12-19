package main

import (
	"github.com/aligang/YandexPracticumGoAdvanced/lib/agent"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/config"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/logging"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/metric"
	"github.com/rs/zerolog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	printBuildInfo()
	conf := config.GetAgentConfig()
	logging.Configure(os.Stdout, zerolog.DebugLevel)
	logging.Logger.Printf("Starting Agent with config: %+v", conf)
	a := agent.New(conf)
	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)

	bus := make(chan metric.Stats)
	go a.CollectMetrics(conf, bus)
	go a.SendMetrics(conf, bus)
	go a.BulkSendMetrics(conf, bus)

	<-exitSignal
	os.Exit(0)
}
