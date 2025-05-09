package business

import (
	"context"

	"thinkflow-service/services/notification/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) DeleteNotification(ctx context.Context, notiId int) error {
	note, err := biz.notiRepo.GetNotificationById(ctx, notiId)
	if err != nil {
		if err == core.ErrRecordNotFound {
			return core.ErrNotFound.
				WithError(entity.ErrCannotGetNotification.Error()).
				WithDebug(err.Error())
		}

		return core.ErrInternalServerError.
			WithError(entity.ErrCannotGetNotification.Error()).
			WithDebug(err.Error())
	}

	requester := core.GetRequester(ctx)

	uid, _ := core.FromBase58(requester.GetSubject())
	requesterId := int(uid.GetLocalID())

	if requesterId != int(note.NotiReceivedID) {
		return core.ErrForbidden.WithError(entity.ErrRequesterIsNotReceivedUser.Error())
	}

	if err := biz.notiRepo.DeleteNotification(ctx, notiId); err != nil {
		return core.ErrInternalServerError.
			WithError(entity.ErrCannotDeleteNotification.Error()).
			WithDebug(err.Error())
	}

	return nil
}
