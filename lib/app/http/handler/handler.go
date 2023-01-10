package handler

import (
	"github.com/aligang/YandexPracticumGoAdvanced/lib/app"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/storage"
	"github.com/go-chi/chi/v5"
)

type APIHandler struct {
	*chi.Mux
	Storage storage.Storage
	Config  app.Config
}

func New(s storage.Storage, h string, t string) APIHandler {
	mux := APIHandler{
		Mux:     chi.NewMux(),
		Storage: s,
		Config: app.Config{
			HashKey:     h,
			StorageType: t,
		},
	}
	return mux
}
