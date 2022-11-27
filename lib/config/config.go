package config

import (
	"time"
)

// ServerConfig represents configuration of server
type ServerConfig struct {
	Address       string        `env:"ADDRESS" envDefault:""`
	StoreInterval time.Duration `env:"STORE_INTERVAL" envDefault:"-1s"`
	StoreFile     string        `env:"STORE_FILE" envDefault:""`
	Restore       bool          `env:"RESTORE" envDefault:"false"`
	Key           string        `env:"KEY" envDefault:""`
	DatabaseDsn   string        `env:"DATABASE_DSN" envDefault:""`
}

// AgentConfig represents configuration of agent
type AgentConfig struct {
	Address        string        `env:"ADDRESS" envDefault:""`
	PollInterval   time.Duration `env:"POLL_INTERVAL" envDefault:"-1s"`
	ReportInterval time.Duration `env:"REPORT_INTERVAL" envDefault:"-1s"`
	Key            string        `env:"KEY" envDefault:""`
}

// NewServer returns blanc server configuration
func NewServer() *ServerConfig {
	return &ServerConfig{}
}

// NewAgent returns blanc server configuration
func NewAgent() *AgentConfig {
	return &AgentConfig{}
}
