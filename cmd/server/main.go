package main

import (
	"fmt"
	"net/http"
	"time"

	"print-gateway/config"
	"print-gateway/handlers"
	"print-gateway/ws"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	// CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins: cfg.Cors.AllowOrigins,
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// ==========================
	// WebSocket Hub
	// ==========================
	hub := ws.NewHub()
	go hub.Run()

	// ==========================
	// HTTP API
	// ==========================

	r.GET("/api/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "online",
		})
	})

	r.GET("/api/printers", handlers.GetPrinters)

	printHandler := handlers.NewPrintHandler()
	r.POST("/api/print", printHandler.Print)
	r.POST("/api/printPdf", printHandler.PrintPdf)

	// ==========================
	// WebSocket
	// ==========================
	r.GET("/ws", handlers.HandleWS(hub))

	// ==========================
	// Run Server
	// ==========================

	r.Run(fmt.Sprintf(":%d", cfg.Server.Port))
}
