package handler

import (
	"github.com/aligang/YandexPracticumGoAdvanced/internal/storage"
	"github.com/go-chi/chi/v5"
)

type APIHandler struct {
	*chi.Mux
	Storage *storage.Storage
}

func New(s *storage.Storage) APIHandler {
	mux := APIHandler{
		Mux:     chi.NewMux(),
		Storage: s,
	}
	return mux
}
