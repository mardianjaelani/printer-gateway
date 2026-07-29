package ws

import (
	"log"

	"print-gateway/printer"
)

// Mengirim printer default
func (c *Client) DefaultPrinter() {

	p, err := printer.DefaultPrinter()
	if err != nil {
		log.Println("GetDefaultPrinter Error:", err)
		c.Error(err.Error())
		return
	}

	c.Success(
		"defaultPrinter",
		p,
		"",
	)
}
