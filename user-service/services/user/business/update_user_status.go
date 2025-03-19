package business

import (
	"context"

	"thinkflow-service/services/user/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) UpdateUserStatus(ctx context.Context, id int, statusStr string) error {
	// Convert string to Status type
	status := entity.Status(statusStr)

	// Validate status
	if status != entity.StatusActive && status != entity.StatusPendingVerify && status != entity.StatusBanned {
		return core.ErrBadRequest.WithError("invalid status")
	}

	// Create update data
	data := &entity.UserDataUpdate{
		Status: &status,
	}

	// Update user status
	if err := biz.userRepo.UpdateUser(ctx, id, data); err != nil {
		return core.ErrInternalServerError.
			WithError("cannot update user status").
			WithDebug(err.Error())
	}

	return nil
}
