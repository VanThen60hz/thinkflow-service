package mysql

import (
	"context"
	"thinkflow-service/services/media/entity"

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

func (repo *mysqlRepo) DeleteAudio(ctx context.Context, id int) error {
	if err := repo.db.Table(entity.Audio{}.TableName()).
		Where("id = ?", id).
		Delete(nil).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
