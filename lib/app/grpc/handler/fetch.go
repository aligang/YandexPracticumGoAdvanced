package handler

import (
	"context"
	"errors"
	appErrors "github.com/aligang/YandexPracticumGoAdvanced/lib/app/base/errors"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/app/grpc/converter"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/app/grpc/generated/common"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *GrpcHandler) Fetch(ctx context.Context, in *common.Metric) (*common.Metric, error) {
	m := converter.ConvertMetric(in)
	res, err := s.BaseFetch(m)
	if err != nil {
		switch {
		case errors.As(err, &appErrors.InvalidMetricType):
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		default:
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}
	}
	return converter.ConvertMetricEntity(res), nil
}
