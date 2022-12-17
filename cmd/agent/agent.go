package main

import (
	"github.com/aligang/YandexPracticumGoAdvanced/lib/agent"
	"os"
	"os/signal"
	"syscall"

	"github.com/aligang/YandexPracticumGoAdvanced/lib/config"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/logging"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/metric"
	"github.com/rs/zerolog"
)

func main() {
	conf := config.NewAgent()
	config.GetAgentCLIConfig(conf)
	config.GetAgentENVConfig(conf)
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

}
