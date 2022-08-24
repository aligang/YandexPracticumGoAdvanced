package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
)

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

	if envConf.StoreInterval >= 0 {
		conf.StoreInterval = envConf.StoreInterval
	}
	conf.Restore = conf.Restore || envConf.Restore

}

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
	if envConf.PollInterval >= 0 {
		conf.PollInterval = envConf.PollInterval
	}
	if envConf.ReportInterval >= 0 {
		conf.ReportInterval = envConf.ReportInterval
	}
}
