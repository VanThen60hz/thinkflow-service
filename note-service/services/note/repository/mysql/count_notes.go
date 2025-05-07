package mysql

import (
	"context"

	"thinkflow-service/services/note/entity"
)

func (repo *mysqlRepo) CountNotes(ctx context.Context) (int64, error) {
	var count int64
	if err := repo.db.Model(&entity.Note{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
