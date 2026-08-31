package handlers

import (
	"encoding/base64"
	"net/http"
	"strings"

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

// =====================================================
// PRINT RAW / TEXT
// =====================================================

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

// =====================================================
// PRINT PDF
// =====================================================

func (h *PrintHandler) PrintPdf(c *gin.Context) {

	var req models.PrintPdfRequest

	// =================================================
	// BIND JSON
	// =================================================

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, models.ApiResponse{
			Success: false,
			Message: "Request tidak valid: " + err.Error(),
		})

		return
	}

	// =================================================
	// VALIDASI PRINTER
	// =================================================

	req.Printer = strings.TrimSpace(req.Printer)

	// Printer boleh kosong karena berarti
	// menggunakan default printer Windows.

	// =================================================
	// VALIDASI DATA PDF
	// =================================================

	if strings.TrimSpace(req.Data) == "" {

		c.JSON(http.StatusBadRequest, models.ApiResponse{
			Success: false,
			Message: "Data PDF kosong",
		})

		return
	}

	// =================================================
	// REMOVE DATA URI PREFIX
	//
	// Contoh:
	//
	// data:application/pdf;base64,JVBERi0x...
	//
	// menjadi:
	//
	// JVBERi0x...
	// =================================================

	if strings.Contains(req.Data, ",") {

		parts := strings.SplitN(
			req.Data,
			",",
			2,
		)

		if len(parts) == 2 {
			req.Data = parts[1]
		}
	}

	req.Data = strings.TrimSpace(req.Data)

	// =================================================
	// VALIDASI BASE64
	// =================================================

	pdfBytes, err := base64.StdEncoding.DecodeString(req.Data)

	if err != nil {

		c.JSON(http.StatusBadRequest, models.ApiResponse{
			Success: false,
			Message: "Base64 PDF tidak valid: " + err.Error(),
		})

		return
	}

	// =================================================
	// VALIDASI PDF HEADER
	// =================================================

	if len(pdfBytes) < 4 ||
		string(pdfBytes[:4]) != "%PDF" {

		c.JSON(http.StatusBadRequest, models.ApiResponse{
			Success: false,
			Message: "Data bukan file PDF yang valid",
		})

		return
	}

	// =================================================
	// NORMALIZE BASE64
	//
	// Setelah validasi, encode ulang agar service
	// mendapatkan base64 yang bersih.
	// =================================================

	req.Data = base64.StdEncoding.EncodeToString(pdfBytes)

	// =================================================
	// COPIES
	// =================================================

	if req.Copies <= 0 {
		req.Copies = 1
	}

	// =================================================
	// PRINT
	// =================================================

	if err := h.service.PrintPdf(req); err != nil {

		c.JSON(http.StatusInternalServerError, models.ApiResponse{
			Success: false,
			Message: err.Error(),
		})

		return
	}

	// =================================================
	// SUCCESS
	// =================================================

	c.JSON(http.StatusOK, models.ApiResponse{
		Success: true,
		Message: "PDF berhasil dikirim ke printer",
	})
}
