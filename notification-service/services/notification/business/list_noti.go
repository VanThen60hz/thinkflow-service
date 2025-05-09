package business

import (
	"context"

	"thinkflow-service/common"
	"thinkflow-service/services/notification/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) ListNotifications(ctx context.Context, filter *entity.Filter, paging *core.Paging) ([]entity.Notification, error) {
	requesterVal := ctx.Value(common.RequesterKey)
	if requesterVal == nil {
		return nil, core.ErrUnauthorized.WithError("requester not found in context")
	}

	requester, ok := requesterVal.(core.Requester)
	if !ok {
		return nil, core.ErrInternalServerError.WithError("invalid requester type in context")
	}
	requesterUid := requester.GetSubject()


	if filter == nil {
		filter = &entity.Filter{}
	}
	filter.NotiReceivedID = &requesterUid

	notis, err := biz.notiRepo.ListNotifications(ctx, filter, paging)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError("cannot list notifications").
			WithDebug(err.Error())
	}

	return notis, nil
}
