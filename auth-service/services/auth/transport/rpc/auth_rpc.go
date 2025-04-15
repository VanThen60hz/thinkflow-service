package rpc

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
)

type Business interface {
	IntrospectToken(ctx context.Context, accessToken string) (*jwt.RegisteredClaims, error)
}

type grpcService struct {
	business Business
}

func NewService(business Business) *grpcService {
	return &grpcService{business: business}
}