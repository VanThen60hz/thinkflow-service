package mysql

import (
	"context"

	"thinkflow-service/services/text/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) GetTextByNoteId(ctx context.Context, noteId int) (*entity.Text, error) {
	var text entity.Text

	db := repo.db.Table(entity.Text{}.TableName())

	if err := db.Where("note_id = ?", noteId).First(&text).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	text.Mask()
	return &text, nil
}
