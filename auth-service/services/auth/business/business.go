package business

import (
	"context"
	"math/rand"
	"time"

	"thinkflow-service/common"
	"thinkflow-service/services/auth/entity"

	"gorm.io/gorm"
)

type AuthRepository interface {
	AddNewAuth(ctx context.Context, data *entity.Auth) error
	GetAuth(ctx context.Context, email string) (*entity.Auth, error)
	UpdatePassword(ctx context.Context, email, salt, hashedPassword string) error
}

type UserRepository interface {
	CreateUser(ctx context.Context, firstName, lastName, email string) (newId int, err error)
	GetUserIdByEmail(ctx context.Context, email string) (int, error)
	UpdateUserStatus(ctx context.Context, id int, status string) error
	GetUserStatus(ctx context.Context, id int) (string, error)
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
	redisClient    *common.RedisClient
	emailService   *common.EmailService
}

func NewBusiness(repository AuthRepository, userRepository UserRepository,
	jwtProvider common.JWTProvider, hasher Hasher,
	redisClient *common.RedisClient, emailService *common.EmailService,
) *business {
	return &business{
		repository:     repository,
		userRepository: userRepository,
		jwtProvider:    jwtProvider,
		hasher:         hasher,
		redisClient:    redisClient,
		emailService:   emailService,
	}
}

func generateOTP() string {
	rand.Seed(time.Now().UnixNano())
	digits := "0123456789"
	otp := ""
	for i := 0; i < 6; i++ {
		otp += string(digits[rand.Intn(len(digits))])
	}
	return otp
}
