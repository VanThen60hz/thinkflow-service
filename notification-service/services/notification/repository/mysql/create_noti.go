package mysql

import (
	"context"

	"thinkflow-service/services/notification/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) CreateNotification(ctx context.Context, data *entity.NotificationCreation) error {
	if err := repo.db.Table(data.TableName()).Create(data).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
