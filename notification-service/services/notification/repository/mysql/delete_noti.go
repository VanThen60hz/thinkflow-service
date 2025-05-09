package mysql

import (
	"context"

	"thinkflow-service/services/notification/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) DeleteNotification(ctx context.Context, id int) error {
	db := repo.db.Table(entity.Notification{}.TableName())

	if err := db.Where("id = ?", id).Delete(&entity.Notification{}).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
