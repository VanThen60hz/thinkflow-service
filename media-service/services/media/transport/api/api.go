package api

import (
	"context"
	"thinkflow-service/services/media/entity"

	sctx "github.com/VanThen60hz/service-context"
	"github.com/VanThen60hz/service-context/core"
)

type ServiceContext interface {
	sctx.ServiceContext
	Business
}

type Business interface {
	CreateNewImage(ctx context.Context, data *entity.ImageDataCreation) error
	CreateNewAudio(ctx context.Context, data *entity.AudioDataCreation) error
	GetImageById(ctx context.Context, id int) (*entity.Image, error)
	GetAudioById(ctx context.Context, id int) (*entity.Audio, error)
	ListImages(ctx context.Context, filter *entity.Filter, paging *core.Paging) ([]entity.Image, error)
	ListAudios(ctx context.Context, filter *entity.Filter, paging *core.Paging) ([]entity.Audio, error)
	UpdateImage(ctx context.Context, id int, data *entity.ImageDataUpdate) error
	UpdateAudio(ctx context.Context, id int, data *entity.AudioDataUpdate) error
	DeleteImage(ctx context.Context, id int) error
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
