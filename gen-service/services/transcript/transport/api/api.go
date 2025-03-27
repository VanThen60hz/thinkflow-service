package api

import (
	"context"

	"thinkflow-service/services/transcript/entity"

	sctx "github.com/VanThen60hz/service-context"
)

type ServiceContext interface {
	sctx.ServiceContext
	Business
}

type Business interface {
	CreateNewTranscript(ctx context.Context, data *entity.TranscriptDataCreation) error
	GetTranscriptById(ctx context.Context, id int) (*entity.Transcript, error)
	UpdateTranscript(ctx context.Context, id int, data *entity.TranscriptDataUpdate) error
	DeleteTranscript(ctx context.Context, id int) error
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
