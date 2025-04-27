package business

import (
	"context"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) HasReadPermission(ctx context.Context, noteId int, userId int) (bool, error) {
	hasPermission, err := biz.collabRepo.HasReadPermission(ctx, noteId, userId)
	if err != nil {
		return false, core.ErrInternalServerError.
			WithError("Cannot check read permission").
			WithDebug(err.Error())
	}

	return hasPermission, nil
} 