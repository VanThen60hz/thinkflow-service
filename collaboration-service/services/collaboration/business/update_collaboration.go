package business

import (
	"context"

	collaborationEntity "thinkflow-service/services/collaboration/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) UpdateCollaboration(ctx context.Context, id int, data *collaborationEntity.Collaboration) error {
	if err := biz.collabRepo.UpdateCollaboration(ctx, id, data); err != nil {
		return core.ErrInternalServerError.
			WithError("Cannot update collaboration").
			WithDebug(err.Error())
	}

	return nil
} 