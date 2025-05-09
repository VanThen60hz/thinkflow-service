package mysql

import (
	"context"

	"thinkflow-service/services/notification/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) MarkAllNotificationsAsRead(ctx context.Context, requesterId int64) error {
	db := repo.db.Table(entity.Notification{}.TableName())

	if err := db.Where("noti_received_id = ?", requesterId).
		Update("is_read", true).Error; err != nil {
			return errors.WithStack(err)
		}
	return nil
}