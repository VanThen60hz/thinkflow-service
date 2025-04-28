package business

import (
	"context"

	"thinkflow-service/services/attachment/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) UpdateAttachment(ctx context.Context, id int64, data *entity.Attachment) error {
	attachment, err := biz.attachmentRepo.GetByID(ctx, id)
	if err != nil {
		if err == core.ErrRecordNotFound {
			return core.ErrNotFound.
				WithDebug(err.Error())
		}

		return core.ErrInternalServerError.
			WithError(entity.ErrCannotGetAttachment.Error()).
			WithDebug(err.Error())
	}

	requester := core.GetRequester(ctx)
	uid, _ := core.FromBase58(requester.GetSubject())
	requesterId := int(uid.GetLocalID())

	hasWritePermission, err := biz.collabRepo.HasWritePermission(ctx, int(attachment.NoteID), requesterId)
	if err != nil {
		return core.ErrInternalServerError.
			WithError(entity.ErrCannotGetAttachment.Error()).
			WithDebug(err.Error())
	}
	note, err := biz.noteRepo.GetNoteById(ctx, int(attachment.NoteID))
	if err != nil {
		if err == core.ErrRecordNotFound {
			return core.ErrNotFound.
				WithError(entity.ErrCannotGetAttachment.Error()).
				WithDebug(err.Error())
		}

		return core.ErrInternalServerError.
			WithError(entity.ErrCannotGetAttachment.Error()).
			WithDebug(err.Error())
	}

	if note.UserId != int64(requesterId) && !hasWritePermission {
		return core.ErrInternalServerError.
			WithError(entity.ErrRequesterCannotModify.Error())
	}

	return biz.attachmentRepo.UpdateAttachment(ctx, int(id), data)
}
