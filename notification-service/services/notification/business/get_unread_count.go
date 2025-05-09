package business

import (
	"context"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) GetUnreadCount(ctx context.Context) (int64, error) {
	requester := core.GetRequester(ctx)
	uid, _ := core.FromBase58(requester.GetSubject())
	requesterId := int64(uid.GetLocalID())

	count, err := biz.notiRepo.GetUnreadCount(ctx, requesterId)
	if err != nil {
		return 0, core.ErrInternalServerError.
			WithError("cannot get unread notification count").
			WithDebug(err.Error())
	}

	return count, nil
}
