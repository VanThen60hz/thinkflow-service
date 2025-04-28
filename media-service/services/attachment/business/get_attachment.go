package business

import (
	"context"

	"thinkflow-service/services/attachment/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) GetAttachment(ctx context.Context, id int64) (*entity.Attachment, error) {
	attachment, err := biz.attachmentRepo.GetByID(ctx, id)
	if err != nil {
		if err == core.ErrRecordNotFound {
			return nil, core.ErrNotFound.
				WithDebug(err.Error())
		}

		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotGetAttachment.Error()).
			WithDebug(err.Error())
	}

	requester := core.GetRequester(ctx)
	uid, _ := core.FromBase58(requester.GetSubject())
	requesterId := int(uid.GetLocalID())

	hasReadPermission, err := biz.collabRepo.HasReadPermission(ctx, int(attachment.NoteID), requesterId)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotGetAttachment.Error()).
			WithDebug(err.Error())
	}
	note, err := biz.noteRepo.GetNoteById(ctx, int(attachment.NoteID))
	if err != nil {
		if err == core.ErrRecordNotFound {
			return nil, core.ErrNotFound.
				WithError(entity.ErrCannotGetAttachment.Error()).
				WithDebug(err.Error())
		}

		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotGetAttachment.Error()).
			WithDebug(err.Error())
	}

	if note.UserId != int64(requesterId) && !hasReadPermission {
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrRequesterCannotRead.Error())
	}

	return biz.attachmentRepo.GetByID(ctx, id)
}
