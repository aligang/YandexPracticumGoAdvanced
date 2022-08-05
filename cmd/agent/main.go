package main

import (
	"github.com/aligang/YandexPracticumGoAdvanced/internal/collector"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/reporter"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	stats := &metric.Stats{}

	go collector.CollectMetrics(stats)
	go reporter.SendMetrics(stats)
	for {
		time.Sleep(time.Second * 10)
	}

}
