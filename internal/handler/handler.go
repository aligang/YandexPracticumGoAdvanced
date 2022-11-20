package handler

import (
	"github.com/aligang/YandexPracticumGoAdvanced/internal/storage"
	"github.com/go-chi/chi/v5"
)

type Config struct {
	HashKey     string
	StorageType string
}

type APIHandler struct {
	*chi.Mux
	Storage storage.Storage
	Config  Config
}

func New(s storage.Storage, h string, t string) APIHandler {
	mux := APIHandler{
		Mux:     chi.NewMux(),
		Storage: s,
		Config: Config{
			HashKey:     h,
			StorageType: t,
		},
	}
	return mux
}
