package ws

import (
	"log"

	"print-gateway/models"
	"print-gateway/services"
)

func (c *Client) PrintPdf(req Message) {

	service := services.NewPrintService()

	err := service.PrintPdf(models.PrintPdfRequest{
		Printer: req.Printer,
		Copies:  req.Copies,
		Data:    req.Data,
	})

	if err != nil {

		log.Println("PrintPdf:", err)

		c.Error(err.Error())

		return
	}

	c.Success(
		"print",
		nil,
		"Print PDF berhasil",
	)
}
