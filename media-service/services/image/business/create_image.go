package business

import (
	"context"

	"thinkflow-service/services/image/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) CreateNewImage(ctx context.Context, data *entity.ImageDataCreation) error {
	data.Prepare()

	if err := data.Validate(); err != nil {
		return core.ErrBadRequest.WithError(err.Error())
	}

	if err := biz.imageRepo.AddNewImage(ctx, data); err != nil {
		return core.ErrInternalServerError.WithError(entity.ErrCannotCreateImage.Error())
	}

	return nil
}
