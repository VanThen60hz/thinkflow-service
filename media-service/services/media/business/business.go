package business

import (
	"context"
	"fmt"

	"thinkflow-service/services/media/entity"

	"github.com/VanThen60hz/service-context/core"
)

type MediaRepository interface {
	AddNewImage(ctx context.Context, data *entity.ImageDataCreation) error
	AddNewAudio(ctx context.Context, data *entity.AudioDataCreation) error
	UpdateImage(ctx context.Context, id int, data *entity.ImageDataUpdate) error
	UpdateAudio(ctx context.Context, id int, data *entity.AudioDataUpdate) error
	DeleteImage(ctx context.Context, id int) error
	DeleteAudio(ctx context.Context, id int) error
	GetImageById(ctx context.Context, id int) (*entity.Image, error)
	GetAudioById(ctx context.Context, id int) (*entity.Audio, error)
	ListImages(ctx context.Context, filter *entity.Filter, paging *core.Paging) ([]entity.Image, error)
	ListAudios(ctx context.Context, filter *entity.Filter, paging *core.Paging) ([]entity.Audio, error)
}

type business struct {
	mediaRepo MediaRepository
}

func NewBusiness(mediaRepo MediaRepository) *business {
	return &business{
		mediaRepo: mediaRepo,
	}
}

func (b *business) GetImagesByIds(ctx context.Context, ids []int) ([]entity.Image, error) {
	// Implement the logic to get images by IDs
	// For example, query the database or another data source
	images := make([]entity.Image, len(ids))
	for i, id := range ids {
		images[i] = entity.Image{
			Url: fmt.Sprintf("http://example.com/image%d.jpg", id), // Example URL
		}
	}
	return images, nil
}
