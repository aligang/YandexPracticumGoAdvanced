package handler

import (
	baseHandler "github.com/aligang/YandexPracticumGoAdvanced/lib/app/base/handler"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/storage"
	"github.com/go-chi/chi/v5"
)

type HTTPHandler struct {
	*chi.Mux
	*baseHandler.BaseHandler
}

func New(s storage.Storage, h string, t string) *HTTPHandler {
	return &HTTPHandler{
		Mux:         chi.NewMux(),
		BaseHandler: baseHandler.New(s, h, t),
	}
}
