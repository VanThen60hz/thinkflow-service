package api

import (
	"context"

	"thinkflow-service/services/text/entity"

	sctx "github.com/VanThen60hz/service-context"
)

type ServiceContext interface {
	sctx.ServiceContext
	Business
}

type Business interface {
	CreateNewText(ctx context.Context, data *entity.TextDataCreation) error
	GetTextById(ctx context.Context, id int) (*entity.Text, error)
	GetTextByNoteId(ctx context.Context, noteId int) (*entity.Text, error)
	UpdateText(ctx context.Context, id int, data *entity.TextDataUpdate) error
	DeleteText(ctx context.Context, id int) error
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
