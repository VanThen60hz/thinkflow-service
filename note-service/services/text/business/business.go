package business

import (
	"context"

	"thinkflow-service/common"
	"thinkflow-service/proto/pb"
	textEntity "thinkflow-service/services/text/entity"

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
	CreateSummary(ctx context.Context, summaryText string) (int64, error)
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
	AddNewCollaboration(ctx context.Context, data *pb.CollaborationCreation) error
	HasReadPermission(ctx context.Context, noteId int, userId int) (bool, error)
	HasWritePermission(ctx context.Context, noteId int, userId int) (bool, error)
	GetCollaborationByNoteId(ctx context.Context, noteId int, paging *core.Paging) ([]*pb.Collaboration, error)
	GetCollaborationByUserId(ctx context.Context, userId int, paging *core.Paging) ([]*pb.Collaboration, error)
}

type NotificationRepository interface {
	CreateNotification(ctx context.Context, notiType string, senderId, receiverId int64, content string, options *string) error
}

type Business interface {
	GetTextById(ctx context.Context, id int) (*textEntity.Text, error)
	GetTextByNoteId(ctx context.Context, noteId int) (*textEntity.Text, error)
	CreateNewText(ctx context.Context, data *textEntity.TextDataCreation) error
	UpdateText(ctx context.Context, id int, data *textEntity.TextDataUpdate) error
	DeleteText(ctx context.Context, id int) error
	SummaryText(ctx context.Context, textID int) (*SummaryResponse, error)
}

type business struct {
	textRepo    TextRepository
	noteRepo    NoteRepository
	collabRepo  CollaborationRepository
	summaryRepo SummaryRepository
	notiRepo    NotificationRepository
}

func NewBusiness(
	textRepo TextRepository,
	noteRepo NoteRepository,
	collabRepo CollaborationRepository,
	summaryRepo SummaryRepository,
	notiRepo NotificationRepository,
) *business {
	return &business{
		textRepo:    textRepo,
		noteRepo:    noteRepo,
		collabRepo:  collabRepo,
		summaryRepo: summaryRepo,
		notiRepo:    notiRepo,
	}
}
