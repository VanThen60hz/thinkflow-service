package mysql

import (
	"context"

	"thinkflow-service/services/collaboration/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) UpdateCollaboration(ctx context.Context, id int, data *entity.Collaboration) error {
	if err := repo.db.Where("id = ?", id).Updates(data).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
