package ws

import (
	"log"
	"net/http"

	"thinkflow-service/common"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		allowedOrigins := []string{
			"http://localhost:3001",
			"http://localhost:3002",
			"http://42.113.255.139:5500",
			"http://127.0.0.1:5500",
			"http://118.70.192.62:3001",
			"http://118.70.192.62:3002",
			"https://thinkflow-web.vercel.app",
		}

		for _, allowedOrigin := range allowedOrigins {
			if origin == allowedOrigin {
				return true
			}
		}
		log.Printf("WebSocket connection rejected from origin: %s", origin)
		return false
	},
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received WebSocket request from: %s", r.RemoteAddr)
	log.Printf("Request headers: %v", r.Header)

	// Get requester from context
	requesterVal := r.Context().Value(common.RequesterKey)
	if requesterVal == nil {
		log.Printf("No requester found in context")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	requester, ok := requesterVal.(core.Requester)
	if !ok {
		log.Printf("Invalid requester type in context")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	log.Printf("WebSocket connection request from user: %s", requester.GetSubject())

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading to websocket: %v", err)
		return
	}

	log.Printf("WebSocket connection established for user: %s", requester.GetSubject())

	client := &NotificationClient{
		conn:   conn,
		send:   make(chan []byte, 256),
		userID: requester.GetSubject(),
	}

	Hub.register <- client

	// Start goroutine to read messages from client
	go func() {
		defer func() {
			Hub.unregister <- client
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
