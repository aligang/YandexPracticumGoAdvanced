package main

import (
	"github.com/aligang/YandexPracticumGoAdvanced/internal/compress"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/config"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/handler"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/storage"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func main() {
	conf := config.NewServer()
	config.GetServerCLIConfig(conf)
	config.GetServerENVConfig(conf)
	Storage := storage.New()
	if conf.Restore {
		Storage.Restore(conf)
	}
	Storage.ConfigureBackup(conf)
	mux := handler.New(Storage)
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Recoverer)

	mux.Post("/update/{metricType}/{metricName}/{metricValue}", mux.Update)
	mux.Post("/update/", compress.GzipHandle(mux.UpdateWithJSON))

	mux.Get("/", compress.GzipHandle(mux.FetchAll))
	mux.Get("/value/{metricType}/{metricName}", mux.Fetch)
	mux.Post("/value/", compress.GzipHandle(mux.FetchWithJSON))

	log.Fatal(http.ListenAndServe(conf.Address, mux))

}
