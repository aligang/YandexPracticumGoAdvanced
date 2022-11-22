package config

func ExampleServerConfig() {
	conf := NewServer()
	GetServerCLIConfig(conf)
	GetServerENVConfig(conf)
}

func ExampleAgentConfig() {
	conf := NewAgent()
	GetAgentCLIConfig(conf)
	GetAgentENVConfig(conf)
}
