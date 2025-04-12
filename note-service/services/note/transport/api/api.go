package api

import (
	"context"

	"thinkflow-service/services/note/entity"

	sctx "github.com/VanThen60hz/service-context"
	"github.com/VanThen60hz/service-context/core"
)

type ServiceContext interface {
	sctx.ServiceContext
	Business
}

type Business interface {
	CreateNewNote(ctx context.Context, data *entity.NoteDataCreation) error
	GetNoteById(ctx context.Context, id int) (*entity.Note, error)
	ListNoteMembersById(ctx context.Context, id int, paging *core.Paging) ([]entity.NoteMember, error)
	ListNotes(ctx context.Context, filter *entity.Filter, paging *core.Paging) ([]entity.Note, error)
	ListNotesSharedWithMe(ctx context.Context, filter *entity.Filter, paging *core.Paging) ([]entity.Note, error)
	ListArchivedNotes(ctx context.Context, filter *entity.Filter, paging *core.Paging) ([]entity.Note, error)
	UpdateNote(ctx context.Context, id int, data *entity.NoteDataUpdate) error
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
