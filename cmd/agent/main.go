package main

import (
	"github.com/aligang/YandexPracticumGoAdvanced/internal/collector"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/reporter"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var wg sync.WaitGroup

func main() {
	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)

	stats := &metric.Stats{
		map[string]float64{},
		map[string]int64{},
	}
	//stats := map[string]metric.Metrics{}

	go collector.CollectMetrics(stats)
	go reporter.SendMetrics(stats)

	<-exitSignal

}
