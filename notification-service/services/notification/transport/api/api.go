package api

import (
	"context"

	"thinkflow-service/services/notification/entity"

	"github.com/VanThen60hz/service-context/core"
)

type Business interface {
	GetUnreadCount(ctx context.Context) (int64, error)
	ListNotifications(ctx context.Context, filter *entity.Filter, paging *core.Paging) ([]entity.Notification, error)
	MarkNotificationAsRead(ctx context.Context, id string) error
	MarkAllNotificationsAsRead(ctx context.Context) error
	DeleteNotification(ctx context.Context, id int) error
	CreateNotification(ctx context.Context, data *entity.NotificationCreation) error
}

type api struct {
	business Business
}

func NewAPI(business Business) *api {
	return &api{business: business}
}
