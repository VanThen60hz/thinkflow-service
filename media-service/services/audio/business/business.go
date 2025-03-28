package business

import (
	"context"

	"thinkflow-service/common"
	"thinkflow-service/services/audio/entity"

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

type MindmapRepository interface {
	GetMindmapById(ctx context.Context, id int64) (*common.SimpleMindmap, error)
}

type business struct {
	audioRepo      AudioRepository
	transcriptRepo TranscriptRepository
	summaryRepo    SummaryRepository
	mindmapRepo    MindmapRepository
}

func NewBusiness(
	audioRepo AudioRepository,
	transcriptRepo TranscriptRepository,
	summaryRepo SummaryRepository,
	mindmapRepo MindmapRepository,
) *business {
	return &business{
		audioRepo:      audioRepo,
		transcriptRepo: transcriptRepo,
		summaryRepo:    summaryRepo,
		mindmapRepo:    mindmapRepo,
	}
}
