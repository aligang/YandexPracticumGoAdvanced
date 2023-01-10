package handler

import (
	"context"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/app/grpc/converter"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/app/grpc/generated/common"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/hash"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/logging"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *GrpcHandler) Update(ctx context.Context, in *common.Metric) (*empty.Empty, error) {
	m := converter.ConvertMetric(in)

	if m.MType != "gauge" && m.MType != "counter" {
		logging.Warn("Invalid Metric Type")
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Metric Type")
	}
	if s.Config.HashKey != "" {
		logging.Debug("Validating hash ...")
		if !hash.CheckHashInfo(&m, s.Config.HashKey) {
			logging.Warn("Invalid Hash")
			return nil, status.Errorf(codes.InvalidArgument, "Invalid hash value")
		} else {
			logging.Debug("Hash validation succeeded")
		}
	} else {
		logging.Debug("Skipping hash validation")
	}
	err := s.Storage.Update(m)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, err.Error())
	}
	return &empty.Empty{}, nil
}
