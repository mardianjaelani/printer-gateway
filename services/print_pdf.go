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

	cfg, err := config.Load()
	if err != nil {
		return err
	}

	// Check SumatraPDF
	if _, err := os.Stat(cfg.Sumatra.Path); err != nil {
		return fmt.Errorf(
			"SumatraPDF tidak ditemukan: %s",
			cfg.Sumatra.Path,
		)
	}

	// Default copies
	if req.Copies <= 0 {
		req.Copies = 1
	}

	// Decode Base64 PDF
	pdfBytes, err := base64.StdEncoding.DecodeString(req.Data)
	if err != nil {
		return fmt.Errorf(
			"decode base64 gagal: %v",
			err,
		)
	}

	// Temp directory
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

	// Unique PDF filename
	pdfPath := filepath.Join(
		tempDir,
		fmt.Sprintf(
			"print_%d.pdf",
			time.Now().UnixNano(),
		),
	)

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

	defer os.Remove(pdfPath)

	// ==============================
	// PRINT
	// ==============================

	for i := 0; i < req.Copies; i++ {

		var cmd *exec.Cmd

		if req.Printer == "" {

			fmt.Println("================================")
			fmt.Println("PRINT PDF")
			fmt.Println("Printer : DEFAULT WINDOWS")
			fmt.Println("Copy    :", i+1)
			fmt.Println("Scale   : NONE")
			fmt.Println("================================")

			cmd = exec.Command(
				cfg.Sumatra.Path,
				"-silent",
				"-print-to-default",
				"-print-settings",
				"noscale,paper=auto",
				pdfPath,
			)

		} else {

			fmt.Println("================================")
			fmt.Println("PRINT PDF")
			fmt.Println("Printer :", req.Printer)
			fmt.Println("Copy    :", i+1)
			fmt.Println("Scale   : NONE")
			fmt.Println("================================")

			cmd = exec.Command(
				cfg.Sumatra.Path,
				"-silent",
				"-print-to",
				req.Printer,
				"-print-settings",
				"noscale,paper=auto",
				pdfPath,
			)
		}

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
