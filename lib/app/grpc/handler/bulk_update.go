package handler

import (
	"context"
	"errors"
	appErrors "github.com/aligang/YandexPracticumGoAdvanced/lib/app/base/errors"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/app/grpc/converter"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/app/grpc/generated/service"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/metric"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *GrpcHandler) BulkUpdate(ctx context.Context, req *service.BulkUpdateRequest) (*empty.Empty, error) {

	metrics := []metric.Metrics{}
	for _, grpcMetric := range req.Metrics {
		metrics = append(metrics, converter.ConvertMetric(grpcMetric))
	}
	err := s.BaseBulkUpdate(metrics)
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
