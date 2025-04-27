package rpc

import (
	"context"
)

type Business interface {
	DeleteUserNotes(ctx context.Context, userId int) (int, error)
}

type grpcService struct {
	business Business
}

func NewService(business Business) *grpcService {
	return &grpcService{business: business}
}
