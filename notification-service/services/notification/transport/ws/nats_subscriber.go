package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"thinkflow-service/services/notification/transport/fcm"

	"github.com/VanThen60hz/service-context/component/natsc"
	"github.com/VanThen60hz/service-context/component/pubsub"
)

// Trong nats_subscriber.go
func StartNatsSubscriber(ctx context.Context, natsClient natsc.Nats, fcmService *fcm.Service) {
	events, close := natsClient.Subscribe(ctx, pubsub.Channel("notifications"), "")
	defer close()

	for event := range events {
		// Parse notification data
		var notification map[string]interface{}
		if err := json.Unmarshal(event.Data.([]byte), &notification); err != nil {
			log.Printf("Error unmarshaling notification: %v", err)
			continue
		}

		// Get receiver ID
		receiver, ok := notification["receiver"].(map[string]interface{})
		if !ok {
			continue
		}

		receiverID, ok := receiver["id"].(string)
		if !ok {
			continue
		}

		// Send to WebSocket
		Hub.BroadcastNotification(event.Data)

		// Send to FCM
		title := notification["title"].(string)
		body := notification["body"].(string)
		data := make(map[string]string)
		for k, v := range notification {
			if k != "title" && k != "body" {
				data[k] = fmt.Sprintf("%v", v)
			}
		}

		if err := fcmService.SendNotification(ctx, receiverID, title, body, data); err != nil {
			log.Printf("Error sending FCM notification: %v", err)
		}
	}
}
