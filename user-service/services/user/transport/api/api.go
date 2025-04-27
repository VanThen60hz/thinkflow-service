package api

import (
	"context"

	"thinkflow-service/services/user/entity"
)

type Business interface {
	GetUserProfile(ctx context.Context) (*entity.User, error)
	UpdateUserProfile(ctx context.Context, data *entity.UserDataUpdate) error
	DeleteUser(ctx context.Context, id int) error
}

type api struct {
	business Business
}

func NewAPI(business Business) *api {
	return &api{business: business}
}
