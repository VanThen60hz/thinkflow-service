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

	requester := core.GetRequester(ctx)

	uid, _ := core.FromBase58(requester.GetSubject())
	requesterId := int(uid.GetLocalID())

	hasReadPermission, err := biz.collabRepo.HasReadPermission(ctx, int(data.NoteID), requesterId)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotGetText.Error()).
			WithDebug(err.Error())
	}
	note, err := biz.noteRepo.GetNoteById(ctx, int(data.NoteID))
	if err != nil {
		if err == core.ErrRecordNotFound {
			return nil, core.ErrNotFound.
				WithError(entity.ErrCannotGetText.Error()).
				WithDebug(err.Error())
		}

		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotGetText.Error()).
			WithDebug(err.Error())
	}

	if note.UserId != requesterId && !hasReadPermission {
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrRequesterCannotRead.Error())
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
