package ws

import (
	"encoding/json"
	"log"
	"sync"
)

type NotificationHub struct {
	clients    map[*NotificationClient]bool
	broadcast  chan []byte
	register   chan *NotificationClient
	unregister chan *NotificationClient
	mutex      sync.Mutex
}

var Hub = &NotificationHub{
	clients:    make(map[*NotificationClient]bool),
	broadcast:  make(chan []byte),
	register:   make(chan *NotificationClient),
	unregister: make(chan *NotificationClient),
}

func (h *NotificationHub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mutex.Lock()
			h.clients[client] = true
			h.mutex.Unlock()
		case client := <-h.unregister:
			h.mutex.Lock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
			h.mutex.Unlock()
		case message := <-h.broadcast:
			h.mutex.Lock()
			// Parse the notification to get receiver ID
			var notification map[string]interface{}
			if err := json.Unmarshal(message, &notification); err != nil {
				log.Printf("Error unmarshaling notification: %v", err)
				h.mutex.Unlock()
				continue
			}

			// Get receiver ID from notification
			receiver, ok := notification["receiver"].(map[string]interface{})
			if !ok {
				log.Printf("Invalid receiver data in notification")
				h.mutex.Unlock()
				continue
			}

			receiverID, ok := receiver["id"].(string)
			if !ok {
				log.Printf("Invalid receiver ID in notification")
				h.mutex.Unlock()
				continue
			}

			// Send notification only to the intended receiver
			for client := range h.clients {
				if client.userID == receiverID {
					select {
					case client.send <- message:
					default:
						close(client.send)
						delete(h.clients, client)
					}
				}
			}
			h.mutex.Unlock()
		}
	}
}

func (h *NotificationHub) BroadcastNotification(notification interface{}) {
	data, err := json.Marshal(notification)
	if err != nil {
		log.Printf("Error marshaling notification: %v", err)
		return
	}
	h.broadcast <- data
}
