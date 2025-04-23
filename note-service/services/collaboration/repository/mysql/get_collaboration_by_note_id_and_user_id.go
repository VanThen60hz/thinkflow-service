package mysql

import (
	"context"

	"thinkflow-service/services/collaboration/entity"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (repo *mysqlRepo) GetCollaborationByNoteIdAndUserId(ctx context.Context, noteId int, userId int) (*entity.Collaboration, error) {
	collab := &entity.Collaboration{}
	if err := repo.db.WithContext(ctx).Where("note_id = ? AND user_id = ?", noteId, userId).First(collab).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithStack(err)
		}
		return nil, errors.Wrap(err, "failed to get collaboration by note id and user id")
	}
	return collab, nil
}
