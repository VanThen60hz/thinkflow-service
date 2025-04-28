package business

import (
	"context"

	"thinkflow-service/services/attachment/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) GetAttachmentsByNoteID(ctx context.Context, noteID int64) ([]entity.Attachment, error) {
	attachments, err := biz.attachmentRepo.GetByNoteID(ctx, noteID)
	if err != nil {
		if err == core.ErrRecordNotFound {
			return nil, core.ErrNotFound.
				WithDebug(err.Error())
		}

		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotListAttachments.Error()).
			WithDebug(err.Error())
	}

	requester := core.GetRequester(ctx)

	uid, _ := core.FromBase58(requester.GetSubject())
	requesterId := int(uid.GetLocalID())

	hasReadPermission, err := biz.collabRepo.HasReadPermission(ctx, int(noteID), requesterId)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotGetAttachment.Error()).
			WithDebug(err.Error())
	}
	note, err := biz.noteRepo.GetNoteById(ctx, int(noteID))
	if err != nil {
		if err == core.ErrRecordNotFound {
			return nil, core.ErrNotFound.
				WithError(entity.ErrCannotListAttachments.Error()).
				WithDebug(err.Error())
		}

		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotListAttachments.Error()).
			WithDebug(err.Error())
	}

	if note.UserId != int64(requesterId) && !hasReadPermission {
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrRequesterCannotRead.Error())
	}

	return attachments, nil
}
