package business

import (
	"context"

	"thinkflow-service/services/summary/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) DeleteSummary(ctx context.Context, id int) error {
	// Get summary data, without extra infos
	// summary, err := biz.summaryRepo.GetSummaryById(ctx, id)
	_, err := biz.summaryRepo.GetSummaryById(ctx, id)
	if err != nil {
		if err == core.ErrRecordNotFound {
			return core.ErrNotFound.
				WithError(entity.ErrCannotGetSummary.Error()).
				WithDebug(err.Error())
		}

		return core.ErrInternalServerError.
			WithError(entity.ErrCannotGetSummary.Error()).
			WithDebug(err.Error())
	}

	// requester := core.GetRequester(ctx)

	// uid, _ := core.FromBase58(requester.GetSubject())
	// requesterId := int(uid.GetLocalID())

	// // Only Summary user can do this
	// if requesterId != Summary.UserId {
	// 	return core.ErrForbidden.WithError(entity.ErrRequesterIsNotOwner.Error())
	// }

	if err := biz.summaryRepo.DeleteSummary(ctx, id); err != nil {
		return core.ErrInternalServerError.
			WithError(entity.ErrCannotDeleteSummary.Error()).
			WithDebug(err.Error())
	}

	return nil
}
