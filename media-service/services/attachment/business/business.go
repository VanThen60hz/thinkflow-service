package business

import (
	"context"

	"thinkflow-service/services/attachment/entity"

	"github.com/VanThen60hz/service-context/component/s3c"
)

type AttachmentRepository interface {
	GetByID(ctx context.Context, id int64) (*entity.Attachment, error)
	GetByNoteID(ctx context.Context, noteID int64) ([]entity.Attachment, error)
	AddNewAttachment(ctx context.Context, data *entity.AttachmentCreation) error
	UpdateAttachment(ctx context.Context, id int, data *entity.Attachment) error
	DeleteAttachment(ctx context.Context, id int) error
}

type NoteRepository interface {
	DeleteUserNotes(ctx context.Context, userId int32) (bool, int32, error)
}

type business struct {
	attachmentRepo AttachmentRepository
	noteRepo       NoteRepository
	s3Client       *s3c.S3Component
}

func NewBusiness(attachRepo AttachmentRepository, noteRepo NoteRepository, s3Client *s3c.S3Component) *business {
	return &business{
		attachmentRepo: attachRepo,
		noteRepo:       noteRepo,
		s3Client:       s3Client,
	}
}
