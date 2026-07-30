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

		log.Println("Incoming WS Request")

		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println("Upgrade Error:", err)
			return
		}

		log.Println("Upgrade Success")

		client := &ws.Client{
			Conn: conn,
			Hub:  hub,
		}

		log.Println("Register Client")

		hub.Register <- client

		log.Println("Start ReadPump")

		go client.ReadPump()
	}
}
