package business

import (
	"context"

	"thinkflow-service/services/image/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) UpdateImage(ctx context.Context, id int, data *entity.ImageDataUpdate) error {
	_, err := biz.imageRepo.GetImageById(ctx, id)
	if err != nil {
		if err == core.ErrRecordNotFound {
			return core.ErrNotFound.
				WithError(entity.ErrCannotGetImage.Error()).
				WithDebug(err.Error())
		}

		return core.ErrInternalServerError.
			WithError(entity.ErrCannotGetImage.Error()).
			WithDebug(err.Error())
	}

	if err := biz.imageRepo.UpdateImage(ctx, id, data); err != nil {
		return core.ErrInternalServerError.
			WithError(entity.ErrCannotUpdateImage.Error()).
			WithDebug(err.Error())
	}

	return nil
}
