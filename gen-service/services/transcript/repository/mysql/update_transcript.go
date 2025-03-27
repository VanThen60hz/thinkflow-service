package mysql

import (
	"context"

	"thinkflow-service/services/transcript/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) UpdateTranscript(ctx context.Context, id int, data *entity.TranscriptDataUpdate) error {
	if err := repo.db.Table(data.TableName()).Where("id = ?", id).Updates(data).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
