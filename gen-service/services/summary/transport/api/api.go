package api

import (
	"context"

	"thinkflow-service/services/summary/entity"

	sctx "github.com/VanThen60hz/service-context"
)

type ServiceContext interface {
	sctx.ServiceContext
	Business
}

type Business interface {
	CreateNewSummary(ctx context.Context, data *entity.SummaryDataCreation) error
	GetSummaryById(ctx context.Context, id int) (*entity.Summary, error)
	UpdateSummary(ctx context.Context, id int, data *entity.SummaryDataUpdate) error
	DeleteSummary(ctx context.Context, id int) error
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
