package ws

import "print-gateway/printer"

func (c *Client) DefaultPrinter() {

	p := printer.GetDefaultPrinter()

	c.Send(Response{
		Success: true,
		Action:  "defaultPrinter",
		Data:    p,
	})

}
