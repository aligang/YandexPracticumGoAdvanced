package config

import (
	"time"
)

// ServerConfig represents configuration of server
type ServerConfig struct {
	Address       string        `env:"ADDRESS" envDefault:"" json:"address"`
	StoreInterval time.Duration `env:"STORE_INTERVAL" envDefault:"-1s" json:"store_interval"`
	StoreFile     string        `env:"STORE_FILE" envDefault:"" json:"store_file"`
	Restore       bool          `env:"RESTORE" envDefault:"false" json:"restore"`
	Key           string        `env:"KEY" envDefault:"" json:"key"`
	DatabaseDsn   string        `env:"DATABASE_DSN" envDefault:"" json:"database_dsn"`
	CryptoKey     string        `env:"CRYPTO_KEY" envDefault:"" json:"crypto_key"`
	Config        string        `env:"CONFIG" envDefault:""`
}

func (s *ServerConfig) Merge(a ...*ServerConfig) *ServerConfig {
	for _, another := range a {
		if s.Address == "" && another.Address != "" {
			s.Address = another.Address
		}
		if s.StoreFile == "" && another.StoreFile != "" {
			s.StoreFile = another.StoreFile
		}
		if s.Key == "" && another.Key != "" {
			s.Key = another.Key
		}
		if s.DatabaseDsn == "" && another.DatabaseDsn != "" {
			s.DatabaseDsn = another.DatabaseDsn
		}
		if s.CryptoKey == "" && another.CryptoKey != "" {
			s.CryptoKey = another.CryptoKey
		}

		if s.StoreInterval <= 0 && another.StoreInterval > 0 {
			s.StoreInterval = another.StoreInterval
		}
		if !s.Restore && another.Restore {
			s.Restore = true
		}
	}
	return s
}

// AgentConfig represents configuration of agent
type AgentConfig struct {
	Address        string        `env:"ADDRESS" envDefault:"" json:"address"`
	PollInterval   time.Duration `env:"POLL_INTERVAL" envDefault:"-1s" json:"poll_interval"`
	ReportInterval time.Duration `env:"REPORT_INTERVAL" envDefault:"-1s" json:"report_interval"`
	Key            string        `env:"KEY" envDefault:"" json:"key"`
	CryptoKey      string        `env:"CRYPTO_KEY" envDefault:"" json:"crypto_key"`
	Config         string        `env:"CONFIG" envDefault:""`
}

func (s *AgentConfig) Merge(a ...*AgentConfig) *AgentConfig {
	for _, another := range a {
		if s.Address == "" && another.Address != "" {
			s.Address = another.Address
		}
		if s.Key == "" && another.Key != "" {
			s.Key = another.Key
		}
		if s.CryptoKey == "" && another.CryptoKey != "" {
			s.CryptoKey = another.CryptoKey
		}
		if s.PollInterval <= 0 && another.PollInterval > 0 {
			s.PollInterval = another.PollInterval
		}
		if s.ReportInterval <= 0 && another.ReportInterval > 0 {
			s.ReportInterval = another.ReportInterval
		}
	}
	return s
}

func GetServerConfig() *ServerConfig {
	envCfg := getServerENVConfig()
	cliCfg := getServerCLIConfig()
	fileCfg := newServerConfig()

	switch {
	case envCfg.Config != "":
		fileCfg = getServerConfigFromFile(envCfg.Config)
	case cliCfg.Config != "":
		fileCfg = getServerConfigFromFile(cliCfg.Config)
	}
	return envCfg.Merge(
		cliCfg,
		fileCfg,
		getDefaultServerConfig(),
	)
}

func GetAgentConfig() *AgentConfig {
	envCfg := getAgentENVConfig()
	cliCfg := getAgentCLIConfig()
	fileCfg := newAgentConfig()

	switch {
	case envCfg.Config != "":
		fileCfg = getAgentConfigFromFile(envCfg.Config)
	case cliCfg.Config != "":
		fileCfg = getAgentConfigFromFile(cliCfg.Config)
	}
	return envCfg.Merge(
		cliCfg,
		fileCfg,
		getDefaultAgentConfig(),
	)
}
