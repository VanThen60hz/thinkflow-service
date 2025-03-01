package business

import (
	"context"

	"thinkflow-service/services/media/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) DeleteImage(ctx context.Context, id int) error {
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

	if err := biz.mediaRepo.DeleteImage(ctx, id); err != nil {
		return core.ErrInternalServerError.
			WithError(entity.ErrCannotDeleteMedia.Error()).
			WithDebug(err.Error())
	}

	return nil
}

func (biz *business) DeleteAudio(ctx context.Context, id int) error {
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

	if err := biz.mediaRepo.DeleteAudio(ctx, id); err != nil {
		return core.ErrInternalServerError.
			WithError(entity.ErrCannotDeleteMedia.Error()).
			WithDebug(err.Error())
	}

	return nil
}
