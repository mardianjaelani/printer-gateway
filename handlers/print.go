package handlers

import (
	"net/http"

	"print-gateway/models"
	"print-gateway/services"

	"github.com/gin-gonic/gin"
)

type PrintHandler struct {
	service *services.PrintService
}

func NewPrintHandler() *PrintHandler {
	return &PrintHandler{
		service: services.NewPrintService(),
	}
}

func (h *PrintHandler) Print(c *gin.Context) {

	var req models.PrintRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ApiResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if err := h.service.Print(req); err != nil {
		c.JSON(http.StatusInternalServerError, models.ApiResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.ApiResponse{
		Success: true,
		Message: "Print job accepted",
	})
}
