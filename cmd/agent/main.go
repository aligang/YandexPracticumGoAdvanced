package main

import (
	"fmt"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/collector"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/config"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/reporter"
	"github.com/caarlos0/env/v6"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var wg sync.WaitGroup

func main() {
	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)

	var conf config.AgentConfig
	err := env.Parse(&conf)
	if err != nil {
		fmt.Println("Could not fetch server ENV params")
		panic(err)
	}

	stats := &metric.Stats{
		map[string]float64{},
		map[string]int64{},
	}
	
	go collector.CollectMetrics(conf.PollInterval, stats)
	go reporter.SendMetrics(conf.ReportInterval, stats)

	<-exitSignal

}
