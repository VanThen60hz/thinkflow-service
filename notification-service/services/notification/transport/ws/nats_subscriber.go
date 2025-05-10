package ws

import (
	"context"
	"log"

	"github.com/VanThen60hz/service-context/component/natsc"
	"github.com/VanThen60hz/service-context/component/pubsub"
)

func StartNatsSubscriber(ctx context.Context, natsClient natsc.Nats) {
	// Subscribe to notifications channel
	events, close := natsClient.Subscribe(ctx, pubsub.Channel("notifications"), "")
	defer close()

	// Handle incoming events
	for event := range events {
		log.Printf("Received notification from NATS: %+v", event)
		// Broadcast to all WebSocket clients
		Hub.BroadcastNotification(event.Data)
	}
}
