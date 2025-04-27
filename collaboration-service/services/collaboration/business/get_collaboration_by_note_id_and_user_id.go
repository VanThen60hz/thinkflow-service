package business

import (
	"context"

	collaborationEntity "thinkflow-service/services/collaboration/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) GetCollaborationByNoteIdAndUserId(ctx context.Context, noteId int, userId int) (*collaborationEntity.Collaboration, error) {
	collab, err := biz.collabRepo.GetCollaborationByNoteIdAndUserId(ctx, noteId, userId)
	if err != nil {
		if err == core.ErrRecordNotFound {
			return nil, core.ErrNotFound.
				WithError("Collaboration not found").
				WithDebug(err.Error())
		}

		return nil, core.ErrInternalServerError.
			WithError("Cannot get collaboration").
			WithDebug(err.Error())
	}

	return collab, nil
} 