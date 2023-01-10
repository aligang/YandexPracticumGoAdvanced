package handler

import (
	"github.com/aligang/YandexPracticumGoAdvanced/lib/app"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/app/grpc/generated/service"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/storage"
)

type GrpcHandler struct {
	service.UnimplementedMetricsServiceServer
	Storage storage.Storage
	Config  app.Config
}

func New(s storage.Storage, h string, t string) *GrpcHandler {
	return &GrpcHandler{
		Storage: s,
		Config: app.Config{
			HashKey:     h,
			StorageType: t,
		},
	}
}
