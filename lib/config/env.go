package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

// GetServerENVConfig enriches app configuration with ENV defined params
func getServerENVConfig() *ServerConfig {
	envConf := &ServerConfig{}
	err := env.Parse(envConf)
	if err != nil {
		fmt.Println("Could not fetch app ENV params")
		panic(err)
	}
	//if envConf.Address != "" {
	//	conf.Address = envConf.Address
	//}
	//if envConf.StoreFile != "" {
	//	conf.StoreFile = envConf.StoreFile
	//}
	//if envConf.Key != "" {
	//	conf.Key = envConf.Key
	//}
	//if envConf.DatabaseDsn != "" {
	//	conf.DatabaseDsn = envConf.DatabaseDsn
	//}
	//if envConf.CryptoKey != "" {
	//	conf.CryptoKey = envConf.CryptoKey
	//}
	//
	//if envConf.StoreInterval >= 0 {
	//	conf.StoreInterval = envConf.StoreInterval
	//}
	//conf.Restore = conf.Restore || envConf.Restore
	return envConf
}

// GetAgentENVConfig enriches agent configuration with ENV defined params
func getAgentENVConfig() *AgentConfig {
	envConf := &AgentConfig{}
	err := env.Parse(envConf)
	if err != nil {
		fmt.Println("Could not fetch Agent ENV params")
		panic(err)
	}
	//if envConf.Address != "" {
	//	conf.Address = envConf.Address
	//}
	//if envConf.Key != "" {
	//	conf.Key = envConf.Key
	//}
	//if envConf.CryptoKey != "" {
	//	conf.CryptoKey = envConf.CryptoKey
	//}
	//
	//if envConf.PollInterval >= 0 {
	//	conf.PollInterval = envConf.PollInterval
	//}
	//if envConf.ReportInterval >= 0 {
	//	conf.ReportInterval = envConf.ReportInterval
	//}
	return envConf
}
