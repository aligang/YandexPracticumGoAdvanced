package handler

import (
	"context"
	"errors"
	appErrors "github.com/aligang/YandexPracticumGoAdvanced/lib/app/base/errors"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/app/grpc/converter"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/app/grpc/generated/common"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *GrpcHandler) Update(ctx context.Context, in *common.Metric) (*empty.Empty, error) {
	m := converter.ConvertMetric(in)
	err := s.BaseUpdate(m)
	if err != nil {
		switch {
		case errors.Is(err, appErrors.ErrInvalidMetricType):
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		case errors.Is(err, appErrors.ErrInvalidHashValue):
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		default:
			return nil, status.Errorf(codes.Unavailable, errors.Unwrap(err).Error())
		}
	}
	return &empty.Empty{}, nil
}
