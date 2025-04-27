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
	DeleteUser(ctx context.Context, id int) error
	GetUserIdByEmail(ctx context.Context, email string) (int, error)
}

type ImageRepository interface {
	GetImageById(ctx context.Context, id int) (*core.Image, error)
	DeleteImage(ctx context.Context, id int) error
}

type NoteRepository interface {
	DeleteUserNotes(ctx context.Context, userId int32) (bool, int32, error)
}

type business struct {
	userRepo  UserRepository
	imageRepo ImageRepository
	noteRepo  NoteRepository
}

func NewBusiness(userRepo UserRepository, imageRepo ImageRepository, noteRepo NoteRepository) *business {
	return &business{
		userRepo:  userRepo,
		imageRepo: imageRepo,
		noteRepo:  noteRepo,
	}
}
