package mysql

import (
	"context"

	"thinkflow-service/services/note/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) UnarchiveNote(ctx context.Context, id int) error {
	if err := repo.db.Table(entity.Note{}.TableName()).
		Where("id = ?", id).
		Update("archived", false).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
