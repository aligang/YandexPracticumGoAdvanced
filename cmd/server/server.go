package main

import (
	"github.com/aligang/YandexPracticumGoAdvanced/lib/encrypt"
	"log"
	"net/http"
	"os"

	"github.com/aligang/YandexPracticumGoAdvanced/lib/compress"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/config"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/handler"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/logging"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/storage"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
)

func main() {
	printBuildInfo()
	conf := config.NewServer()
	config.GetServerCLIConfig(conf)
	config.GetServerENVConfig(conf)
	logging.Configure(os.Stdout, zerolog.DebugLevel)
	logging.Debug("Starting Server with config : %+v\n", *conf)
	Storage, Type := storage.New(conf)

	encryption := encrypt.GetServerPlugin(conf)
	mux := handler.New(Storage, conf.Key, Type)
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Recoverer)

	mux.Post("/update/{metricType}/{metricName}/{metricValue}", mux.Update)
	mux.With(compress.GzipHandle, encryption.DecryptWithPublicKey).Post("/update/", mux.UpdateWithJSON)
	mux.With(compress.GzipHandle, encryption.DecryptWithPublicKey).Post("/updates/", mux.BulkUpdate)

	mux.With(compress.GzipHandle).Get("/", mux.FetchAll)
	mux.Get("/value/{metricType}/{metricName}", mux.Fetch)
	mux.With(compress.GzipHandle, encryption.DecryptWithPublicKey).Post("/value/", mux.FetchWithJSON)

	mux.Get("/ping", mux.Ping)

	log.Fatal(http.ListenAndServe(conf.Address, mux))

}
