package business

import (
	"context"

	"thinkflow-service/services/mindmap/entity"

	"gorm.io/datatypes"

	"github.com/VanThen60hz/service-context/core"
)

type MindmapRepository interface {
	AddNewMindmap(ctx context.Context, data *entity.MindmapDataCreation) error
	UpdateMindmap(ctx context.Context, id int, data *entity.MindmapDataUpdate) error
	DeleteMindmap(ctx context.Context, id int) error
	GetMindmapById(ctx context.Context, id int) (*entity.Mindmap, error)
}

type UserRepository interface {
	GetUsersByIds(ctx context.Context, ids []int) ([]core.SimpleUser, error)
	GetUserById(ctx context.Context, id int) (*core.SimpleUser, error)
}

type business struct {
	mindmapRepo MindmapRepository
}

func NewBusiness(mindmapRepo MindmapRepository) *business {
	return &business{
		mindmapRepo: mindmapRepo,
	}
}

func (b *business) CreateMindmap(ctx context.Context, mindmapData string) (int, error) {
	data := &entity.MindmapDataCreation{
		MindmapData: datatypes.JSON(mindmapData),
	}

	if err := b.mindmapRepo.AddNewMindmap(ctx, data); err != nil {
		return 0, err
	}

	return data.Id, nil
}
