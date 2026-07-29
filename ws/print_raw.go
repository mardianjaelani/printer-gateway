package ws

import (
	"print-gateway/models"
	"print-gateway/services"
)

func (c *Client) PrintRaw(req Message) {

	service := services.NewPrintService()

	err := service.Print(models.PrintRequest{
		Printer: req.Printer,
		Copies:  req.Copies,
		Data:    req.Data,
	})

	if err != nil {
		c.Error(err.Error())
		return
	}

	c.Success("print", nil, "Print Raw berhasil")
}
