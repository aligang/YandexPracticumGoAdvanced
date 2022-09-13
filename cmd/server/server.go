package main

import (
	"github.com/aligang/YandexPracticumGoAdvanced/internal/compress"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/config"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/handler"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/logging"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/storage"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
	"log"
	"net/http"
	"os"
)

func main() {
	conf := config.NewServer()
	config.GetServerCLIConfig(conf)
	config.GetServerENVConfig(conf)
	logging.Configure(os.Stdout, zerolog.DebugLevel)
	logging.Debug("Starting Server with config : %+v\n", *conf)
	Storage, Type := storage.New(conf)
	mux := handler.New(Storage, conf.Key, Type)
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Recoverer)

	mux.Post("/update/{metricType}/{metricName}/{metricValue}", mux.Update)
	mux.Post("/update/", compress.GzipHandle(mux.UpdateWithJSON))
	mux.Post("/updates/", compress.GzipHandle(mux.BulkUpdate))

	mux.Get("/", compress.GzipHandle(mux.FetchAll))
	mux.Get("/value/{metricType}/{metricName}", mux.Fetch)
	mux.Post("/value/", compress.GzipHandle(mux.FetchWithJSON))

	mux.Get("/ping", mux.Ping)

	log.Fatal(http.ListenAndServe(conf.Address, mux))

}
