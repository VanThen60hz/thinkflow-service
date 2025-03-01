package business

import (
	"context"

	"thinkflow-service/services/media/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) UpdateImage(ctx context.Context, id int, data *entity.ImageDataUpdate) error {
	// Get media data, without extra infos
	_, err := biz.mediaRepo.GetImageById(ctx, id)
	if err != nil {
		if err == core.ErrRecordNotFound {
			return core.ErrNotFound.
				WithError(entity.ErrCannotGetMedia.Error()).
				WithDebug(err.Error())
		}

		return core.ErrInternalServerError.
			WithError(entity.ErrCannotGetMedia.Error()).
			WithDebug(err.Error())
	}

	if err := biz.mediaRepo.UpdateImage(ctx, id, data); err != nil {
		return core.ErrInternalServerError.
			WithError(entity.ErrCannotUpdateMedia.Error()).
			WithDebug(err.Error())
	}

	return nil
}

func (biz *business) UpdateAudio(ctx context.Context, id int, data *entity.AudioDataUpdate) error {
	// Get media data, without extra infos
	_, err := biz.mediaRepo.GetAudioById(ctx, id)
	if err != nil {
		if err == core.ErrRecordNotFound {
			return core.ErrNotFound.
				WithError(entity.ErrCannotGetMedia.Error()).
				WithDebug(err.Error())
		}

		return core.ErrInternalServerError.
			WithError(entity.ErrCannotGetMedia.Error()).
			WithDebug(err.Error())
	}

	if err := biz.mediaRepo.UpdateAudio(ctx, id, data); err != nil {
		return core.ErrInternalServerError.
			WithError(entity.ErrCannotUpdateMedia.Error()).
			WithDebug(err.Error())
	}

	return nil
}
