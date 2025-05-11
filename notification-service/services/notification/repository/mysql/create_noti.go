package mysql

import (
	"context"

	"thinkflow-service/services/notification/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) CreateNotification(ctx context.Context, data *entity.NotificationCreation) (*entity.Notification, error) {
	if err := repo.db.Table(data.TableName()).Create(data).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	notification := &entity.Notification{
		SQLModel:       data.SQLModel,
		NotiType:       data.NotiType,
		NotiSenderID:   data.NotiSenderID,
		NotiReceivedID: data.NotiReceivedID,
		NotiContent:    data.NotiContent,
		NotiOptions:    data.NotiOptions,
		IsRead:         data.IsRead,
	}

	return notification, nil
}
