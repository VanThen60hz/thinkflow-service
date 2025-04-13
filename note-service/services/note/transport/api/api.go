package api

import (
	"context"

	noteShareLinkEntity "thinkflow-service/services/note-share-links/entity"
	noteEntity "thinkflow-service/services/note/entity"

	sctx "github.com/VanThen60hz/service-context"
	"github.com/VanThen60hz/service-context/core"
)

type ServiceContext interface {
	sctx.ServiceContext
	Business
}

type Business interface {
	CreateNewNote(ctx context.Context, data *noteEntity.NoteDataCreation) error
	CreateNoteShareLink(ctx context.Context, noteId int64, permission string) (*noteShareLinkEntity.NoteShareLink, error)
	GetNoteById(ctx context.Context, id int) (*noteEntity.Note, error)
	ListNoteMembersById(ctx context.Context, id int, paging *core.Paging) ([]noteEntity.NoteMember, error)
	ListNotes(ctx context.Context, filter *noteEntity.Filter, paging *core.Paging) ([]noteEntity.Note, error)
	ListNotesSharedWithMe(ctx context.Context, filter *noteEntity.Filter, paging *core.Paging) ([]noteEntity.Note, error)
	ListArchivedNotes(ctx context.Context, filter *noteEntity.Filter, paging *core.Paging) ([]noteEntity.Note, error)
	AcceptSharedNote(ctx context.Context, token string) (int, error)
	UpdateNote(ctx context.Context, id int, data *noteEntity.NoteDataUpdate) error
	ArchiveNote(ctx context.Context, id int) error
	UnarchiveNote(ctx context.Context, id int) error
	DeleteNote(ctx context.Context, id int) error
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
