package handler

import (
	"context"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/app/grpc/generated/common"
)

func (s *GrpcHandler) Get(context.Context, *common.Metric) (*common.Metric, error) {
	return nil, nil
}
