package business

import (
	"context"
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

type UserRepository interface {
	GetUsersByIds(ctx context.Context, ids []int) ([]core.SimpleUser, error)
	GetUserById(ctx context.Context, id int) (*core.SimpleUser, error)
}

type business struct {
	mediaRepo MediaRepository
	userRepo  UserRepository
}

func NewBusiness(mediaRepo MediaRepository, userRepo UserRepository) *business {
	return &business{
		mediaRepo: mediaRepo,
		userRepo:  userRepo,
	}
}
