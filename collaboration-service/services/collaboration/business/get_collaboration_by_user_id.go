package business

import (
	"context"

	collaborationEntity "thinkflow-service/services/collaboration/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) GetCollaborationByUserId(ctx context.Context, userId int, paging *core.Paging) ([]collaborationEntity.Collaboration, error) {
	collabs, err := biz.collabRepo.GetCollaborationByUserId(ctx, userId, paging)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError("Cannot get collaborations").
			WithDebug(err.Error())
	}

	return collabs, nil
} 