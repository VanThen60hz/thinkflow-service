package mysql

import (
	"context"

	"thinkflow-service/services/summary/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) UpdateSummary(ctx context.Context, id int, data *entity.SummaryDataUpdate) error {
	if err := repo.db.Table(data.TableName()).Where("id = ?", id).Updates(data).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
