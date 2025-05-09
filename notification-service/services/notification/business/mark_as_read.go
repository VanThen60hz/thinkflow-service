package business

import (
	"context"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) MarkNotificationAsRead(ctx context.Context, notiId string) error {
	err := biz.notiRepo.MarkNotificationAsRead(ctx, notiId)
	if err != nil {
		return core.ErrInternalServerError.
			WithError("cannot mark notification as read").
			WithDebug(err.Error())
	}

	return nil
}
