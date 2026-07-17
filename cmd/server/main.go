package main

import (
	"net/http"

	"print-gateway/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/api/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "online",
		})
	})

	r.GET("/api/printers", handlers.GetPrinters)
	r.POST("/api/print/text", handlers.PrintText)

	r.Run(":8080")
}
