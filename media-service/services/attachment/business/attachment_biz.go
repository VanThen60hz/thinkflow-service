package business

import (
	"context"

	"thinkflow-service/services/attachment/entity"
)

type AttachmentRepo interface {
	GetByID(ctx context.Context, id int64) (*entity.Attachment, error)
	GetByNoteID(ctx context.Context, noteID int64) ([]entity.Attachment, error)
	AddNewAttachment(ctx context.Context, data *entity.AttachmentCreation) error
	UpdateAttachment(ctx context.Context, id int, data *entity.Attachment) error
	DeleteAttachment(ctx context.Context, id int) error
}

type business struct {
	attachmentRepo AttachmentRepo
}

func NewBusiness(repo AttachmentRepo) *business {
	return &business{
		attachmentRepo: repo,
	}
}

func (biz *business) CreateAttachment(ctx context.Context, data *entity.AttachmentCreation) error {
	return biz.attachmentRepo.AddNewAttachment(ctx, data)
}

func (biz *business) GetAttachment(ctx context.Context, id int64) (*entity.Attachment, error) {
	return biz.attachmentRepo.GetByID(ctx, id)
}

func (biz *business) GetAttachmentsByNoteID(ctx context.Context, noteID int64) ([]entity.Attachment, error) {
	return biz.attachmentRepo.GetByNoteID(ctx, noteID)
}

func (biz *business) UpdateAttachment(ctx context.Context, id int64, data *entity.Attachment) error {
	return biz.attachmentRepo.UpdateAttachment(ctx, int(id), data)
}

func (biz *business) DeleteAttachment(ctx context.Context, id int64) error {
	return biz.attachmentRepo.DeleteAttachment(ctx, int(id))
}
