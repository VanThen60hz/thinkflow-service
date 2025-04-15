package business

import (
	"context"

	"thinkflow-service/common"
	"thinkflow-service/services/auth/entity"

	"github.com/VanThen60hz/service-context/component/emailc"
	"github.com/VanThen60hz/service-context/component/oauthc"
	"github.com/VanThen60hz/service-context/component/redisc"
	"gorm.io/gorm"
)

type AuthRepository interface {
	AddNewAuth(ctx context.Context, data *entity.Auth) error
	GetAuth(ctx context.Context, email string) (*entity.Auth, error)
	UpdatePassword(ctx context.Context, email, salt, hashedPassword string) error
	DeleteAuth(ctx context.Context, email string) error
}

type UserRepository interface {
	CreateUser(ctx context.Context, firstName, lastName, email string) (newId int, err error)
	GetUserIdByEmail(ctx context.Context, email string) (int, error)
	UpdateUserStatus(ctx context.Context, id int, status string) error
	GetUserStatus(ctx context.Context, id int) (string, error)
	DeleteUser(ctx context.Context, id int) error
	GetDB() *gorm.DB
}

type Hasher interface {
	RandomStr(length int) (string, error)
	HashPassword(salt, password string) (string, error)
	CompareHashPassword(hashedPassword, salt, password string) bool
}

type business struct {
	repository     AuthRepository
	userRepository UserRepository
	jwtProvider    common.JWTProvider
	hasher         Hasher
	redisClient    redisc.Redis
	emailService   emailc.Email
	oauthProvider  oauthc.OAuth
}

func NewBusiness(repository AuthRepository, userRepository UserRepository,
	jwtProvider common.JWTProvider, hasher Hasher,
	redisClient redisc.Redis, emailService emailc.Email,
	oauth oauthc.OAuth,
) *business {
	return &business{
		repository:     repository,
		userRepository: userRepository,
		jwtProvider:    jwtProvider,
		hasher:         hasher,
		redisClient:    redisClient,
		emailService:   emailService,
		oauthProvider:  oauth,
	}
}
