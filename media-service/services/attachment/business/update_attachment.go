package business

import (
	"context"
	"thinkflow-service/services/attachment/entity"
)

func (biz *business) UpdateAttachment(ctx context.Context, id int64, data *entity.Attachment) error {
	return biz.attachmentRepo.UpdateAttachment(ctx, int(id), data)
}