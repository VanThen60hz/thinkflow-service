package utils

import (
	"context"

	"github.com/VanThen60hz/service-context/core"
)

// CheckUserStatus checks if the user has the expected status
func CheckUserStatus(ctx context.Context, userRepository UserRepository, userId int, expectedStatus string) error {
	status, err := userRepository.GetUserStatus(ctx, userId)
	if err != nil {
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	if status != expectedStatus {
		return core.ErrBadRequest.WithError("User status is not " + expectedStatus)
	}

	return nil
}

// IsUserWaitingVerification checks if the user is waiting for email verification
func IsUserWaitingVerification(ctx context.Context, userRepository UserRepository, userId int) (bool, error) {
	status, err := userRepository.GetUserStatus(ctx, userId)
	if err != nil {
		return false, core.ErrInternalServerError.WithDebug(err.Error())
	}

	return status == "waiting_verify", nil
}

// UserRepository interface for user operations
type UserRepository interface {
	GetUserStatus(ctx context.Context, userId int) (string, error)
	UpdateUserStatus(ctx context.Context, userId int, status string) error
}
