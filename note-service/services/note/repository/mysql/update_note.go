package mysql

import (
	"context"

	"thinkflow-service/services/note/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) UpdateNote(ctx context.Context, id int, data *entity.NoteDataUpdate) error {
	if err := repo.db.Table(data.TableName()).Where("id = ?", id).Updates(data).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
