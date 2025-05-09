package business

import (
	"context"

	"thinkflow-service/services/notification/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) CreateNotification(ctx context.Context, data *entity.NotificationCreation) error {
	err := biz.notiRepo.CreateNotification(ctx, data)
	if err != nil {
		return core.ErrInternalServerError.
			WithError(entity.ErrCannotCreateNotification.Error()).
			WithDebug(err.Error())
	}

	return nil
}
