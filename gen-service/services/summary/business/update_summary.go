package business

import (
	"context"

	"thinkflow-service/services/summary/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) UpdateSummary(ctx context.Context, id int, data *entity.SummaryDataUpdate) error {
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

	if err := biz.summaryRepo.UpdateSummary(ctx, id, data); err != nil {
		return core.ErrInternalServerError.
			WithError(entity.ErrCannotUpdateSummary.Error()).
			WithDebug(err.Error())
	}

	return nil
}
