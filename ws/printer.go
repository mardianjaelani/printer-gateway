package ws

import (
	"log"

	"print-gateway/printer"
)

// Mengirim daftar printer
func (c *Client) Printers() {

	printers, err := printer.GetPrinters()
	if err != nil {
		log.Println("GetPrinters Error:", err)
		c.Error(err.Error())
		return
	}

	c.Success(
		"printers",
		printers,
		"",
	)
}

// Mengirim printer default
func (c *Client) DefaultPrinter() {

	p, err := printer.GetDefaultPrinter()
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
