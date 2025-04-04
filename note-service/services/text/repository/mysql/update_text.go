package mysql

import (
	"context"

	"thinkflow-service/services/text/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) UpdateText(ctx context.Context, id int, data *entity.TextDataUpdate) error {
	if err := repo.db.Where("id = ?", id).Updates(data).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
