package api

import (
	"context"
	"mime/multipart"

	"thinkflow-service/services/audio/entity"

	sctx "github.com/VanThen60hz/service-context"
	"github.com/VanThen60hz/service-context/core"
)

type ServiceContext interface {
	sctx.ServiceContext
	Business
}

type Business interface {
	UploadAudio(ctx context.Context, tempFile string, file *multipart.FileHeader, noteID int64) (*entity.AudioDataCreation, error)
	GetAudioById(ctx context.Context, id int) (*entity.Audio, error)
	GetAudiosByNoteId(ctx context.Context, noteId int) ([]entity.Audio, error)
	ListAudios(ctx context.Context, filter *entity.Filter, paging *core.Paging) ([]entity.Audio, error)
	UpdateAudio(ctx context.Context, id int, data *entity.AudioDataUpdate) error
	DeleteAudio(ctx context.Context, id int) error
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
