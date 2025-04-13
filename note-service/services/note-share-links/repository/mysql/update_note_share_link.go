package mysql

import (
	"context"

	"thinkflow-service/services/note-share-links/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) UpdateNoteShareLink(ctx context.Context, id int64, data *entity.NoteShareLinkUpdate) error {
	if err := repo.db.Model(&entity.NoteShareLink{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return errors.WithStack(entity.ErrCannotUpdateShareLink)
	}
	return nil
}
