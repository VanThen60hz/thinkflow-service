package business

import (
	"context"

	collaborationEntity "thinkflow-service/services/collaboration/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) GetCollaborationByNoteId(ctx context.Context, noteId int, paging *core.Paging) ([]collaborationEntity.Collaboration, error) {
	collabs, err := biz.collabRepo.GetCollaborationByNoteId(ctx, noteId, paging)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError("Cannot get collaborations").
			WithDebug(err.Error())
	}

	return collabs, nil
} 