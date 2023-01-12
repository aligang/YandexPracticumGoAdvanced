package config

import (
	"flag"
	"time"
)

// GetServerCLIConfig enriches app configuration with CLI defined params
//func GetServerCLIConfig(conf *ServerConfig) {
//	flag.StringVar(&conf.Address, "a", "127.0.0.1:8080", "host to listen on")
//	flag.DurationVar(&conf.StoreInterval, "i", 300*time.Second, "period to backup data")
//	flag.StringVar(&conf.StoreFile, "f", "/tmp/devops-metrics-db.json", "backup filepath")
//	flag.BoolVar(&conf.Restore, "r", true, "Read from backup before Startup")
//	flag.StringVar(&conf.Key, "k", "", "Hashing key")
//	flag.StringVar(&conf.DatabaseDsn, "d", "", "Database")
//	flag.StringVar(&conf.CryptoKey, "crypto-key", "", "CryptoKey")
//	flag.StringVar(&conf.Config, "c", "", "Config")
//	flag.Parse()
//}
//
//// GetAgentCLIConfig enriches agent configuration with CLI defined params
//func GetAgentCLIConfig(conf *AgentConfig) {
//	flag.StringVar(&conf.Address, "a", "127.0.0.1:8080", "host to listen on")
//	flag.DurationVar(&conf.PollInterval, "p", 2*time.Second, "period for collection metrics by agent")
//	flag.DurationVar(&conf.ReportInterval, "r", 10*time.Second, "period for pushing metrics by agent")
//	flag.StringVar(&conf.Key, "k", "", "Hashing key")
//	flag.StringVar(&conf.CryptoKey, "crypto-key", "", "CryptoKey")
//	flag.StringVar(&conf.Config, "c", "", "Config")
//	flag.Parse()
//}

func getServerCLIConfig() *ServerConfig {
	conf := &ServerConfig{}
	flag.StringVar(&conf.Address, "a", "", "host to listen on")
	flag.DurationVar(&conf.StoreInterval, "i", -1*time.Second, "period to backup data")
	flag.StringVar(&conf.StoreFile, "f", "", "backup filepath")
	flag.BoolVar(&conf.Restore, "r", false, "Read from backup before Startup")
	flag.StringVar(&conf.Key, "k", "", "Hashing key")
	flag.StringVar(&conf.DatabaseDsn, "d", "", "Database")
	flag.StringVar(&conf.CryptoKey, "crypto-key", "", "CryptoKey")
	flag.StringVar(&conf.Config, "c", "", "Config")
	flag.Parse()
	return conf
}

// GetAgentCLIConfig enriches agent configuration with CLI defined params
func getAgentCLIConfig() *AgentConfig {
	conf := &AgentConfig{}
	flag.StringVar(&conf.Address, "a", "", "host to listen on")
	flag.DurationVar(&conf.PollInterval, "p", -1*time.Second, "period for collection metrics by agent")
	flag.DurationVar(&conf.ReportInterval, "r", -1*time.Second, "period for pushing metrics by agent")
	flag.StringVar(&conf.Key, "k", "", "Hashing key")
	flag.StringVar(&conf.CryptoKey, "crypto-key", "", "CryptoKey")
	flag.StringVar(&conf.Config, "c", "", "Config")
	flag.Parse()
	return conf
}
