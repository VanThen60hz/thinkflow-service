package business

import (
	"context"

	"thinkflow-service/services/summary/entity"

	"github.com/VanThen60hz/service-context/core"
)

type SummaryRepository interface {
	AddNewSummary(ctx context.Context, data *entity.SummaryDataCreation) error
	UpdateSummary(ctx context.Context, id int, data *entity.SummaryDataUpdate) error
	DeleteSummary(ctx context.Context, id int) error
	GetSummaryById(ctx context.Context, id int) (*entity.Summary, error)
}

type UserRepository interface {
	GetUsersByIds(ctx context.Context, ids []int) ([]core.SimpleUser, error)
	GetUserById(ctx context.Context, id int) (*core.SimpleUser, error)
}

type business struct {
	summaryRepo SummaryRepository
}

func NewBusiness(summaryRepo SummaryRepository) *business {
	return &business{
		summaryRepo: summaryRepo,
	}
}

func (b *business) CreateSummary(ctx context.Context, summaryText string) (int, error) {
	data := &entity.SummaryDataCreation{
		SummaryText: summaryText,
	}

	if err := b.summaryRepo.AddNewSummary(ctx, data); err != nil {
		return 0, err
	}

	return data.Id, nil
}
