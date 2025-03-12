package business

import (
	"context"

	"thinkflow-service/services/media/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) ListImages(ctx context.Context, filter *entity.Filter, paging *core.Paging) ([]entity.Image, error) {
	images, err := biz.mediaRepo.ListImages(ctx, filter, paging)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotListMedia.Error()).
			WithDebug(err.Error())
	}

	return images, nil
}

func (biz *business) ListAudios(ctx context.Context, filter *entity.Filter, paging *core.Paging) ([]entity.Audio, error) {
	audios, err := biz.mediaRepo.ListAudios(ctx, filter, paging)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotListMedia.Error()).
			WithDebug(err.Error())
	}

	// Get extra infos: User
	userIds := make([]int, len(audios))

	for i := range userIds {
		userIds[i] = audios[i].UserId
	}

	return audios, nil
}
