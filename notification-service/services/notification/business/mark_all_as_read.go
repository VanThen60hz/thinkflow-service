package business

import (
	"context"
	"errors"

	"thinkflow-service/common"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) MarkAllNotificationsAsRead(ctx context.Context) error {
	requesterVal := ctx.Value(common.RequesterKey)
	if requesterVal == nil {
		return core.ErrUnauthorized.WithError("requester not found in context")
	}

	requester, ok := requesterVal.(core.Requester)
	if !ok {
		return core.ErrInternalServerError.WithError("invalid requester type in context")
	}
	uid, _ := core.FromBase58(requester.GetSubject())
	requesterId := int(uid.GetLocalID())

	if requesterId == 0 {
		return core.ErrInvalidRequest(errors.New("requester id is required"))
	}

	if err := biz.notiRepo.MarkAllNotificationsAsRead(ctx, int64(requesterId)); err != nil {
		return core.ErrInternalServerError.WithError("cannot mark all notifications as read").WithDebug(err.Error())
	}

	return nil
}
