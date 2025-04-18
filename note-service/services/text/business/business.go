package business

import (
	"context"

	"thinkflow-service/common"
	textEntity "thinkflow-service/services/text/entity"

	collaborationEntity "thinkflow-service/services/collaboration/entity"
	noteEntity "thinkflow-service/services/note/entity"

	"github.com/VanThen60hz/service-context/core"
)

type TextRepository interface {
	AddNewText(ctx context.Context, data *textEntity.TextDataCreation) error
	UpdateText(ctx context.Context, id int, data *textEntity.TextDataUpdate) error
	DeleteText(ctx context.Context, id int) error
	GetTextById(ctx context.Context, id int) (*textEntity.Text, error)
	GetTextByNoteId(ctx context.Context, noteId int) (*textEntity.Text, error)
}

type SummaryRepository interface {
	GetSummaryById(ctx context.Context, id int64) (*common.SimpleSummary, error)
}

type MindmapRepository interface {
	GetMindmapById(ctx context.Context, id int64) (*common.SimpleMindmap, error)
}

type NoteRepository interface {
	AddNewNote(ctx context.Context, data *noteEntity.NoteDataCreation) error
	GetNoteById(ctx context.Context, id int) (*noteEntity.Note, error)
	ListNotes(ctx context.Context, filter *noteEntity.Filter, paging *core.Paging) ([]noteEntity.Note, error)
	ListArchivedNotes(ctx context.Context, filter *noteEntity.Filter, paging *core.Paging) ([]noteEntity.Note, error)
	UpdateNote(ctx context.Context, id int, data *noteEntity.NoteDataUpdate) error
	ArchiveNote(ctx context.Context, id int) error
	UnarchiveNote(ctx context.Context, id int) error
	DeleteNote(ctx context.Context, id int) error
}

type CollaborationRepository interface {
	AddNewCollaboration(ctx context.Context, data *collaborationEntity.CollaborationCreation) error
	HasReadPermission(ctx context.Context, noteId int, userId int) (bool, error)
	HasWritePermission(ctx context.Context, noteId int, userId int) (bool, error)
	GetCollaborationByNoteId(ctx context.Context, noteId int, paging *core.Paging) ([]collaborationEntity.Collaboration, error)
	GetCollaborationByUserId(ctx context.Context, userId int, paging *core.Paging) ([]collaborationEntity.Collaboration, error)
}

type business struct {
	textRepo    TextRepository
	noteRepo    NoteRepository
	collabRepo  CollaborationRepository
	summaryRepo SummaryRepository
	mindmapRepo MindmapRepository
}

func NewBusiness(
	textRepo TextRepository,
	noteRepo NoteRepository,
	collabRepo CollaborationRepository,
	summaryRepo SummaryRepository,
	mindmapRepo MindmapRepository,
) *business {
	return &business{
		textRepo:    textRepo,
		noteRepo:    noteRepo,
		collabRepo:  collabRepo,
		summaryRepo: summaryRepo,
		mindmapRepo: mindmapRepo,
	}
}
