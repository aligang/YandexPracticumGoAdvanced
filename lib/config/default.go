package config

import "time"

func getDefaultServerConfig() *ServerConfig {
	conf := newServerConfig()
	conf.Address = "127.0.0.1:8080"
	conf.StoreInterval = 300 * time.Second
	conf.StoreFile = "/tmp/devops-metrics-db.json"
	conf.Restore = true
	conf.Key = ""
	conf.DatabaseDsn = ""
	conf.CryptoKey = ""
	return conf
}

func getDefaultAgentConfig() *AgentConfig {
	conf := newAgentConfig()
	conf.Address = "127.0.0.1:8080"
	conf.Key = ""
	conf.CryptoKey = ""
	conf.PollInterval = 2 * time.Second
	conf.ReportInterval = 10 * time.Second
	return conf
}
