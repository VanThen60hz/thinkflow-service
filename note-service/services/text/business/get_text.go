package business

import (
	"context"

	"thinkflow-service/services/text/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) GetTextById(ctx context.Context, id int) (*entity.Text, error) {
	data, err := biz.textRepo.GetTextById(ctx, id)
	if err != nil {
		if err == core.ErrRecordNotFound {
			return nil, core.ErrNotFound.
				WithDebug(err.Error())
		}

		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotGetText.Error()).
			WithDebug(err.Error())
	}

	if data.SummaryID != nil {
		summary, err := biz.summaryRepo.GetSummaryById(ctx, *data.SummaryID)
		if err != nil {
			return nil, core.ErrInternalServerError.
				WithError(entity.ErrCannotGetSummary.Error()).
				WithDebug(err.Error())
		}
		data.Summary = summary
	}

	return data, nil
}
