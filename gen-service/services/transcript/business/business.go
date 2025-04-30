package business

import (
	"context"

	"thinkflow-service/services/transcript/entity"

	"github.com/VanThen60hz/service-context/core"
)

type TranscriptRepository interface {
	AddNewTranscript(ctx context.Context, data *entity.TranscriptDataCreation) error
	UpdateTranscript(ctx context.Context, id int, data *entity.TranscriptDataUpdate) error
	DeleteTranscript(ctx context.Context, id int) error
	GetTranscriptById(ctx context.Context, id int) (*entity.Transcript, error)
}

type UserRepository interface {
	GetUsersByIds(ctx context.Context, ids []int) ([]core.SimpleUser, error)
	GetUserById(ctx context.Context, id int) (*core.SimpleUser, error)
}

type business struct {
	transcriptRepo TranscriptRepository
}

func NewBusiness(transcriptRepo TranscriptRepository) *business {
	return &business{
		transcriptRepo: transcriptRepo,
	}
}

func (b *business) CreateTranscript(ctx context.Context, content string) (int, error) {
	data := &entity.TranscriptDataCreation{
		Content: content,
	}

	if err := b.transcriptRepo.AddNewTranscript(ctx, data); err != nil {
		return 0, err
	}

	return data.Id, nil
}
