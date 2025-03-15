package business

import (
	"context"

	"thinkflow-service/services/image/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) ListImages(ctx context.Context, filter *entity.Filter, paging *core.Paging) ([]entity.Image, error) {
	images, err := biz.imageRepo.ListImages(ctx, filter, paging)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotListImage.Error()).
			WithDebug(err.Error())
	}

	return images, nil
}
