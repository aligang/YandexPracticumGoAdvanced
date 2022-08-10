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

	stats := &metric.Stats{}

	go collector.CollectMetrics(stats)
	go reporter.SendMetrics(stats)

	<-exitSignal

}
