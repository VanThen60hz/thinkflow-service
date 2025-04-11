package business

import (
	"context"

	"thinkflow-service/services/note/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) UpdateNote(ctx context.Context, id int, data *entity.NoteDataUpdate) error {
	// Get note data, without extra infos
	note, err := biz.noteRepo.GetNoteById(ctx, id)
	if err != nil {
		if err == core.ErrRecordNotFound {
			return core.ErrNotFound.
				WithError(entity.ErrCannotGetNote.Error()).
				WithDebug(err.Error())
		}

		return core.ErrInternalServerError.
			WithError(entity.ErrCannotGetNote.Error()).
			WithDebug(err.Error())
	}

	requester := core.GetRequester(ctx)

	uid, _ := core.FromBase58(requester.GetSubject())
	requesterId := int(uid.GetLocalID())

	hasWritePermission, err := biz.collabRepo.HasWritePermission(ctx, id, requesterId)
	if err != nil {
		return core.ErrInternalServerError.
			WithError(entity.ErrCannotGetNote.Error()).
			WithDebug(err.Error())
	}

	// Only note user or collab user can do this
	if requesterId != note.UserId && !hasWritePermission {
		return core.ErrForbidden.WithError(entity.ErrRequesterCannotModify.Error())
	}

	if err := biz.noteRepo.UpdateNote(ctx, id, data); err != nil {
		return core.ErrInternalServerError.
			WithError(entity.ErrCannotUpdateNote.Error()).
			WithDebug(err.Error())
	}

	return nil
}
