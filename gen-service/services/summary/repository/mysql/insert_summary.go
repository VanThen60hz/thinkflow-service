package mysql

import (
	"context"

	"thinkflow-service/services/summary/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) AddNewSummary(ctx context.Context, data *entity.SummaryDataCreation) error {
	if err := repo.db.Create(data).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
