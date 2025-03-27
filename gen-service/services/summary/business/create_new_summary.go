package business

import (
	"context"

	"thinkflow-service/services/summary/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) CreateNewSummary(ctx context.Context, data *entity.SummaryDataCreation) error {
	requester := core.GetRequester(ctx)

	uid, _ := core.FromBase58(requester.GetSubject())
	requesterId := int(uid.GetLocalID()) // summary user id, id of who creates this new Summary

	data.Prepare(requesterId)

	if err := data.Validate(); err != nil {
		return core.ErrBadRequest.WithError(err.Error())
	}

	if err := biz.summaryRepo.AddNewSummary(ctx, data); err != nil {
		return core.ErrInternalServerError.WithError(entity.ErrCannotCreateSummary.Error())
	}

	return nil
}
