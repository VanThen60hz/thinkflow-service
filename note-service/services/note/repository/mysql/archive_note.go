package mysql

import (
	"context"

	"thinkflow-service/services/note/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) ArchiveNote(ctx context.Context, id int) error {
	if err := repo.db.Table(entity.Note{}.TableName()).
		Where("id = ?", id).
		Update("archived", true).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
