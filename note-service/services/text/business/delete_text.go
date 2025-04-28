package business

import (
	"context"

	"thinkflow-service/services/text/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) DeleteText(ctx context.Context, id int) error {
	text, err := biz.textRepo.GetTextById(ctx, id)
	if err != nil {
		if err == core.ErrRecordNotFound {
			return core.ErrNotFound.
				WithError(entity.ErrCannotGetText.Error()).
				WithDebug(err.Error())
		}

		return core.ErrInternalServerError.
			WithError(entity.ErrCannotGetText.Error()).
			WithDebug(err.Error())
	}

	requester := core.GetRequester(ctx)
	uid, _ := core.FromBase58(requester.GetSubject())
	requesterId := int(uid.GetLocalID())

	note, err := biz.noteRepo.GetNoteById(ctx, int(text.NoteID))
	if err != nil {
		if err == core.ErrRecordNotFound {
			return core.ErrNotFound.
				WithError(entity.ErrCannotGetText.Error()).
				WithDebug(err.Error())
		}

		return core.ErrInternalServerError.
			WithError(entity.ErrCannotGetText.Error()).
			WithDebug(err.Error())
	}

	if requesterId != note.UserId {
		return core.ErrForbidden.WithError(entity.ErrRequesterIsNotOwner.Error())
	}

	if err := biz.textRepo.DeleteText(ctx, id); err != nil {
		return core.ErrInternalServerError.
			WithError(entity.ErrCannotDeleteText.Error()).
			WithDebug(err.Error())
	}

	return nil
}
