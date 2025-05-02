package api

import (
	"context"
	"time"

	"thinkflow-service/proto/pb"
	"thinkflow-service/services/note/business"
	noteEntity "thinkflow-service/services/note/entity"

	sctx "github.com/VanThen60hz/service-context"
	"github.com/VanThen60hz/service-context/core"
	"gorm.io/datatypes"
)

type ServiceContext interface {
	sctx.ServiceContext
	Business
}

type Business interface {
	CreateNewNote(ctx context.Context, data *noteEntity.NoteDataCreation) error
	CreateNoteShareLink(ctx context.Context, noteId int64, permission string, expiresAt *time.Time) (*pb.NoteShareLink, error)
	NoteShareLinkToEmail(ctx context.Context, noteId int64, email, permission string, expiresAt *time.Time) error
	AcceptSharedNote(ctx context.Context, token string) (*noteEntity.Note, error)
	SummaryNote(ctx context.Context, noteID int) (*business.SummaryNoteResponse, error)
	MindmapNote(ctx context.Context, noteID int) (datatypes.JSON, error)
	GetNoteById(ctx context.Context, id int) (*noteEntity.Note, error)
	ListNoteMembersByNoteId(ctx context.Context, id int, paging *core.Paging) ([]noteEntity.NoteMember, error)
	ListNotes(ctx context.Context, filter *noteEntity.Filter, paging *core.Paging) ([]noteEntity.Note, error)
	ListNotesSharedWithMe(ctx context.Context, filter *noteEntity.Filter, paging *core.Paging) ([]noteEntity.Note, error)
	ListArchivedNotes(ctx context.Context, filter *noteEntity.Filter, paging *core.Paging) ([]noteEntity.Note, error)
	UpdateNote(ctx context.Context, id int, data *noteEntity.NoteDataUpdate) error
	ArchiveNote(ctx context.Context, id int) error
	UnarchiveNote(ctx context.Context, id int) error
	UpdateNoteMemberAccess(ctx context.Context, noteId int, userId int, permission string) error
	DeleteNote(ctx context.Context, id int) error
	DeleteNoteMember(ctx context.Context, noteId int, userId int) error
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
