package mysql

import (
	"context"

	"thinkflow-service/services/note/entity"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (repo *mysqlRepo) DeleteUserNotes(ctx context.Context, userId int) (int, error) {
	var result *gorm.DB
	var deletedCount int64

	result = repo.db.
		Table(entity.Note{}.TableName()).
		Where("user_id = ?", userId).
		Delete(&entity.Note{})

	if result.Error != nil {
		return 0, errors.WithStack(result.Error)
	}

	deletedCount = result.RowsAffected

	return int(deletedCount), nil
}
