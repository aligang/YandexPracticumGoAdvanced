package handler

import (
	"github.com/aligang/YandexPracticumGoAdvanced/internal/storage"
	"github.com/go-chi/chi/v5"
)

type APIHandler struct {
	*chi.Mux
	Storage *storage.Storage
	HashKey string
}

func New(s *storage.Storage, h string) APIHandler {
	mux := APIHandler{
		Mux:     chi.NewMux(),
		Storage: s,
		HashKey: h,
	}
	return mux
}
