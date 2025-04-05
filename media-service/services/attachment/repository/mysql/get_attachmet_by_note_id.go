package mysql

import (
	"context"

	"thinkflow-service/services/attachment/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) GetAttachmentsByNoteId(ctx context.Context, noteId int) (*entity.Attachment, error) {
	var text entity.Attachment

	db := repo.db.Table(entity.Attachment{}.TableName())

	if err := db.Where("note_id = ?", noteId).First(&text).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	text.Mask()
	return &text, nil
}
