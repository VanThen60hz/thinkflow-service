package business

import (
	"context"

	collaborationEntity "thinkflow-service/services/collaboration/entity"

	"github.com/VanThen60hz/service-context/core"
)

type CollaborationRepository interface {
	AddNewCollaboration(ctx context.Context, data *collaborationEntity.CollaborationCreation) error
	HasReadPermission(ctx context.Context, noteId int, userId int) (bool, error)
	HasWritePermission(ctx context.Context, noteId int, userId int) (bool, error)
	GetCollaborationByNoteIdAndUserId(ctx context.Context, noteId int, userId int) (*collaborationEntity.Collaboration, error)
	GetCollaborationByNoteId(ctx context.Context, noteId int, paging *core.Paging) ([]collaborationEntity.Collaboration, error)
	GetCollaborationByUserId(ctx context.Context, userId int, paging *core.Paging) ([]collaborationEntity.Collaboration, error)
	UpdateCollaboration(ctx context.Context, id int, data *collaborationEntity.Collaboration) error
	DeleteCollaboration(ctx context.Context, id int) error
	RemoveCollaborationByNoteIdAndUserId(ctx context.Context, noteId int, userId int) error
}

type business struct {
	collabRepo CollaborationRepository
}

func NewBusiness(
	collabRepo CollaborationRepository,
) *business {
	return &business{
		collabRepo: collabRepo,
	}
}
