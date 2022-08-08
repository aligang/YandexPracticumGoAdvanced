package handler

import (
	"github.com/aligang/YandexPracticumGoAdvanced/internal/storage"
	"github.com/go-chi/chi/v5"
)

type ApiHandler struct {
	*chi.Mux
	Storage *storage.Storage
}

func New(s *storage.Storage) ApiHandler {
	mux := ApiHandler{
		Mux:     chi.NewMux(),
		Storage: s,
	}
	return mux
}
