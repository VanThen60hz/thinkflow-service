package business

import (
	"context"

	"thinkflow-service/services/image/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) GetImageById(ctx context.Context, id int) (*entity.Image, error) {
	data, err := biz.imageRepo.GetImageById(ctx, id)
	if err != nil {
		if err == core.ErrRecordNotFound {
			return nil, core.ErrNotFound.
				WithDebug(err.Error())
		}

		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotGetImage.Error()).
			WithDebug(err.Error())
	}

	return data, nil
}
