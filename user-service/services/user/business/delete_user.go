package business

import (
	"context"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) DeleteUser(ctx context.Context, id int) error {
	// Check if user exists
	user, err := biz.userRepo.GetUserById(ctx, id)
	if err != nil {
		return core.ErrInternalServerError.WithError("Failed to get user").WithDebug(err.Error())
	}
	if user == nil {
		return core.ErrBadRequest.WithError("User not found")
	}

	// Delete user
	if err := biz.userRepo.DeleteUser(ctx, id); err != nil {
		return core.ErrInternalServerError.WithError("Failed to delete user").WithDebug(err.Error())
	}

	return nil
} 