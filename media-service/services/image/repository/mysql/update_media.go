package mysql

import (
	"context"

	"thinkflow-service/services/image/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) UpdateImage(ctx context.Context, id int, data *entity.ImageDataUpdate) error {
	if err := repo.db.Where("id = ?", id).Updates(data).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
