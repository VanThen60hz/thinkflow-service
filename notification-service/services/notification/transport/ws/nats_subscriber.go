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

func getTitleFromType(notiType string) string {
	switch notiType {
	case "TRANSCRIPT_GENERATED":
		return "Transcript has been generated"
	case "SUMMARY_GENERATED":
		return "Summary is ready"
	case "MINDMAP_GENERATED":
		return "Mindmap has been created"
	case "COLLAB_INVITE":
		return "You've been invited to collaborate"
	default:
		return "You have a new notification"
	}
}

func StartNatsSubscriber(ctx context.Context, natsClient natsc.Nats, fcmService *fcm.Service) {
	events, close := natsClient.Subscribe(ctx, pubsub.Channel("notifications"), "")
	defer close()

	for event := range events {
		var notification map[string]interface{}

		switch data := event.Data.(type) {
		case []byte:
			if err := json.Unmarshal(data, &notification); err != nil {
				log.Printf("Error unmarshaling notification from []byte: %v", err)
				continue
			}
		case string:
			if err := json.Unmarshal([]byte(data), &notification); err != nil {
				log.Printf("Error unmarshaling notification from string: %v", err)
				continue
			}
		case map[string]interface{}:
			notification = data
		default:
			log.Printf("Unexpected data type: %T", event.Data)
			continue
		}

		receiver, ok := notification["receiver"].(map[string]interface{})
		if !ok {
			log.Printf("Invalid receiver in notification: %v", notification)
			continue
		}

		receiverID, ok := receiver["id"].(string)
		if !ok {
			log.Printf("Invalid or missing receiver ID in notification")
			continue
		}

		var notificationData []byte
		if byteData, ok := event.Data.([]byte); ok {
			notificationData = byteData
		} else if strData, ok := event.Data.(string); ok {
			notificationData = []byte(strData)
		} else {
			var err error
			notificationData, err = json.Marshal(notification)
			if err != nil {
				log.Printf("Error marshaling notification: %v", err)
				continue
			}
		}
		Hub.BroadcastNotification(notificationData)

		// Extract fields
		title, _ := notification["title"].(string)
		body, _ := notification["body"].(string)
		notiType, _ := notification["noti_type"].(string)
		notiContent, _ := notification["noti_content"].(string)

		// Improved title fallback
		if title == "" {
			title = getTitleFromType(notiType)
		}
		if body == "" && notiContent != "" {
			body = notiContent
		}

		data := make(map[string]string)
		for k, v := range notification {
			if k != "title" && k != "body" {
				data[k] = fmt.Sprintf("%v", v)
			}
		}

		if (title != "" || body != "") && receiverID != "" {
			if err := fcmService.SendNotification(ctx, receiverID, title, body, data); err != nil {
				log.Printf("Error sending FCM notification: %v", err)
			}
		}
	}
}
