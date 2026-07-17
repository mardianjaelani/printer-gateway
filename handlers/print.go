package handlers

import (
	"net/http"
	"print-gateway/models"

	"github.com/gin-gonic/gin"
)

func PrintText(c *gin.Context) {

	var req models.PrintTextRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	// sementara hanya tampilkan request
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"printer": req.Printer,
		"text":    req.Text,
	})
}
