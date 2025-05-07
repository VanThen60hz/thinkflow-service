package business

import (
	"context"
	"time"

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
	ListUsers(ctx context.Context, filter *entity.UserFilter, paging *core.Paging) ([]entity.User, error)
	CountUsers(ctx context.Context) (int64, error)
	CountUsersByStatus(ctx context.Context, status string) (int64, error)
	CountUsersCreatedAfter(ctx context.Context, time time.Time) (int64, error)
}

type ImageRepository interface {
	GetImageById(ctx context.Context, id int) (*core.Image, error)
	DeleteImage(ctx context.Context, id int) error
}

type NoteRepository interface {
	CountNotes(ctx context.Context) (int64, error)
	DeleteUserNotes(ctx context.Context, userId int32) (bool, int32, error)
}

type AuthRepository interface {
	RegisterWithUserId(ctx context.Context, userId int32, email, password string) error
}

type Business interface {
	GetUserProfile(ctx context.Context) (*entity.User, error)
	ListUsers(ctx context.Context, filter *entity.UserFilter, paging *core.Paging) ([]entity.User, error)
	UpdateUserProfile(ctx context.Context, data *entity.UserDataUpdate) error
	UpdateUser(ctx context.Context, userId int, data *entity.UserDataUpdate) error
	DeleteUser(ctx context.Context, id int) error
	GetUserById(ctx context.Context, id int) (*entity.User, error)
	GetUserDetails(ctx context.Context, id int) (*entity.User, error)
	DeactivateUser(ctx context.Context, userId int) (*DeactivateUserResult, error)
	CreateUser(ctx context.Context, data *entity.UserDataCreation) (*entity.User, error)
	GetDashboardStats(ctx context.Context) (*DashboardStats, error)
}

type business struct {
	userRepo  UserRepository
	imageRepo ImageRepository
	noteRepo  NoteRepository
	authRepo  AuthRepository
}

func NewBusiness(userRepo UserRepository, imageRepo ImageRepository, noteRepo NoteRepository, authRepo AuthRepository) *business {
	return &business{
		userRepo:  userRepo,
		imageRepo: imageRepo,
		noteRepo:  noteRepo,
		authRepo:  authRepo,
	}
}
