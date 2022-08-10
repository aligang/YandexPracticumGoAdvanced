package main

import (
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
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)

	mux.Post("/update/{metricType}/{metricName}/{metricValue}", mux.Update)

	mux.Get("/", mux.FetchAll)
	mux.Get("/value/{metricType}/{metricName}", mux.Fetch)

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", mux))

}
