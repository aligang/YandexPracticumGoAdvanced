package config

import (
	"encoding/json"
	"os"
)

func getServerConfigFromFile(filePath string) *ServerConfig {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0400)
	configFromFile := &ServerConfig{}
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(configFromFile)
	if err != nil {
		panic(err)
	}
	return configFromFile
}

func getAgentConfigFromFile(filePath string) *AgentConfig {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0400)
	configFromFile := &AgentConfig{}
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(configFromFile)
	if err != nil {
		panic(err)
	}
	return configFromFile
}
