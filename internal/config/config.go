package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"time"
)

type ServerConfig struct {
	Address       string        `env:"ADDRESS" envDefault:"127.0.0.1:8080"`
	StoreInterval time.Duration `env:"STORE_INTERVAL" envDefault:"300s"`
	StoreFile     string        `env:"STORE_FILE" envDefault:"/tmp/devops-metrics-db.json"`
	Restore       bool          `env:"RESTORE" envDefault:"true"`
}

type AgentConfig struct {
	Address        string        `env:"ADDRESS" envDefault:"127.0.0.1:8080"`
	PollInterval   time.Duration `env:"POLL_INTERVAL" envDefault:"2s"`
	ReportInterval time.Duration `env:"REPORT_INTERVAL" envDefault:"10s"`
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
