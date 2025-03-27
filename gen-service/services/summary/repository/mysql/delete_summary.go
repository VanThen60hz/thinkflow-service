package mysql

import (
	"context"

	"thinkflow-service/services/summary/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) DeleteSummary(ctx context.Context, id int) error {
	if err := repo.db.Table(entity.Summary{}.TableName()).
		Where("id = ?", id).
		Delete(nil).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
