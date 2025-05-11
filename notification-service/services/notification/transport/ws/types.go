package ws

import "github.com/gorilla/websocket"

type NotificationClient struct {
	conn   *websocket.Conn
	send   chan []byte
	userID string
}
