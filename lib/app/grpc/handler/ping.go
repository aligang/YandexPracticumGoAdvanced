package handler

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *GrpcHandler) Ping(context.Context, *empty.Empty) (*empty.Empty, error) {
	err := s.BasePing()
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, err.Error())
	}
	return &empty.Empty{}, nil
}
