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

func (s *PrintService) PrintPdf(req models.PrintPdfRequest) error {

	// ==============================
	// LOAD CONFIG
	// ==============================

	cfg, err := config.Load()
	if err != nil {
		return err
	}

	// ==============================
	// CHECK SUMATRA PDF
	// ==============================

	if _, err := os.Stat(cfg.Sumatra.Path); err != nil {
		return fmt.Errorf(
			"SumatraPDF tidak ditemukan: %s",
			cfg.Sumatra.Path,
		)
	}

	// ==============================
	// DEFAULT COPIES
	// ==============================

	if req.Copies <= 0 {
		req.Copies = 1
	}

	// ==============================
	// DECODE BASE64 PDF
	// ==============================

	pdfBytes, err := base64.StdEncoding.DecodeString(req.Data)
	if err != nil {
		return fmt.Errorf(
			"decode base64 gagal: %v",
			err,
		)
	}

	// ==============================
	// TEMP DIRECTORY
	// ==============================

	tempDir := filepath.Join(
		os.TempDir(),
		"print-gateway",
	)

	if err := os.MkdirAll(tempDir, os.ModePerm); err != nil {
		return fmt.Errorf(
			"gagal membuat folder temporary: %v",
			err,
		)
	}

	// ==============================
	// TEMP PDF FILE
	// ==============================

	pdfPath := filepath.Join(
		tempDir,
		fmt.Sprintf(
			"label_%d.pdf",
			time.Now().UnixNano(),
		),
	)

	if err := os.WriteFile(
		pdfPath,
		pdfBytes,
		0644,
	); err != nil {
		return fmt.Errorf(
			"gagal menyimpan PDF temporary: %v",
			err,
		)
	}

	// Hapus file setelah selesai
	defer os.Remove(pdfPath)

	// ==============================
	// PRINT
	// ==============================

	for i := 0; i < req.Copies; i++ {

		var cmd *exec.Cmd

		if req.Printer == "" {

			fmt.Println("--------------------------------")
			fmt.Println("PRINT PDF")
			fmt.Println("Printer : DEFAULT")
			fmt.Println("Copy    :", i+1)
			fmt.Println("File    :", pdfPath)
			fmt.Println("--------------------------------")

			cmd = exec.Command(
				cfg.Sumatra.Path,

				// Jangan tampilkan UI SumatraPDF
				"-silent",

				// Printer default Windows
				"-print-to-default",

				// Cetak ukuran asli PDF
				"-print-settings",
				"noscale",

				// File PDF
				pdfPath,
			)

		} else {

			fmt.Println("--------------------------------")
			fmt.Println("PRINT PDF")
			fmt.Println("Printer :", req.Printer)
			fmt.Println("Copy    :", i+1)
			fmt.Println("File    :", pdfPath)
			fmt.Println("--------------------------------")

			cmd = exec.Command(
				cfg.Sumatra.Path,

				// Jangan tampilkan UI SumatraPDF
				"-silent",

				// Printer tujuan
				"-print-to",
				req.Printer,

				// Cetak ukuran asli PDF
				"-print-settings",
				"noscale",

				// File PDF
				pdfPath,
			)
		}

		// ==============================
		// EXECUTE PRINT
		// ==============================

		output, err := cmd.CombinedOutput()

		fmt.Println("Sumatra Output:", string(output))
		fmt.Println("Sumatra Error :", err)

		if err != nil {
			return fmt.Errorf(
				"gagal print PDF: %v\n%s",
				err,
				string(output),
			)
		}
	}

	return nil
}
