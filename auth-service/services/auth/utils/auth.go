package utils

import (
	"context"
	"fmt"

	"thinkflow-service/services/auth/entity"

	"github.com/VanThen60hz/service-context/core"
)

func ValidateEmailAndGetAuth(ctx context.Context, repository AuthRepository, email string) (*entity.Auth, error) {
	auth, err := repository.GetAuth(ctx, email)
	if err != nil {
		if err == core.ErrRecordNotFound {
			return nil, core.ErrBadRequest.WithError(entity.ErrEmailNotFound.Error())
		}
		return nil, core.ErrInternalServerError.WithDebug(err.Error())
	}
	return auth, nil
}

func ProcessPassword(hasher Hasher, password string) (salt, hashedPassword string, err error) {
	salt, err = hasher.RandomStr(16)
	if err != nil {
		return "", "", core.ErrInternalServerError.WithDebug(err.Error())
	}

	hashedPassword, err = hasher.HashPassword(salt, password)
	if err != nil {
		return "", "", core.ErrInternalServerError.WithDebug(err.Error())
	}

	return salt, hashedPassword, nil
}

func CompensateUserCreation(ctx context.Context, userRepository UserRepository, userId int) {
	if err := userRepository.DeleteUser(ctx, userId); err != nil {
		fmt.Printf("Failed to compensate user creation: %v\n", err)
	}
}

func CompensateAuthCreation(ctx context.Context, repository AuthRepository, email string) {
	if err := repository.DeleteAuth(ctx, email); err != nil {
		fmt.Printf("Failed to compensate auth creation: %v\n", err)
	}
}
