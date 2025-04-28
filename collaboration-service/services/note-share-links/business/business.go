package business

import (
	"context"

	"thinkflow-service/services/note-share-links/entity"
)

type NoteShareLinkRepo interface {
	AddNewNoteShareLink(ctx context.Context, data *entity.NoteShareLinkCreation) error
	GetNoteShareLinkByID(ctx context.Context, id int64) (*entity.NoteShareLink, error)
	GetNoteShareLinkByToken(ctx context.Context, token string) (*entity.NoteShareLink, error)
	UpdateNoteShareLink(ctx context.Context, id int64, data *entity.NoteShareLinkUpdate) error
	DeleteNoteShareLink(ctx context.Context, id int64) error
}

type Business interface {
	CreateNoteShareLink(ctx context.Context, data *entity.NoteShareLinkCreation) (*entity.NoteShareLink, error)
	GetNoteShareLinkByID(ctx context.Context, id int64) (*entity.NoteShareLink, error)
	GetNoteShareLinkByToken(ctx context.Context, token string) (*entity.NoteShareLink, error)
	UpdateNoteShareLink(ctx context.Context, id int64, data *entity.NoteShareLinkUpdate) (*entity.NoteShareLink, error)
	DeleteNoteShareLink(ctx context.Context, id int64) error
}

type business struct {
	noteShareLinkRepo NoteShareLinkRepo
}

func NewBusiness(noteShareLinkRepo NoteShareLinkRepo) Business {
	return &business{
		noteShareLinkRepo: noteShareLinkRepo,
	}
}
