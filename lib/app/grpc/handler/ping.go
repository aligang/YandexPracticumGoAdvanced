package handler

import (
	"context"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/app/grpc/generated/common"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *GrpcHandler) Ping(context.Context, *empty.Empty) (*common.Metric, error) {
	return nil, nil
}
