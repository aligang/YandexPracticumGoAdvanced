package handler

import (
	"context"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/app/grpc/converter"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/app/grpc/generated/service"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/hash"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/logging"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/metric"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *GrpcHandler) BulkUpdate(ctx context.Context, req *service.BulkUpdateRequest) (*empty.Empty, error) {

	aggregatedMetrics := map[string]metric.Metrics{}
	for _, grpcMetric := range req.Metrics {
		m := converter.ConvertMetric(grpcMetric)
		if m.MType != "gauge" && m.MType != "counter" {
			logging.Warn("Invalid Metric Type")
			return nil, status.Errorf(codes.InvalidArgument, "Invalid Metric Type")
		}
		if s.Config.HashKey != "" {
			logging.Debug("Validating hash ...")
			if !hash.CheckHashInfo(&m, s.Config.HashKey) {
				logging.Warn("Invalid Hash")
				return nil, status.Errorf(codes.InvalidArgument, "Invalid hash value")
			}
			logging.Debug("Hash validation succeeded")
		} else {
			logging.Debug("Skipping hash validation")
		}
		_, found := aggregatedMetrics[m.ID]
		if m.MType == "counter" && found {
			*aggregatedMetrics[m.ID].Delta += *m.Delta
		} else {
			aggregatedMetrics[m.ID] = m
		}
	}

	err := s.Storage.BulkUpdate(aggregatedMetrics)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, err.Error())
	}

	return &empty.Empty{}, nil
}
