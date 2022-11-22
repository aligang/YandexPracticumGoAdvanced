package handler

import (
	"github.com/aligang/YandexPracticumGoAdvanced/lib/config"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/storage"
)

func ExampleApiHandler() {
	conf := config.NewServer()
	config.GetServerCLIConfig(conf)
	config.GetServerENVConfig(conf)
	Storage, Type := storage.New(conf)
	mux := New(Storage, conf.Key, Type)
	mux.Post("/update/{metricType}/{metricName}/{metricValue}", mux.Update)

}
