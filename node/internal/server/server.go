package server

import (
	"context"
	"github.com/ShudderStorm/CRStorage/api/service/kvs"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Storage interface {
	Set(key, value string) error
	Get(key string) (string, error)
}

type Server struct {
	kvs.UnimplementedKVSServer
	storage Storage
}

func (s Server) Set(ctx context.Context, request *kvs.SetRequest) (*kvs.Status, error) {
	err := s.storage.Set(request.Key, request.Value)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to set: %v", err)
	}

	return &kvs.Status{Success: true}, nil
}

func (s Server) Get(ctx context.Context, request *kvs.GetRequest) (*kvs.Value, error) {
	value, err := s.storage.Get(request.Key)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get: %v", err)
	}

	if value == "" {
		return nil, status.Errorf(codes.NotFound, "key not found")
	}

	return &kvs.Value{Value: value}, nil
}
