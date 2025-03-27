package business

import (
	"context"

	"thinkflow-service/services/mindmap/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) CreateNewMindmap(ctx context.Context, data *entity.MindmapDataCreation) error {
	requester := core.GetRequester(ctx)

	uid, _ := core.FromBase58(requester.GetSubject())
	requesterId := int(uid.GetLocalID()) // mindmap user id, id of who creates this new Mindmap

	data.Prepare(requesterId)

	if err := data.Validate(); err != nil {
		return core.ErrBadRequest.WithError(err.Error())
	}

	if err := biz.mindmapRepo.AddNewMindmap(ctx, data); err != nil {
		return core.ErrInternalServerError.WithError(entity.ErrCannotCreateMindmap.Error())
	}

	return nil
}
