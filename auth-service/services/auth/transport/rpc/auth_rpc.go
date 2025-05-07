package rpc

import (
	"context"

	"thinkflow-service/services/auth/entity"

	"github.com/golang-jwt/jwt/v5"
)

type Business interface {
	IntrospectToken(ctx context.Context, accessToken string) (*jwt.RegisteredClaims, error)
	RegisterWithUserId(ctx context.Context, data *entity.AuthRegister, newUserId int) error
}

type grpcService struct {
	business Business
}

func NewService(business Business) *grpcService {
	return &grpcService{business: business}
}
