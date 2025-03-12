package business

import (
	"context"

	"thinkflow-service/services/user/entity"

	"github.com/VanThen60hz/service-context/core"
)

type UserRepository interface {
	GetUserById(ctx context.Context, id int) (*entity.User, error)
	GetUsersByIds(ctx context.Context, ids []int) ([]entity.User, error)
	CreateNewUser(ctx context.Context, data *entity.UserDataCreation) error
	UpdateUser(ctx context.Context, id int, data *entity.UserDataUpdate) error
}

// type business struct {
// 	repository UserRepository
// }

// func NewBusiness(repository UserRepository) *business {
// 	return &business{repository: repository}
// }

type ImageRepository interface {
	GetImagesByIds(ctx context.Context, ids []int) ([]core.Image, error)
	GetImageById(ctx context.Context, id int) (*core.Image, error)
}

type business struct {
	userRepo  UserRepository
	imageRepo ImageRepository
}

func NewBusiness(userRepo UserRepository, imageRepo ImageRepository) *business {
	return &business{
		userRepo:  userRepo,
		imageRepo: imageRepo,
	}
}
