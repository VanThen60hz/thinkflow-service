package mysql

import (
	"context"

	"thinkflow-service/services/image/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) AddNewImage(ctx context.Context, data *entity.ImageDataCreation) error {
	if err := repo.db.Create(data).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
