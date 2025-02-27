package business

import (
	"context"

	"thinkflow-service/services/note/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) DeleteNote(ctx context.Context, id int) error {
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

	// Only note user can do this
	if requesterId != note.UserId {
		return core.ErrForbidden.WithError(entity.ErrRequesterIsNotOwner.Error())
	}

	// Only delete note with doing status
	// if note.Status == entity.StatusDeleted {
	// 	return core.ErrForbidden.
	// 		WithError(entity.ErrNoteDeleted.Error())
	// }

	if err := biz.noteRepo.DeleteNote(ctx, id); err != nil {
		return core.ErrInternalServerError.
			WithError(entity.ErrCannotDeleteNote.Error()).
			WithDebug(err.Error())
	}

	return nil
}
