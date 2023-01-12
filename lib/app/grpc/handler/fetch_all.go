package handler

import (
	"context"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/app/grpc/converter"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/app/grpc/generated/service"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *GrpcHandler) FetchAll(ctx context.Context, in *empty.Empty) (*service.FetchAllResponse, error) {
	r := &service.FetchAllResponse{}
	for name, m := range s.BaseFetchAll() {
		r.Metrics[name] = converter.ConvertMetricEntity(&m)
	}
	return r, nil
}
