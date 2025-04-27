package business

import (
	"context"

	collaborationEntity "thinkflow-service/services/collaboration/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) AddNewCollaboration(ctx context.Context, data *collaborationEntity.CollaborationCreation) error {
	if err := biz.collabRepo.AddNewCollaboration(ctx, data); err != nil {
		return core.ErrInternalServerError.
			WithError(collaborationEntity.ErrCannotCreateCollab.Error()).
			WithDebug(err.Error())
	}

	return nil
}
