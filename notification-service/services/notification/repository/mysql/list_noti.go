package mysql

import (
	"context"

	"thinkflow-service/services/notification/entity"

	"github.com/VanThen60hz/service-context/core"
	"github.com/pkg/errors"
)

func (repo *mysqlRepo) ListNotifications(ctx context.Context, filter *entity.Filter, paging *core.Paging) ([]entity.Notification, error) {
	var notis []entity.Notification
	db := repo.db.Table(entity.Notification{}.TableName())

	if notiReceivedID := filter.NotiReceivedID; notiReceivedID != nil {
		uid, err := core.FromBase58(*notiReceivedID)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		db = db.Where("noti_received_id = ?", uid.GetLocalID())
	}

	if filter.NotiType != nil && *filter.NotiType != "" {
		db = db.Where("noti_type = ?", *filter.NotiType)
	}

	if err := db.Select("id").Count(&paging.Total).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	if v := paging.FakeCursor; v != "" {
		uid, err := core.FromBase58(v)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		db = db.Where("id < ?", uid.GetLocalID())
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.
		Select("id, noti_type, noti_sender_id, noti_received_id, noti_content, noti_options, is_read, created_at, updated_at").
		Order("id desc").
		Limit(paging.Limit).
		Find(&notis).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	if len(notis) > 0 {
		notis[len(notis)-1].Mask()
		paging.NextCursor = notis[len(notis)-1].FakeId.String()
	}

	return notis, nil
}
