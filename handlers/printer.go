package handlers

import (
	"net/http"

	"print-gateway/models"

	"github.com/alexbrainman/printer"
	"github.com/gin-gonic/gin"
)

func GetPrinters(c *gin.Context) {

	names, err := printer.ReadNames()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	var printers []models.Printer

	for _, name := range names {
		printers = append(printers, models.Printer{
			Name: name,
		})
	}

	c.JSON(http.StatusOK, printers)
}
