package business

import (
	"context"

	"thinkflow-service/services/audio/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) CreateNewAudio(ctx context.Context, data *entity.AudioDataCreation) error {
	data.Prepare()

	if err := data.Validate(); err != nil {
		return core.ErrBadRequest.WithError(err.Error())
	}

	if err := biz.audioRepo.AddNewAudio(ctx, data); err != nil {
		return core.ErrInternalServerError.WithError(entity.ErrCannotCreateAudio.Error())
	}

	return nil
}
