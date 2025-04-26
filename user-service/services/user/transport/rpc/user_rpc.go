package rpc

import (
	"context"

	"thinkflow-service/services/user/entity"
)

type Business interface {
	GetUserDetails(ctx context.Context, id int) (*entity.User, error)
	GetUsersByIds(ctx context.Context, ids []int) ([]entity.User, error)
	CreateNewUser(ctx context.Context, data *entity.UserDataCreation) error
	UpdateUserStatus(ctx context.Context, id int, status string) error
	GetUserStatus(ctx context.Context, id int) (string, error)
	DeleteUser(ctx context.Context, id int) error
	GetUserIdByEmail(ctx context.Context, email string) (int, error)
}

type grpcService struct {
	business Business
}

func NewService(business Business) *grpcService {
	return &grpcService{business: business}
}
