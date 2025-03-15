package business

import (
	"context"

	"thinkflow-service/services/audio/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) CreateNewAudio(ctx context.Context, data *entity.AudioDataCreation) error {
	requester := core.GetRequester(ctx)

	uid, _ := core.FromBase58(requester.GetSubject())
	requesterId := int(uid.GetLocalID())

	data.Prepare(requesterId)

	if err := data.Validate(); err != nil {
		return core.ErrBadRequest.WithError(err.Error())
	}

	if err := biz.audioRepo.AddNewAudio(ctx, data); err != nil {
		return core.ErrInternalServerError.WithError(entity.ErrCannotCreateAudio.Error())
	}

	return nil
}
