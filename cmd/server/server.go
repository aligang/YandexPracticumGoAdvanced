package main

import (
	"github.com/aligang/YandexPracticumGoAdvanced/internal/config"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/handler"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/storage"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func main() {
	Storage := storage.New()
	mux := handler.New(Storage)
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	//mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)

	mux.Post("/update/{metricType}/{metricName}/{metricValue}", mux.Update)
	mux.Post("/update/", mux.UpdateWithJson)

	mux.Get("/", mux.FetchAll)
	mux.Get("/value/{metricType}/{metricName}", mux.Fetch)
	mux.Post("/value/", mux.FetchWithJson)

	conf := config.GetServerConfig()
	log.Fatal(http.ListenAndServe(conf.Address, mux))

}
