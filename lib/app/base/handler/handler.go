package handler

import (
	"github.com/aligang/YandexPracticumGoAdvanced/lib/app"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/storage"
)

type BaseHandler struct {
	Storage storage.Storage
	Config  app.Config
}

func New(s storage.Storage, h string, t string) *BaseHandler {
	return &BaseHandler{
		Storage: s,
		Config: app.Config{
			HashKey:     h,
			StorageType: t,
		},
	}
}
