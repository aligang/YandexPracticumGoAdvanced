package storage

import "github.com/aligang/YandexPracticumGoAdvanced/lib/config"

func ExampleStorage() {
	//import "github.com/aligang/YandexPracticumGoAdvanced/lib/config"

	conf := config.NewServer()
	config.GetServerCLIConfig(conf)
	config.GetServerENVConfig(conf)
	New(conf)
}
