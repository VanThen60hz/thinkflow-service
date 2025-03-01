package mysql

import (
	"context"
	"thinkflow-service/services/media/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) AddNewImage(ctx context.Context, data *entity.ImageDataCreation) error {
	if err := repo.db.Create(data).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (repo *mysqlRepo) AddNewAudio(ctx context.Context, data *entity.AudioDataCreation) error {
	if err := repo.db.Create(data).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
