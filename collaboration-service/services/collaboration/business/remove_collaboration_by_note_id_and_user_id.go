package business

import (
	"context"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) RemoveCollaborationByNoteIdAndUserId(ctx context.Context, noteId int, userId int) error {
	if err := biz.collabRepo.RemoveCollaborationByNoteIdAndUserId(ctx, noteId, userId); err != nil {
		return core.ErrInternalServerError.
			WithError("Cannot remove collaboration").
			WithDebug(err.Error())
	}

	return nil
} 