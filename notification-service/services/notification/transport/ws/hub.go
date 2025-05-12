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
			log.Printf("Client registered: %s, total clients: %d", client.userID, len(h.clients))
		case client := <-h.unregister:
			h.mutex.Lock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
				log.Printf("Client unregistered: %s, remaining clients: %d", client.userID, len(h.clients))
			}
			h.mutex.Unlock()
		case message := <-h.broadcast:
			h.mutex.Lock()
			// Parse the notification to get receiver ID
			var notification map[string]interface{}
			if err := json.Unmarshal(message, &notification); err != nil {
				log.Printf("Error unmarshaling notification in hub: %v", err)
				log.Printf("Raw message: %s", string(message))
				h.mutex.Unlock()
				continue
			}

			// Get receiver ID from notification
			receiver, ok := notification["receiver"].(map[string]interface{})
			if !ok {
				log.Printf("Invalid receiver data in notification: %v", notification)
				h.mutex.Unlock()
				continue
			}

			receiverID, ok := receiver["id"].(string)
			if !ok {
				log.Printf("Invalid receiver ID in notification: %v", receiver)
				h.mutex.Unlock()
				continue
			}

			log.Printf("Broadcasting notification to user: %s", receiverID)
			clientsFound := false

			// Send notification only to the intended receiver
			for client := range h.clients {
				if client.userID == receiverID {
					clientsFound = true
					select {
					case client.send <- message:
						log.Printf("Notification sent to client: %s", client.userID)
					default:
						log.Printf("Failed to send to client (buffer full): %s", client.userID)
						close(client.send)
						delete(h.clients, client)
					}
				}
			}

			if !clientsFound {
				log.Printf("No connected clients found for user: %s", receiverID)
			}

			h.mutex.Unlock()
		}
	}
}

func (h *NotificationHub) BroadcastNotification(notification interface{}) {
	var data []byte
	var err error

	// Handle different types of input
	switch n := notification.(type) {
	case []byte:
		data = n
	case string:
		data = []byte(n)
	default:
		data, err = json.Marshal(n)
		if err != nil {
			log.Printf("Error marshaling notification: %v", err)
			return
		}
	}

	// Verify the data is valid JSON
	var testMap map[string]interface{}
	if err := json.Unmarshal(data, &testMap); err != nil {
		log.Printf("Invalid JSON in notification: %v", err)
		return
	}

	log.Printf("Broadcasting notification: %s", string(data))
	h.broadcast <- data
}
