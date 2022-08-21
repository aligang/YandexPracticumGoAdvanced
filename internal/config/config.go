package config

type ServerConfig struct {
	Address string `env:"ADDRESS" envDefault:"127.0.0.1:8080"`
}

type AgentConfig struct {
	PollInterval   int `env:"POLL_INTERVAL" envDefault:"2"`
	ReportInterval int `env:"REPORT_INTERVAL" envDefault:"10"`
}
