package business

import (
	"context"

	"thinkflow-service/services/summary/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) GetSummaryById(ctx context.Context, id int) (*entity.Summary, error) {
	data, err := biz.summaryRepo.GetSummaryById(ctx, id)
	if err != nil {
		if err == core.ErrRecordNotFound {
			return nil, core.ErrNotFound.
				WithDebug(err.Error())
		}

		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotGetSummary.Error()).
			WithDebug(err.Error())
	}

	// requesterVal := ctx.Value(common.RequesterKey)
	// if requesterVal == nil {
	// 	return nil, core.ErrUnauthorized.WithError("requester not found in context")
	// }

	// requester, ok := requesterVal.(core.Requester)
	// if !ok {
	// 	return nil, core.ErrInternalServerError.
	// 		WithError("invalid requester type in context")
	// }

	// uid, _ := core.FromBase58(requester.GetSubject())
	// requesterId := int(uid.GetLocalID())

	// if requesterId != data.UserId {
	// 	return nil, core.ErrForbidden.WithError(entity.ErrRequesterIsNotOwner.Error())
	// }

	// // Get extra infos: User
	// user, err := biz.userRepo.GetUserById(ctx, data.UserId)
	// if err != nil {
	// 	return nil, core.ErrInternalServerError.
	// 		WithError(entity.ErrCannotGetSummary.Error()).
	// 		WithDebug(err.Error())
	// }

	// data.User = user

	return data, nil
}
