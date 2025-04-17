package business

import (
	"context"

	"thinkflow-service/services/image/entity"

	"github.com/VanThen60hz/service-context/component/s3c"
	"github.com/VanThen60hz/service-context/core"
)

type ImageRepository interface {
	AddNewImage(ctx context.Context, data *entity.ImageDataCreation) error
	UpdateImage(ctx context.Context, id int, data *entity.ImageDataUpdate) error
	DeleteImage(ctx context.Context, id int) error
	GetImageById(ctx context.Context, id int) (*entity.Image, error)
	ListImages(ctx context.Context, filter *entity.Filter, paging *core.Paging) ([]entity.Image, error)
}

type business struct {
	imageRepo ImageRepository
	s3Client  *s3c.S3Component
}

func NewBusiness(imageRepo ImageRepository, s3Client *s3c.S3Component) *business {
	return &business{
		imageRepo: imageRepo,
		s3Client:  s3Client,
	}
}
