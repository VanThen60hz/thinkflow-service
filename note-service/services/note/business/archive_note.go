package business

import (
	"context"

	"thinkflow-service/common"
	"thinkflow-service/services/note/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) ArchiveNote(ctx context.Context, id int) error {
	data, err := biz.noteRepo.GetNoteById(ctx, id)
	if err != nil {
		if err == core.ErrRecordNotFound {
			return core.ErrNotFound.
				WithDebug(err.Error())
		}

		return core.ErrInternalServerError.
			WithError(entity.ErrCannotGetNote.Error()).
			WithDebug(err.Error())
	}

	requesterVal := ctx.Value(common.RequesterKey)
	if requesterVal == nil {
		return core.ErrUnauthorized.WithError("requester not found in context")
	}

	requester, ok := requesterVal.(core.Requester)
	if !ok {
		return core.ErrInternalServerError.
			WithError("invalid requester type in context")
	}

	uid, _ := core.FromBase58(requester.GetSubject())
	requesterId := int(uid.GetLocalID())

	if requesterId != data.UserId {
		return core.ErrForbidden.WithError(entity.ErrRequesterIsNotOwner.Error())
	}

	if err := biz.noteRepo.ArchiveNote(ctx, id); err != nil {
		return core.ErrInternalServerError.
			WithError(entity.ErrCannotArchiveNote.Error()).
			WithDebug(err.Error())
	}

	return nil
}
