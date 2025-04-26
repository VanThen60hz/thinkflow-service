package business

import (
	"context"

	"thinkflow-service/common"
	"thinkflow-service/services/audio/entity"

	"github.com/VanThen60hz/service-context/component/s3c"
	"github.com/VanThen60hz/service-context/core"
)

type AudioRepository interface {
	AddNewAudio(ctx context.Context, data *entity.AudioDataCreation) error
	UpdateAudio(ctx context.Context, id int, data *entity.AudioDataUpdate) error
	DeleteAudio(ctx context.Context, id int) error
	GetAudioById(ctx context.Context, id int) (*entity.Audio, error)
	ListAudios(ctx context.Context, filter *entity.Filter, paging *core.Paging) ([]entity.Audio, error)
	GetAudiosByNoteId(ctx context.Context, noteId int) ([]entity.Audio, error)
}

type TranscriptRepository interface {
	GetTranscriptById(ctx context.Context, id int64) (*common.SimpleTranscript, error)
}

type SummaryRepository interface {
	GetSummaryById(ctx context.Context, id int64) (*common.SimpleSummary, error)
}

type business struct {
	audioRepo      AudioRepository
	s3Client       *s3c.S3Component
	transcriptRepo TranscriptRepository
	summaryRepo    SummaryRepository
}

func NewBusiness(
	audioRepo AudioRepository,
	s3Client *s3c.S3Component,
	transcriptRepo TranscriptRepository,
	summaryRepo SummaryRepository,
) *business {
	return &business{
		audioRepo:      audioRepo,
		s3Client:       s3Client,
		transcriptRepo: transcriptRepo,
		summaryRepo:    summaryRepo,
	}
}
