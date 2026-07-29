package ws

import (
	"encoding/json"
	"log"
)

func (c *Client) ReadPump() {

	defer func() {

		c.Hub.Unregister <- c

		c.Conn.Close()

		log.Println("Client Disconnected :", c.Conn.RemoteAddr())

	}()

	for {

		_, data, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println("Read Error :", err)
			break
		}

		var req Message

		if err := json.Unmarshal(data, &req); err != nil {
			c.Error("Format JSON tidak valid")
			continue
		}

		if req.Action == "" {
			c.Error("Action tidak boleh kosong")
			continue
		}

		switch req.Action {

		case "status":
			c.Status()

		// case "printers":
		// 	c.ListPrinters()

		case "defaultPrinter":
			c.DefaultPrinter()

		case "print":

			switch req.Type {

			case "pdf":
				c.PrintPdf(req)

			case "raw":
				c.PrintRaw(req)

			default:
				c.Error("Unknown print type")
			}

		default:
			c.Error("Unknown Action : " + req.Action)
		}
	}
}
