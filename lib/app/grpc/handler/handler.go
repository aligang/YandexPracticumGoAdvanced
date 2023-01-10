package handler

import (
	baseHandler "github.com/aligang/YandexPracticumGoAdvanced/lib/app/base/handler"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/app/grpc/generated/service"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/storage"
)

type GrpcHandler struct {
	service.UnimplementedMetricsServiceServer
	*baseHandler.BaseHandler
}

func New(s storage.Storage, h string, t string) *GrpcHandler {
	return &GrpcHandler{
		BaseHandler: baseHandler.New(s, h, t),
	}
}
