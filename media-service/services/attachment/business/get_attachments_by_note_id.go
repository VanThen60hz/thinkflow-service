package business

import (
	"context"

	"thinkflow-service/services/attachment/entity"
)

func (biz *business) GetAttachmentsByNoteID(ctx context.Context, noteID int64) ([]entity.Attachment, error) {
	return biz.attachmentRepo.GetByNoteID(ctx, noteID)
}
