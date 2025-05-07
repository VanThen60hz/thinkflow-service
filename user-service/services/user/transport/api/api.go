package api

import (
	"context"

	"thinkflow-service/services/user/business"
	"thinkflow-service/services/user/entity"

	"github.com/VanThen60hz/service-context/core"
)

type Business interface {
	CreateUserByAdmin(ctx context.Context, data *business.CreateUserData) (*entity.User, error)
	DeactivateUser(ctx context.Context, userId int) (*business.DeactivateUserResult, error)
	GetUserProfile(ctx context.Context) (*entity.User, error)
	ListUsers(ctx context.Context, filter *entity.UserFilter, paging *core.Paging) ([]entity.User, error)
	GetDashboardStats(ctx context.Context) (*business.DashboardStats, error)
	UpdateUserProfile(ctx context.Context, data *entity.UserDataUpdate) error
	UpdateUser(ctx context.Context, userId int, data *entity.UserDataUpdate) error
	DeleteUser(ctx context.Context, id int) error
}

type api struct {
	business Business
}

func NewAPI(business Business) *api {
	return &api{business: business}
}
