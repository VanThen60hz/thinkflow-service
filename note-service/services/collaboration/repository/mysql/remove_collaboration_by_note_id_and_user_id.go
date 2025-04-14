package mysql

import (
	"context"

	"thinkflow-service/services/collaboration/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) RemoveCollaborationByNoteIdAndUserId(ctx context.Context, noteId int, userId int) error {
	db := repo.db.Table(entity.Collaboration{}.TableName()).
		Where("note_id = ? AND user_id = ?", noteId, userId)

	if err := db.Delete(nil).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
