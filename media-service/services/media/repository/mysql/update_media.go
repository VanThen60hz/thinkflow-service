package mysql

import (
	"context"
	"thinkflow-service/services/media/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) UpdateImage(ctx context.Context, id int, data *entity.ImageDataUpdate) error {
	if err := repo.db.Where("id = ?", id).Updates(data).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (repo *mysqlRepo) UpdateAudio(ctx context.Context, id int, data *entity.AudioDataUpdate) error {
	if err := repo.db.Where("id = ?", id).Updates(data).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
