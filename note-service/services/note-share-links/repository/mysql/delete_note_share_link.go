package mysql

import (
	"context"

	"thinkflow-service/services/note-share-links/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) DeleteNoteShareLink(ctx context.Context, id int64) error {
	if err := repo.db.Where("id = ?", id).Delete(&entity.NoteShareLink{}).Error; err != nil {
		return errors.WithStack(entity.ErrCannotDeleteShareLink)
	}
	return nil
}
