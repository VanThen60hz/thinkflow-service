package mysql

import (
	"context"

	"thinkflow-service/services/notification/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) GetUnreadCount(ctx context.Context, requesterId int64) (int64, error) {
	var count int64
	db := repo.db.Table(entity.Notification{}.TableName())

	if err := db.Where("noti_received_id = ? AND is_read = ?", requesterId, false).
		Count(&count).Error; err != nil {
		return 0, errors.WithStack(err)
	}

	return count, nil
}
