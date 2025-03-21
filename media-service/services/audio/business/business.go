package business

import (
	"context"

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

type business struct {
	audioRepo AudioRepository
}

func NewBusiness(audioRepo AudioRepository) *business {
	return &business{
		audioRepo: audioRepo,
	}
}
