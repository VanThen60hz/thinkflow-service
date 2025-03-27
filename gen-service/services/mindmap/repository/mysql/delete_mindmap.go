package mysql

import (
	"context"

	"thinkflow-service/services/mindmap/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) DeleteMindmap(ctx context.Context, id int) error {
	if err := repo.db.Table(entity.Mindmap{}.TableName()).
		Where("id = ?", id).
		Delete(nil).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
