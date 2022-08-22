package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
)

type ServerConfig struct {
	Address string `env:"ADDRESS" envDefault:"127.0.0.1:8080"`
}

type AgentConfig struct {
	PollInterval   int `env:"POLL_INTERVAL" envDefault:"2"`
	ReportInterval int `env:"REPORT_INTERVAL" envDefault:"10"`
}

func GetServerConfig() ServerConfig {
	var conf ServerConfig
	err := env.Parse(&conf)
	if err != nil {
		fmt.Println("Could not fetch server ENV params")
		panic(err)
	}
	return conf
}

func GetAgentConfig() AgentConfig {
	var conf AgentConfig
	err := env.Parse(&conf)
	if err != nil {
		fmt.Println("Could not fetch agent ENV params")
		panic(err)
	}
	return conf
}
