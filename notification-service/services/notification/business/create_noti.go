package business

import (
	"context"
	"encoding/json"

	"thinkflow-service/services/notification/entity"
	"thinkflow-service/services/notification/transport/ws"

	"github.com/VanThen60hz/service-context/component/pubsub"
	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) CreateNotification(ctx context.Context, data *entity.NotificationCreation) error {
	err := biz.notiRepo.CreateNotification(ctx, data)
	if err != nil {
		return core.ErrInternalServerError.
			WithError(entity.ErrCannotCreateNotification.Error()).
			WithDebug(err.Error())
	}

	event := &pubsub.Event{
		Title:   data.NotiType,
		Channel: pubsub.Channel("notifications"),
		Data: map[string]interface{}{
			"id":               data.FakeId.String(),
			"noti_type":        data.NotiType,
			"noti_sender_id":   data.NotiSenderID,
			"noti_received_id": data.NotiReceivedID,
			"noti_content":     data.NotiContent,
			"noti_options":     data.NotiOptions.String,
			"is_read":          data.IsRead,
			"created_at":       data.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			"updated_at":       data.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		},
	}

	// Log the event data before publishing
	eventData, _ := json.Marshal(event.Data)
	biz.logger.Infoln("Publishing notification to NATS:", string(eventData))

	// Use PublishAsync to avoid blocking
	if err := biz.natsClient.PublishAsync(ctx, pubsub.Channel("notifications"), event); err != nil {
		biz.logger.Errorln("Failed to publish notification to NATS:", err)
		return core.ErrInternalServerError.
			WithError("Failed to publish notification").
			WithDebug(err.Error())
	}

	ws.Hub.BroadcastNotification(event.Data)

	biz.logger.Infoln("Successfully published notification to NATS channel:", event.Channel)
	return nil
}
