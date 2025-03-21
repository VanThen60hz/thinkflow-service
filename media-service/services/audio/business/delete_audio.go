package business

import (
	"context"

	"thinkflow-service/services/audio/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) DeleteAudio(ctx context.Context, id int) error {
	// Get media data, without extra infos
	_, err := biz.audioRepo.GetAudioById(ctx, id)
	if err != nil {
		if err == core.ErrRecordNotFound {
			return core.ErrNotFound.
				WithError(entity.ErrCannotGetAudio.Error()).
				WithDebug(err.Error())
		}

		return core.ErrInternalServerError.
			WithError(entity.ErrCannotGetAudio.Error()).
			WithDebug(err.Error())
	}

	if err := biz.audioRepo.DeleteAudio(ctx, id); err != nil {
		return core.ErrInternalServerError.
			WithError(entity.ErrCannotDeleteAudio.Error()).
			WithDebug(err.Error())
	}

	return nil
}
