package handlers

import (
	"net/http"

	"print-gateway/models"

	"github.com/gin-gonic/gin"
)

func (h *PrintHandler) PrintPdf(c *gin.Context) {

	var req models.PrintPdfRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ApiResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if err := h.service.PrintPdf(req); err != nil {
		c.JSON(http.StatusInternalServerError, models.ApiResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.ApiResponse{
		Success: true,
		Message: "Print PDF Success",
	})
}
