package api

import (
	"context"

	"thinkflow-service/services/mindmap/entity"

	sctx "github.com/VanThen60hz/service-context"
)

type ServiceContext interface {
	sctx.ServiceContext
	Business
}

type Business interface {
	CreateNewMindmap(ctx context.Context, data *entity.MindmapDataCreation) error
	GetMindmapById(ctx context.Context, id int) (*entity.Mindmap, error)
	UpdateMindmap(ctx context.Context, id int, data *entity.MindmapDataUpdate) error
	DeleteMindmap(ctx context.Context, id int) error
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
