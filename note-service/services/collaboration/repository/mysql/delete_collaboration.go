package mysql

import (
	"context"
	"thinkflow-service/services/collaboration/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) DeleteCollaboration(ctx context.Context, id int) error {
	db := repo.db.Table(entity.Collaboration{}.TableName()).Where("id = ?", id)

	if err := db.Delete(nil).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}