package handler

import (
	"github.com/aligang/YandexPracticumGoAdvanced/lib/config"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/storage"
)

func ExampleHTTPHandler() {
	conf := config.GetServerConfig()
	Storage, Type := storage.New(conf)
	mux := New(Storage, conf.Key, Type)
	mux.Post("/update/{metricType}/{metricName}/{metricValue}", mux.Update)

}
