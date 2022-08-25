package main

import (
	"github.com/aligang/YandexPracticumGoAdvanced/internal/collector"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/config"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/reporter"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	conf := config.NewAgent()
	config.GetAgentCLIConfig(conf)
	config.GetAgentENVConfig(conf)
	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)

	stats := &metric.Stats{
		Gauge:   map[string]float64{},
		Counter: map[string]int64{},
	}

	go collector.CollectMetrics(conf, stats)
	go reporter.SendMetrics(conf, stats)

	<-exitSignal

}
