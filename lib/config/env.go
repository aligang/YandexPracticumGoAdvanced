package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

// GetServerENVConfig enriches server configuration with ENV defined params
func GetServerENVConfig(conf *ServerConfig) {
	envConf := ServerConfig{}
	err := env.Parse(&envConf)
	if err != nil {
		fmt.Println("Could not fetch server ENV params")
		panic(err)
	}
	if envConf.Address != "" {
		conf.Address = envConf.Address
	}
	if envConf.StoreFile != "" {
		conf.StoreFile = envConf.StoreFile
	}
	if envConf.Key != "" {
		conf.Key = envConf.Key
	}
	if envConf.DatabaseDsn != "" {
		conf.DatabaseDsn = envConf.DatabaseDsn
	}

	if envConf.StoreInterval >= 0 {
		conf.StoreInterval = envConf.StoreInterval
	}
	conf.Restore = conf.Restore || envConf.Restore

}

// GetAgentENVConfig enriches agent configuration with ENV defined params
func GetAgentENVConfig(conf *AgentConfig) {
	envConf := AgentConfig{}
	err := env.Parse(&envConf)
	if err != nil {
		fmt.Println("Could not fetch Agent ENV params")
		panic(err)
	}
	if envConf.Address != "" {
		conf.Address = envConf.Address
	}
	if envConf.Key != "" {
		conf.Key = envConf.Key
	}

	if envConf.PollInterval >= 0 {
		conf.PollInterval = envConf.PollInterval
	}
	if envConf.ReportInterval >= 0 {
		conf.ReportInterval = envConf.ReportInterval
	}
}