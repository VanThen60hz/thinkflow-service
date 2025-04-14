package business

import (
	"context"

	"thinkflow-service/common"
	collaborationEntity "thinkflow-service/services/collaboration/entity"
	noteShareLinkEntity "thinkflow-service/services/note-share-links/entity"
	noteEntity "thinkflow-service/services/note/entity"

	"github.com/VanThen60hz/service-context/component/emailc"
	"github.com/VanThen60hz/service-context/component/redisc"
	"github.com/VanThen60hz/service-context/core"
)

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

type UserRepository interface {
	GetUsersByIds(ctx context.Context, ids []int) ([]core.SimpleUser, error)
	GetUserById(ctx context.Context, id int) (*core.SimpleUser, error)
}

type CollaborationRepository interface {
	AddNewCollaboration(ctx context.Context, data *collaborationEntity.CollaborationCreation) error
	HasReadPermission(ctx context.Context, noteId int, userId int) (bool, error)
	HasWritePermission(ctx context.Context, noteId int, userId int) (bool, error)
	GetCollaborationByNoteId(ctx context.Context, noteId int, paging *core.Paging) ([]collaborationEntity.Collaboration, error)
	GetCollaborationByUserId(ctx context.Context, userId int, paging *core.Paging) ([]collaborationEntity.Collaboration, error)
	RemoveCollaborationByNoteIdAndUserId(ctx context.Context, noteId int, userId int) error
}

type NoteShareLinkRepository interface {
	AddNewNoteShareLink(ctx context.Context, data *noteShareLinkEntity.NoteShareLinkCreation) error
	GetNoteShareLinkByID(ctx context.Context, id int64) (*noteShareLinkEntity.NoteShareLink, error)
	GetNoteShareLinkByToken(ctx context.Context, token string) (*noteShareLinkEntity.NoteShareLink, error)
	UpdateNoteShareLink(ctx context.Context, id int64, data *noteShareLinkEntity.NoteShareLinkUpdate) error
	DeleteNoteShareLink(ctx context.Context, id int64) error
}

type business struct {
	noteRepo          NoteRepository
	userRepo          UserRepository
	collabRepo        CollaborationRepository
	noteShareLinkRepo NoteShareLinkRepository
	jwtProvider       common.JWTProvider
	redisClient       redisc.Redis
	emailService      emailc.Email
}

func NewBusiness(
	noteRepo NoteRepository,
	userRepo UserRepository,
	collabRepo CollaborationRepository,
	noteShareLinkRepo NoteShareLinkRepository,
	jwtProvider common.JWTProvider,
	redisClient redisc.Redis,
	emailService emailc.Email,
) *business {
	return &business{
		noteRepo:          noteRepo,
		userRepo:          userRepo,
		collabRepo:        collabRepo,
		noteShareLinkRepo: noteShareLinkRepo,
		jwtProvider:       jwtProvider,
		redisClient:       redisClient,
		emailService:      emailService,
	}
}
