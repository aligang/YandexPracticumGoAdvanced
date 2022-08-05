package main

import (
	"github.com/aligang/YandexPracticumGoAdvanced/internal/collector"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/reporter"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	stats := &metric.Stats{
		MemStats:  &runtime.MemStats{},
		OperStats: &metric.OperStats{},
	}

	go collector.CollectMetrics(stats)
	go reporter.SendMetrics(stats)
	for {
		time.Sleep(time.Second * 10)
	}

}
