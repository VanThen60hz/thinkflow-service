package business

import (
	"context"

	"thinkflow-service/common"
	"thinkflow-service/services/media/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) GetImageById(ctx context.Context, id int) (*entity.Image, error) {
	data, err := biz.mediaRepo.GetImageById(ctx, id)
	if err != nil {
		if err == core.ErrRecordNotFound {
			return nil, core.ErrNotFound.
				WithDebug(err.Error())
		}

		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotGetMedia.Error()).
			WithDebug(err.Error())
	}

	return data, nil
}

func (biz *business) GetAudioById(ctx context.Context, id int) (*entity.Audio, error) {
	data, err := biz.mediaRepo.GetAudioById(ctx, id)
	if err != nil {
		if err == core.ErrRecordNotFound {
			return nil, core.ErrNotFound.
				WithDebug(err.Error())
		}

		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotGetMedia.Error()).
			WithDebug(err.Error())
	}

	requesterVal := ctx.Value(common.RequesterKey)
	if requesterVal == nil {
		return nil, core.ErrUnauthorized.WithError("requester not found in context")
	}

	requester, ok := requesterVal.(core.Requester)
	if !ok {
		return nil, core.ErrInternalServerError.
			WithError("invalid requester type in context")
	}

	uid, _ := core.FromBase58(requester.GetSubject())
	requesterId := int(uid.GetLocalID())

	if requesterId != data.UserId {
		return nil, core.ErrForbidden.WithError(entity.ErrRequesterIsNotOwner.Error())
	}

	return data, nil
}
