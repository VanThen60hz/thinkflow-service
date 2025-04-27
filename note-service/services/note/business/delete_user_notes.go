package business

import (
	"context"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) DeleteUserNotes(ctx context.Context, userId int) (int, error) {
	deletedCount, err := biz.noteRepo.DeleteUserNotes(ctx, userId)
	if err != nil {
		return 0, core.ErrInternalServerError.
			WithError("Failed to delete user notes").
			WithDebug(err.Error())
	}

	return deletedCount, nil
}
