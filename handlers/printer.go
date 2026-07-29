package handlers

import (
	"net/http"
	"os/exec"

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

func GetPrinterConfig(c *gin.Context) {

	cmd := exec.Command(
		"powershell",
		"-Command",
		"Get-PrintConfiguration -PrinterName ((Get-Printer | Where Default).Name) | ConvertTo-Json",
	)

	out, err := cmd.Output()

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Data(
		200,
		"application/json",
		out,
	)
}
