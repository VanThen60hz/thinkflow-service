package business

import (
	"context"

	"thinkflow-service/services/image/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) DeleteImage(ctx context.Context, id int) error {
	// Get media data, without extra infos
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

	if err := biz.imageRepo.DeleteImage(ctx, id); err != nil {
		return core.ErrInternalServerError.
			WithError(entity.ErrCannotDeleteImage.Error()).
			WithDebug(err.Error())
	}

	return nil
}
