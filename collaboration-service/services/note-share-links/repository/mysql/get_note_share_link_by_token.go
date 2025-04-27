package mysql

import (
	"context"

	"thinkflow-service/services/note-share-links/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) GetNoteShareLinkByToken(ctx context.Context, token string) (*entity.NoteShareLink, error) {
	var link entity.NoteShareLink
	if err := repo.db.Where("token = ?", token).First(&link).Error; err != nil {
		return nil, errors.WithStack(entity.ErrNoteShareLinkNotFound)
	}
	return &link, nil
}
