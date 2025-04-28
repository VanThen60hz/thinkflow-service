package rpc

import (
	"context"

	"thinkflow-service/services/note-share-links/entity"
)

type Business interface {
	CreateNoteShareLink(ctx context.Context, data *entity.NoteShareLinkCreation) (*entity.NoteShareLink, error)
	GetNoteShareLinkByID(ctx context.Context, id int64) (*entity.NoteShareLink, error)
	GetNoteShareLinkByToken(ctx context.Context, token string) (*entity.NoteShareLink, error)
	UpdateNoteShareLink(ctx context.Context, id int64, data *entity.NoteShareLinkUpdate) (*entity.NoteShareLink, error)
	DeleteNoteShareLink(ctx context.Context, id int64) error
}

type grpcService struct {
	business Business
}

func NewService(business Business) *grpcService {
	return &grpcService{business: business}
}
