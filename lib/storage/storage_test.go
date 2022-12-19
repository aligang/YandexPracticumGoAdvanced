package storage

import "github.com/aligang/YandexPracticumGoAdvanced/lib/config"

func ExampleStorage() {
	//import "github.com/aligang/YandexPracticumGoAdvanced/lib/config"

	conf := config.GetServerConfig()
	New(conf)
}
