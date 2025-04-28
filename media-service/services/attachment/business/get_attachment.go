package business

import (
	"context"
	"thinkflow-service/services/attachment/entity"
)

func (biz *business) GetAttachment(ctx context.Context, id int64) (*entity.Attachment, error) {
	return biz.attachmentRepo.GetByID(ctx, id)
}