package mysql

import (
	"context"

	"thinkflow-service/services/note/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) UpdateNote(ctx context.Context, id int, data *entity.NoteDataUpdate) error {
	var note entity.Note
	if err := repo.db.Table(entity.Note{}.TableName()).
		Where("id = ?", id).
		First(&note).Error; err != nil {
		return errors.WithStack(err)
	}

	if note.Archived {
		return entity.ErrCannotUpdateArchivedNote
	}

	if err := repo.db.Table(entity.Note{}.TableName()).
		Where("id = ?", id).
		Updates(data).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
