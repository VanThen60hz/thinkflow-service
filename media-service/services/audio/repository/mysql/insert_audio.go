package mysql

import (
	"context"

	"thinkflow-service/services/audio/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) AddNewAudio(ctx context.Context, data *entity.AudioDataCreation) error {
	if err := repo.db.Create(data).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
