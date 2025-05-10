package ws

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type NotificationClient struct {
	conn *websocket.Conn
	send chan []byte
}

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
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
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
