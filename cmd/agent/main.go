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

	var agentConfig config.AgentConfig
	var serverConfig config.ServerConfig
	aerr := env.Parse(&agentConfig)
	serr := env.Parse(&serverConfig)
	if aerr != nil || serr != nil {
		panic("Could not fetch ENV params")
	}
	fmt.Printf("%+v\n", agentConfig)
	fmt.Printf("%+v\n", serverConfig)

	stats := &metric.Stats{
		map[string]float64{},
		map[string]int64{},
	}

	go collector.CollectMetrics(agentConfig.PollInterval, stats)
	go reporter.SendMetrics(serverConfig.Address, agentConfig.ReportInterval, stats)

	<-exitSignal

}
