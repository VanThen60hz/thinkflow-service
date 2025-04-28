package business

import (
	"context"
	"fmt"
	"strings"

	"thinkflow-service/services/attachment/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) DeleteAttachment(ctx context.Context, id int64) error {
	attachment, err := biz.attachmentRepo.GetByID(ctx, id)
	if err != nil {
		return core.ErrInternalServerError.WithError(err.Error())
	}

	requester := core.GetRequester(ctx)
	uid, _ := core.FromBase58(requester.GetSubject())
	requesterId := int(uid.GetLocalID())

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

	if note.UserId != int64(requesterId) {
		return core.ErrInternalServerError.
			WithError(entity.ErrRequesterIsNotAttachmentOwner.Error())
	}

	if err := biz.attachmentRepo.DeleteAttachment(ctx, int(id)); err != nil {
		return core.ErrInternalServerError.WithError(err.Error())
	}

	urlParts := strings.Split(attachment.FileURL, "/attachments/")
	attachmentId := urlParts[len(urlParts)-1]
	fileKey := fmt.Sprintf("attachments/%s", attachmentId)
	if err := biz.s3Client.DeleteObject(ctx, fileKey); err != nil {
		fmt.Printf("Failed to delete file from S3: %v\n", err)
	}

	return nil
}
