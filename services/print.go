package services

import (
	"encoding/base64"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"print-gateway/config"
	"print-gateway/models"
)

type PrintService struct{}

// =====================================================
// CONSTRUCTOR
// =====================================================

func NewPrintService() *PrintService {
	return &PrintService{}
}

// =====================================================
// PRINT RAW / TEXT
// =====================================================

func (s *PrintService) Print(req models.PrintRequest) error {

	if req.Printer == "" {
		return fmt.Errorf("printer wajib diisi")
	}

	if req.Data == "" {
		return fmt.Errorf("data print kosong")
	}

	// =================================================
	// Masukkan implementasi raw printing Anda di sini.
	// =================================================

	fmt.Println("================================")
	fmt.Println("PRINT RAW")
	fmt.Println("Printer :", req.Printer)
	fmt.Println("================================")

	return nil
}

// =====================================================
// PRINT PDF
// =====================================================

func (s *PrintService) PrintPdf(
	req models.PrintPdfRequest,
) error {

	// =================================================
	// LOAD CONFIG
	// =================================================

	cfg, err := config.Load()

	if err != nil {
		return fmt.Errorf(
			"gagal load config: %v",
			err,
		)
	}

	// =================================================
	// CHECK SUMATRA
	// =================================================

	if _, err := os.Stat(cfg.Sumatra.Path); err != nil {

		return fmt.Errorf(
			"SumatraPDF tidak ditemukan: %s",
			cfg.Sumatra.Path,
		)
	}

	// =================================================
	// COPIES
	// =================================================

	if req.Copies <= 0 {
		req.Copies = 1
	}

	// =================================================
	// DECODE BASE64
	// =================================================

	pdfBytes, err := base64.StdEncoding.DecodeString(req.Data)

	if err != nil {

		return fmt.Errorf(
			"decode base64 gagal: %v",
			err,
		)
	}

	if len(pdfBytes) == 0 {
		return fmt.Errorf("PDF kosong")
	}

	// =================================================
	// TEMP DIRECTORY
	// =================================================

	tempDir := filepath.Join(
		os.TempDir(),
		"print-gateway",
	)

	if err := os.MkdirAll(
		tempDir,
		os.ModePerm,
	); err != nil {

		return fmt.Errorf(
			"gagal membuat folder temporary: %v",
			err,
		)
	}

	// =================================================
	// UNIQUE FILE
	// =================================================

	pdfPath := filepath.Join(
		tempDir,
		fmt.Sprintf(
			"print_%d.pdf",
			time.Now().UnixNano(),
		),
	)

	// =================================================
	// WRITE PDF
	// =================================================

	if err := os.WriteFile(
		pdfPath,
		pdfBytes,
		0644,
	); err != nil {

		return fmt.Errorf(
			"gagal menyimpan PDF: %v",
			err,
		)
	}

	defer func() {

		if err := os.Remove(pdfPath); err != nil {
			fmt.Println(
				"Gagal menghapus temporary PDF:",
				err,
			)
		}

	}()

	// =================================================
	// PRINT
	// =================================================

	for i := 0; i < req.Copies; i++ {

		fmt.Println("")
		fmt.Println("================================")
		fmt.Println("PRINT PDF")
		fmt.Println("================================")

		if req.Printer == "" {

			fmt.Println(
				"Printer : DEFAULT WINDOWS",
			)

		} else {

			fmt.Println(
				"Printer :",
				req.Printer,
			)
		}

		fmt.Println(
			"Copy    :",
			i+1,
			"/",
			req.Copies,
		)

		fmt.Println(
			"Scale   : 100% / NOSCALE",
		)

		fmt.Println(
			"PDF     :",
			pdfPath,
		)

		fmt.Println("================================")

		var cmd *exec.Cmd

		// =================================================
		// DEFAULT PRINTER
		// =================================================

		if req.Printer == "" {

			cmd = exec.Command(
				cfg.Sumatra.Path,

				"-silent",

				"-print-to-default",

				"-print-settings",
				"noscale",

				pdfPath,
			)

		} else {

			// =================================================
			// SPECIFIC PRINTER
			// =================================================

			cmd = exec.Command(
				cfg.Sumatra.Path,

				"-silent",

				"-print-to",
				req.Printer,

				"-print-settings",
				"noscale",

				pdfPath,
			)
		}

		// =================================================
		// EXECUTE
		// =================================================

		output, err := cmd.CombinedOutput()

		fmt.Println(
			"Sumatra Output:",
			string(output),
		)

		if err != nil {

			fmt.Println(
				"Sumatra Error:",
				err,
			)

			return fmt.Errorf(
				"gagal print PDF: %v\n%s",
				err,
				string(output),
			)
		}

		fmt.Println(
			"Print berhasil dikirim",
		)

		// Beri sedikit waktu agar spooler
		// menerima job sebelum copy berikutnya.
		if i < req.Copies-1 {
			time.Sleep(300 * time.Millisecond)
		}
	}

	return nil
}
