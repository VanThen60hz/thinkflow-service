package mysql

import (
	"context"

	"thinkflow-service/services/collaboration/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) HasReadPermission(ctx context.Context, noteId int, userId int) (bool, error) {
	var count int64

	db := repo.db.Table(entity.Collaboration{}.TableName())

	if err := db.
		Where("note_id = ? AND user_id = ? AND permission IN (?, ?)",
			noteId, userId, entity.PermissionRead, entity.PermissionWrite).
		Count(&count).Error; err != nil {
		return false, errors.WithStack(err)
	}

	return count > 0, nil
}
