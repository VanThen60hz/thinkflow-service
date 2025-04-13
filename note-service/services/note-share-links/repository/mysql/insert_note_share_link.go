package mysql

import (
	"context"

	"thinkflow-service/services/note-share-links/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) AddNewNoteShareLink(ctx context.Context, data *entity.NoteShareLinkCreation) error {
	if err := repo.db.Create(data).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
