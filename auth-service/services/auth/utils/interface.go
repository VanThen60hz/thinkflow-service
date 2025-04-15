package utils

import (
	"context"

	"thinkflow-service/services/auth/entity"

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
