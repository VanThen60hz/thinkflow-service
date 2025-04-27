package rpc

import (
	"context"

	"thinkflow-service/services/collaboration/entity"

	"github.com/VanThen60hz/service-context/core"
)

type Business interface {
	AddNewCollaboration(ctx context.Context, data *entity.CollaborationCreation) error
	HasReadPermission(ctx context.Context, noteId int, userId int) (bool, error)
	HasWritePermission(ctx context.Context, noteId int, userId int) (bool, error)
	GetCollaborationByNoteIdAndUserId(ctx context.Context, noteId int, userId int) (*entity.Collaboration, error)
	GetCollaborationByNoteId(ctx context.Context, noteId int, paging *core.Paging) ([]entity.Collaboration, error)
	GetCollaborationByUserId(ctx context.Context, userId int, paging *core.Paging) ([]entity.Collaboration, error)
	UpdateCollaboration(ctx context.Context, id int, data *entity.Collaboration) error
	DeleteCollaboration(ctx context.Context, id int) error
	RemoveCollaborationByNoteIdAndUserId(ctx context.Context, noteId int, userId int) error
}

type grpcService struct {
	business Business
}

func NewService(business Business) *grpcService {
	return &grpcService{business: business}
}
