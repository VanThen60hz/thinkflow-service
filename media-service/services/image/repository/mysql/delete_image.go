package mysql

import (
	"context"

	"thinkflow-service/services/image/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) DeleteImage(ctx context.Context, id int) error {
	if err := repo.db.Table(entity.Image{}.TableName()).
		Where("id = ?", id).
		Delete(nil).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
