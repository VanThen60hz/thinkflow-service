package mysql

import (
	"context"
	"time"

	"thinkflow-service/services/note/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) DeleteArchivedNotesOlderThan(ctx context.Context, days int) error {
	cutoffDate := time.Now().AddDate(0, 0, -days)

	if err := repo.db.Table(entity.Note{}.TableName()).
		Where("archived = ? AND updated_at < ?", true, cutoffDate).
		Delete(&entity.Note{}).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
