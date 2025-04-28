package business

import (
	"context"

	"thinkflow-service/services/text/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) UpdateText(ctx context.Context, id int, data *entity.TextDataUpdate) error {
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

	hasWritePermission, err := biz.collabRepo.HasWritePermission(ctx, int(text.NoteID), requesterId)
	if err != nil {
		return core.ErrInternalServerError.
			WithError(entity.ErrCannotGetText.Error()).
			WithDebug(err.Error())
	}

	if requesterId != note.UserId && !hasWritePermission {
		return core.ErrForbidden.WithError(entity.ErrRequesterCannotModify.Error())
	}

	if err := biz.textRepo.UpdateText(ctx, id, data); err != nil {
		return core.ErrInternalServerError.
			WithError(entity.ErrCannotUpdateText.Error()).
			WithDebug(err.Error())
	}

	return nil
}
