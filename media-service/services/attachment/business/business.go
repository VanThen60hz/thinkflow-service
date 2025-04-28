package business

import (
	"context"

	"thinkflow-service/proto/pb"
	"thinkflow-service/services/attachment/entity"

	"github.com/VanThen60hz/service-context/component/s3c"
	"github.com/VanThen60hz/service-context/core"
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
	GetNoteById(ctx context.Context, id int) (*pb.GetNoteByIdResp, error)
}

type CollaborationRepository interface {
	AddNewCollaboration(ctx context.Context, data *pb.CollaborationCreation) error
	HasReadPermission(ctx context.Context, noteId int, userId int) (bool, error)
	HasWritePermission(ctx context.Context, noteId int, userId int) (bool, error)
	GetCollaborationByNoteId(ctx context.Context, noteId int, paging *core.Paging) ([]*pb.Collaboration, error)
	GetCollaborationByUserId(ctx context.Context, userId int, paging *core.Paging) ([]*pb.Collaboration, error)
}

type business struct {
	attachmentRepo AttachmentRepository
	s3Client       *s3c.S3Component
	noteRepo       NoteRepository
	collabRepo     CollaborationRepository
}

func NewBusiness(attachRepo AttachmentRepository, s3Client *s3c.S3Component, noteRepo NoteRepository, collabRepo CollaborationRepository) *business {
	return &business{
		attachmentRepo: attachRepo,
		s3Client:       s3Client,
		noteRepo:       noteRepo,
		collabRepo:     collabRepo,
	}
}
