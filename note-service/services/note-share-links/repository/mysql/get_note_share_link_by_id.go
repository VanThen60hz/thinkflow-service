package mysql

import (
	"context"

	"thinkflow-service/services/note-share-links/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) GetNoteShareLinkByID(ctx context.Context, id int64) (*entity.NoteShareLink, error) {
	var link entity.NoteShareLink
	if err := repo.db.First(&link, id).Error; err != nil {
		return nil, errors.WithStack(entity.ErrNoteShareLinkNotFound)
	}
	return &link, nil
}
