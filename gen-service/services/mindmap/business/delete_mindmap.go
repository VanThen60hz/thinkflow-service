package business

import (
	"context"

	"thinkflow-service/services/mindmap/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) DeleteMindmap(ctx context.Context, id int) error {
	// Get mindmap data, without extra infos
	// mindmap, err := biz.MindmapRepo.GetMindmapById(ctx, id)
	_, err := biz.mindmapRepo.GetMindmapById(ctx, id)
	if err != nil {
		if err == core.ErrRecordNotFound {
			return core.ErrNotFound.
				WithError(entity.ErrCannotGetMindmap.Error()).
				WithDebug(err.Error())
		}

		return core.ErrInternalServerError.
			WithError(entity.ErrCannotGetMindmap.Error()).
			WithDebug(err.Error())
	}

	// requester := core.GetRequester(ctx)

	// uid, _ := core.FromBase58(requester.GetSubject())
	// requesterId := int(uid.GetLocalID())

	// // Only Mindmap user can do this
	// if requesterId != Mindmap.UserId {
	// 	return core.ErrForbidden.WithError(entity.ErrRequesterIsNotOwner.Error())
	// }

	if err := biz.mindmapRepo.DeleteMindmap(ctx, id); err != nil {
		return core.ErrInternalServerError.
			WithError(entity.ErrCannotDeleteMindmap.Error()).
			WithDebug(err.Error())
	}

	return nil
}
