package api

import (
	"context"
	"mime/multipart"

	"thinkflow-service/services/image/entity"

	sctx "github.com/VanThen60hz/service-context"
	"github.com/VanThen60hz/service-context/core"
)

type ServiceContext interface {
	sctx.ServiceContext
	Business
}

type Business interface {
	UploadImage(ctx context.Context, tempFile string, file *multipart.FileHeader) (*entity.ImageDataCreation, error)
	GetImageById(ctx context.Context, id int) (*entity.Image, error)
	ListImages(ctx context.Context, filter *entity.Filter, paging *core.Paging) ([]entity.Image, error)
	UpdateImage(ctx context.Context, id int, data *entity.ImageDataUpdate) error
	DeleteImage(ctx context.Context, id int) error
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
