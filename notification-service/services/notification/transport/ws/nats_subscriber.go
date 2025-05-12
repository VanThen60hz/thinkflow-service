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
		
		// Check the type of event.Data and handle accordingly
		switch data := event.Data.(type) {
		case []byte:
			if err := json.Unmarshal(data, &notification); err != nil {
				log.Printf("Error unmarshaling notification from []byte: %v", err)
				continue
			}
		case map[string]interface{}:
			notification = data
		default:
			log.Printf("Unexpected data type: %T", event.Data)
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
		// Convert to JSON if it's not already in []byte format
		var notificationData []byte
		if byteData, ok := event.Data.([]byte); ok {
			notificationData = byteData
		} else {
			var err error
			notificationData, err = json.Marshal(notification)
			if err != nil {
				log.Printf("Error marshaling notification: %v", err)
				continue
			}
		}
		Hub.BroadcastNotification(notificationData)

		// Send to FCM
		title, _ := notification["title"].(string)
		body, _ := notification["body"].(string)
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
