package business

import (
	"context"

	"thinkflow-service/services/media/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) CreateNewImage(ctx context.Context, data *entity.ImageDataCreation) error {
	data.Prepare()

	if err := data.Validate(); err != nil {
		return core.ErrBadRequest.WithError(err.Error())
	}

	if err := biz.mediaRepo.AddNewImage(ctx, data); err != nil {
		return core.ErrInternalServerError.WithError(entity.ErrCannotCreateMedia.Error())
	}

	return nil
}

func (biz *business) CreateNewAudio(ctx context.Context, data *entity.AudioDataCreation) error {
	requester := core.GetRequester(ctx)

	uid, _ := core.FromBase58(requester.GetSubject())
	requesterId := int(uid.GetLocalID())

	data.Prepare(requesterId)

	if err := data.Validate(); err != nil {
		return core.ErrBadRequest.WithError(err.Error())
	}

	if err := biz.mediaRepo.AddNewAudio(ctx, data); err != nil {
		return core.ErrInternalServerError.WithError(entity.ErrCannotCreateMedia.Error())
	}

	return nil
}
