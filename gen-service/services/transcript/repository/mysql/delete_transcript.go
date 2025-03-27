package mysql

import (
	"context"

	"thinkflow-service/services/transcript/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) DeleteTranscript(ctx context.Context, id int) error {
	if err := repo.db.Table(entity.Transcript{}.TableName()).
		Where("id = ?", id).
		Delete(nil).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
