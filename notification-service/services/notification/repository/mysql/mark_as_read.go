package mysql

import (
	"context"

	"thinkflow-service/services/notification/entity"

	"github.com/VanThen60hz/service-context/core"
	"github.com/pkg/errors"
)

func (repo *mysqlRepo) MarkNotificationAsRead(ctx context.Context, id string) error {
	uid, err := core.FromBase58(id)
	if err != nil {
		return errors.WithStack(err)
	}

	result := repo.db.Table(entity.Notification{}.TableName()).
		Where("id = ?", uid.GetLocalID()).
		Update("is_read", true)

	if result.Error != nil {
		return errors.WithStack(result.Error)
	}

	if result.RowsAffected == 0 {
		return errors.New("notification not found")
	}

	return nil
}
