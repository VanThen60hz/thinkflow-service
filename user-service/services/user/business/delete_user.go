package business

import (
	"context"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) DeleteUser(ctx context.Context, id int) error {
	user, err := biz.userRepo.GetUserById(ctx, id)
	if err != nil {
		return core.ErrInternalServerError.WithError("Failed to get user").WithDebug(err.Error())
	}
	if user == nil {
		return core.ErrBadRequest.WithError("User not found")
	}

	// Delete user's notes first
	if _, _, err := biz.noteRepo.DeleteUserNotes(ctx, int32(id)); err != nil {
		return core.ErrInternalServerError.
			WithError("Failed to delete user's notes").
			WithDebug(err.Error())
	}

	if err := biz.userRepo.DeleteUser(ctx, id); err != nil {
		return core.ErrInternalServerError.WithError("Failed to delete user").WithDebug(err.Error())
	}

	if user.AvatarId != 0 {
		if err := biz.imageRepo.DeleteImage(ctx, user.AvatarId); err != nil {
			return core.ErrInternalServerError.
				WithError("User deleted but failed to delete image").
				WithDebug(err.Error())
		}
	}

	return nil
}
