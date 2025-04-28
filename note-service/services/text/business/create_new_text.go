package business

import (
	"context"

	"thinkflow-service/services/text/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) CreateNewText(ctx context.Context, data *entity.TextDataCreation) error {
	requester := core.GetRequester(ctx)
	uid, _ := core.FromBase58(requester.GetSubject())
	requesterId := int(uid.GetLocalID())

	data.Prepare(requesterId)

	if err := data.Validate(); err != nil {
		return core.ErrBadRequest.WithError(err.Error())
	}

	if err := biz.textRepo.AddNewText(ctx, data); err != nil {
		return core.ErrInternalServerError.WithError(entity.ErrCannotCreateText.Error())
	}

	return nil
}
