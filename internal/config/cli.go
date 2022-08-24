package config

import (
	"flag"
	"time"
)

func GetServerCLIConfig(conf *ServerConfig) {
	flag.StringVar(&conf.Address, "a", "127.0.0.1:8080", "host to listen on")
	flag.DurationVar(&conf.StoreInterval, "i", 300*time.Second, "period to backup data")
	flag.StringVar(&conf.StoreFile, "f", "/tmp/devops-metrics-db.json", "backup filepath")
	flag.BoolVar(&conf.Restore, "r", true, "Read from backup before Startup")
	flag.Parse()
}

func GetAgentCLIConfig(conf *AgentConfig) {
	flag.StringVar(&conf.Address, "a", "127.0.0.1:8080", "host to listen on")
	flag.DurationVar(&conf.PollInterval, "i", 2*time.Second, "period for collection metrics by agent")
	flag.DurationVar(&conf.ReportInterval, "i", 10*time.Second, "period for pushing metrics by agent")
	flag.Parse()
}
