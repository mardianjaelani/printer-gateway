package services

import (
	"encoding/base64"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"print-gateway/config"
	"print-gateway/models"
)

func (s *PrintService) PrintPdf(req models.PrintPdfRequest) error {

	cfg, err := config.Load()
	if err != nil {
		return err
	}

	if _, err := os.Stat(cfg.Sumatra.Path); err != nil {
		return fmt.Errorf("SumatraPDF tidak ditemukan: %s", cfg.Sumatra.Path)
	}

	if req.Copies <= 0 {
		req.Copies = 1
	}

	// Decode Base64 PDF
	pdfBytes, err := base64.StdEncoding.DecodeString(req.Data)
	if err != nil {
		return fmt.Errorf("decode base64 gagal : %v", err)
	}

	// Folder temp
	tempDir := filepath.Join(os.TempDir(), "print-gateway")

	err = os.MkdirAll(tempDir, os.ModePerm)
	if err != nil {
		return err
	}

	pdfPath := filepath.Join(tempDir, "label.pdf")

	err = os.WriteFile(pdfPath, pdfBytes, 0644)
	if err != nil {
		return err
	}

	defer os.Remove(pdfPath)

	for i := 0; i < req.Copies; i++ {

		var cmd *exec.Cmd

		if req.Printer == "" {

			fmt.Println("Menggunakan default printer")

			cmd = exec.Command(
				cfg.Sumatra.Path,
				"-silent",
				"-print-to-default",
				pdfPath,
			)

		} else {

			fmt.Println("Menggunakan printer :", req.Printer)

			cmd = exec.Command(
				cfg.Sumatra.Path,
				"-silent",
				"-print-to",
				req.Printer,
				pdfPath,
			)

		}

		output, err := cmd.CombinedOutput()

		fmt.Println("Output :", string(output))
		fmt.Println("Error :", err)

		if err != nil {
			return fmt.Errorf("%v\n%s", err, string(output))
		}

	}

	return nil
}
