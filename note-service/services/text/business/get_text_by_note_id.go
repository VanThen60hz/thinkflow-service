package business

import (
	"context"

	"thinkflow-service/services/text/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) GetTextByNoteId(ctx context.Context, noteId int) (*entity.Text, error) {
	text, err := biz.textRepo.GetTextByNoteId(ctx, noteId)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotListText.Error()).
			WithDebug(err.Error())
	}

	requester := core.GetRequester(ctx)
	uid, _ := core.FromBase58(requester.GetSubject())
	requesterId := int(uid.GetLocalID())

	hasReadPermission, err := biz.collabRepo.HasReadPermission(ctx, int(text.NoteID), requesterId)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotGetText.Error()).
			WithDebug(err.Error())
	}

	note, err := biz.noteRepo.GetNoteById(ctx, int(text.NoteID))
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

	if text.SummaryID != nil {
		summary, err := biz.summaryRepo.GetSummaryById(ctx, *text.SummaryID)
		if err != nil {
			return nil, core.ErrInternalServerError.
				WithError(entity.ErrCannotGetSummary.Error()).
				WithDebug(err.Error())
		}
		text.Summary = summary
	}

	return text, nil
}
