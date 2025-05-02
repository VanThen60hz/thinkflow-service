package business

import (
	"context"

	"thinkflow-service/common"
	"thinkflow-service/proto/pb"
	noteEntity "thinkflow-service/services/note/entity"
	textEntity "thinkflow-service/services/text/entity"

	"github.com/VanThen60hz/service-context/component/emailc"
	"github.com/VanThen60hz/service-context/component/redisc"
	"github.com/VanThen60hz/service-context/component/s3c"
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
	DeleteUserNotes(ctx context.Context, userId int) (int, error)
}

type SummaryRepository interface {
	GetSummaryById(ctx context.Context, id int64) (*common.SimpleSummary, error)
	CreateSummary(ctx context.Context, summaryText string) (int64, error)
}

type MindmapRepository interface {
	GetMindmapById(ctx context.Context, id int64) (*common.SimpleMindmap, error)
	CreateMindmap(ctx context.Context, mindmapText string) (int64, error)
}

type UserRepository interface {
	GetUsersByIds(ctx context.Context, ids []int) ([]core.SimpleUser, error)
	GetUserById(ctx context.Context, id int) (*core.SimpleUser, error)
}

type CollaborationRepository interface {
	AddNewCollaboration(ctx context.Context, data *pb.CollaborationCreation) error
	HasReadPermission(ctx context.Context, noteId int, userId int) (bool, error)
	HasWritePermission(ctx context.Context, noteId int, userId int) (bool, error)
	GetCollaborationByNoteIdAndUserId(ctx context.Context, noteId int, userId int) (*pb.Collaboration, error)
	GetCollaborationByNoteId(ctx context.Context, noteId int, paging *core.Paging) ([]*pb.Collaboration, error)
	GetCollaborationByUserId(ctx context.Context, userId int, paging *core.Paging) ([]*pb.Collaboration, error)
	UpdateCollaboration(ctx context.Context, id int, data *pb.CollaborationUpdate) error
	DeleteCollaboration(ctx context.Context, id int) error
	RemoveCollaborationByNoteIdAndUserId(ctx context.Context, noteId int, userId int) error
}

type NoteShareLinkRepository interface {
	AddNewNoteShareLink(ctx context.Context, data *pb.NoteShareLinkCreation) error
	GetNoteShareLinkByID(ctx context.Context, id int64) (*pb.NoteShareLink, error)
	GetNoteShareLinkByToken(ctx context.Context, token string) (*pb.NoteShareLink, error)
	UpdateNoteShareLink(ctx context.Context, id int64, data *pb.NoteShareLinkUpdate) error
	DeleteNoteShareLink(ctx context.Context, id int64) error
}

type TextRepository interface {
	AddNewText(ctx context.Context, data *textEntity.TextDataCreation) error
	UpdateText(ctx context.Context, id int, data *textEntity.TextDataUpdate) error
	DeleteText(ctx context.Context, id int) error
	GetTextById(ctx context.Context, id int) (*textEntity.Text, error)
	GetTextByNoteId(ctx context.Context, noteId int) (*textEntity.Text, error)
}

type AudioRepository interface {
	GetAudioById(ctx context.Context, id int64) (*pb.PublicAudioInfo, error)
	GetAudiosByNoteId(ctx context.Context, noteId int64) ([]*pb.PublicAudioInfo, error)
	DeleteAudio(ctx context.Context, id int64) error
}

type TranscriptRepository interface {
	GetTranscriptById(ctx context.Context, id int64) (*common.SimpleTranscript, error)
	CreateTranscript(ctx context.Context, content string) (int64, error)
}

type business struct {
	noteRepo          NoteRepository
	textRepo          TextRepository
	userRepo          UserRepository
	audioRepo         AudioRepository
	collabRepo        CollaborationRepository
	noteShareLinkRepo NoteShareLinkRepository
	transcriptRepo    TranscriptRepository
	summaryRepo       SummaryRepository
	mindmapRepo       MindmapRepository
	jwtProvider       common.JWTProvider
	s3Client          *s3c.S3Component
	redisClient       redisc.Redis
	emailService      emailc.Email
}

func NewBusiness(
	noteRepo NoteRepository,
	textRepo TextRepository,
	userRepo UserRepository,
	audioRepo AudioRepository,
	collabRepo CollaborationRepository,
	noteShareLinkRepo NoteShareLinkRepository,
	transcriptRepo TranscriptRepository,
	summaryRepo SummaryRepository,
	mindmapRepo MindmapRepository,
	jwtProvider common.JWTProvider,
	s3Client *s3c.S3Component,
	redisClient redisc.Redis,
	emailService emailc.Email,
) *business {
	return &business{
		noteRepo:          noteRepo,
		textRepo:          textRepo,
		userRepo:          userRepo,
		audioRepo:         audioRepo,
		collabRepo:        collabRepo,
		noteShareLinkRepo: noteShareLinkRepo,
		transcriptRepo:    transcriptRepo,
		summaryRepo:       summaryRepo,
		mindmapRepo:       mindmapRepo,
		jwtProvider:       jwtProvider,
		s3Client:          s3Client,
		redisClient:       redisClient,
		emailService:      emailService,
	}
}
