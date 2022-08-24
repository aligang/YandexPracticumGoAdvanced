package config

import (
	"time"
)

type ServerConfig struct {
	Address       string        `env:"ADDRESS" envDefault:""`
	StoreInterval time.Duration `env:"STORE_INTERVAL" envDefault:"-1s"`
	StoreFile     string        `env:"STORE_FILE" envDefault:""`
	Restore       bool          `env:"RESTORE" envDefault:"false"`
}

type AgentConfig struct {
	Address        string        `env:"ADDRESS" envDefault:""`
	PollInterval   time.Duration `env:"POLL_INTERVAL" envDefault:"-1s"`
	ReportInterval time.Duration `env:"REPORT_INTERVAL" envDefault:"-1s"`
}

func NewServer() *ServerConfig {
	return &ServerConfig{}
}

func NewAgent() *ServerConfig {
	return &ServerConfig{}
}
