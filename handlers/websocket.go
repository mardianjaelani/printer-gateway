package handlers

import (
	"log"
	"net/http"

	"print-gateway/ws"

	"github.com/gin-gonic/gin"
	gws "github.com/gorilla/websocket"
)

var upgrader = gws.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	CheckOrigin: func(r *http.Request) bool {
		// Nanti bisa divalidasi berdasarkan config.json
		return true
	},
}

func HandleWS(hub *ws.Hub) gin.HandlerFunc {

	return func(c *gin.Context) {

		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println("WebSocket Upgrade Error:", err)
			return
		}

		client := &ws.Client{
			Conn: conn,
			Hub:  hub,
		}

		hub.Register <- client

		log.Printf("Client Connected : %s", conn.RemoteAddr())

		go client.ReadPump()
	}
}
