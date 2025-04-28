package api

import (
	"context"
	"mime/multipart"

	"thinkflow-service/services/attachment/entity"

	sctx "github.com/VanThen60hz/service-context"
)

type ServiceContext interface {
	sctx.ServiceContext
	Business
}

type Business interface {
	UploadAttachment(ctx context.Context, tempFile string, file *multipart.FileHeader, noteID int64) (*entity.AttachmentCreation, error)
	GetAttachment(ctx context.Context, id int64) (*entity.Attachment, error)
	GetAttachmentsByNoteID(ctx context.Context, noteID int64) ([]entity.Attachment, error)
	UpdateAttachment(ctx context.Context, id int64, data *entity.Attachment) error
	DeleteAttachment(ctx context.Context, id int64) error
}

type api struct {
	serviceCtx sctx.ServiceContext
	business   Business
}

func NewAPI(serviceCtx sctx.ServiceContext, business Business) *api {
	return &api{
		serviceCtx: serviceCtx,
		business:   business,
	}
}
