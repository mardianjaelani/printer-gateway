package ws

import (
	"print-gateway/services"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
	Hub  *Hub

	PrintService *services.PrintService
}
