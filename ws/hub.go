package ws

import (
	"log"
)

type Hub struct {
	Clients map[*Client]bool

	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan Response
}

func NewHub() *Hub {

	return &Hub{
		Clients:    make(map[*Client]bool),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan Response),
	}
}

func (h *Hub) Run() {

	for {

		select {

		case client := <-h.Register:

			h.Clients[client] = true

			log.Printf(
				"Client Connected (%d)",
				len(h.Clients),
			)

		case client := <-h.Unregister:

			if _, ok := h.Clients[client]; ok {

				delete(h.Clients, client)

				log.Printf(
					"Client Disconnected (%d)",
					len(h.Clients),
				)
			}

		case message := <-h.Broadcast:

			for client := range h.Clients {

				if err := client.Send(message); err != nil {

					client.Conn.Close()

					delete(h.Clients, client)
				}
			}
		}
	}
}
