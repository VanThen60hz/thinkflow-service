package api

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // You may want to add proper origin checking
	},
}

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

var hub = &NotificationHub{
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

func (hdl *api) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading to websocket: %v", err)
		return
	}

	client := &NotificationClient{
		conn: conn,
		send: make(chan []byte, 256),
	}

	hub.register <- client

	// Start goroutine to read messages from client
	go func() {
		defer func() {
			hub.unregister <- client
			client.conn.Close()
		}()

		for {
			_, _, err := client.conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Printf("Error reading message: %v", err)
				}
				break
			}
		}
	}()

	// Start goroutine to write messages to client
	go func() {
		defer func() {
			client.conn.Close()
		}()

		for {
			select {
			case message, ok := <-client.send:
				if !ok {
					client.conn.WriteMessage(websocket.CloseMessage, []byte{})
					return
				}

				w, err := client.conn.NextWriter(websocket.TextMessage)
				if err != nil {
					return
				}
				w.Write(message)

				if err := w.Close(); err != nil {
					return
				}
			}
		}
	}()
}
