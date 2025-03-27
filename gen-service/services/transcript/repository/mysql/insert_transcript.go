package mysql

import (
	"context"

	"thinkflow-service/services/transcript/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) AddNewTranscript(ctx context.Context, data *entity.TranscriptDataCreation) error {
	if err := repo.db.Create(data).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
