package business

import (
	"context"
	"encoding/json"

	"thinkflow-service/services/notification/entity"

	"github.com/VanThen60hz/service-context/component/pubsub"
	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) CreateNotification(ctx context.Context, data *entity.NotificationCreation) (*entity.Notification, error) {
	notification, err := biz.notiRepo.CreateNotification(ctx, data)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotCreateNotification.Error()).
			WithDebug(err.Error())
	}

	if notification.NotiSenderID > 0 {
		notification.Sender, err = biz.userRepo.GetUserById(ctx, int(notification.NotiSenderID))
		if err != nil {
			return nil, core.ErrInternalServerError.
				WithError(entity.ErrCannotGetUser.Error()).
				WithDebug(err.Error())
		}
	}

	if notification.NotiReceivedID > 0 {
		notification.Receiver, err = biz.userRepo.GetUserById(ctx, int(notification.NotiReceivedID))
		if err != nil {
			return nil, core.ErrInternalServerError.
				WithError(entity.ErrCannotGetUser.Error()).
				WithDebug(err.Error())
		}
	}

	notification.Mask()

	// Create a new map for event data with only necessary fields
	eventData := map[string]interface{}{
		"id":           notification.FakeId.String(),
		"noti_type":    notification.NotiType,
		"noti_content": notification.NotiContent,
		"noti_options": notification.NotiOptions.String,
		"is_read":      notification.IsRead,
		"created_at":   notification.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		"updated_at":   notification.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		"sender":       notification.Sender,
		"receiver":     notification.Receiver,
	}

	event := &pubsub.Event{
		Title:   data.NotiType,
		Channel: pubsub.Channel("notifications"),
		Data:    eventData,
	}

	eventDataJson, _ := json.Marshal(event.Data)
	biz.logger.Infoln("Publishing notification to NATS:", string(eventDataJson))

	if err := biz.natsClient.PublishAsync(ctx, pubsub.Channel("notifications"), event); err != nil {
		biz.logger.Errorln("Failed to publish notification to NATS:", err)
		return nil, core.ErrInternalServerError.
			WithError("Failed to publish notification").
			WithDebug(err.Error())
	}

	biz.logger.Infoln("Successfully published notification to NATS channel:", event.Channel)
	return notification, nil
}
