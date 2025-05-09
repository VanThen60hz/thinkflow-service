package business

import (
	"context"

	"thinkflow-service/services/notification/entity"

	"github.com/VanThen60hz/service-context/core"
)

type NotificationRepository interface {
	GetUnreadCount(ctx context.Context, requesterId int64) (int64, error)
	CreateNotification(ctx context.Context, data *entity.NotificationCreation) error
	GetNotificationById(ctx context.Context, id int) (*entity.Notification, error)
	ListNotifications(ctx context.Context, filter *entity.Filter, paging *core.Paging) ([]entity.Notification, error)
	MarkNotificationAsRead(ctx context.Context, id string) error
	MarkAllNotificationsAsRead(ctx context.Context, requesterId int64) error
	DeleteNotification(ctx context.Context, id int) error
}

type AuthRepository interface {
	RegisterWithUserId(ctx context.Context, userId int32, email, password string) error
}

type Business interface{}

type business struct {
	notiRepo NotificationRepository
	authRepo AuthRepository
}

func NewBusiness(notiRepo NotificationRepository, authRepo AuthRepository) *business {
	return &business{
		notiRepo: notiRepo,
		authRepo: authRepo,
	}
}
