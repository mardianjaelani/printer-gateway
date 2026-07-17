package services

import (
	"encoding/base64"
	"errors"

	"print-gateway/models"
	"print-gateway/printer"
)

type PrintService struct{}

func NewPrintService() *PrintService {
	return &PrintService{}
}

func (s *PrintService) Print(req models.PrintRequest) error {

	data, err := base64.StdEncoding.DecodeString(req.Data)
	if err != nil {
		return errors.New("invalid base64 data")
	}

	if req.Copies <= 0 {
		req.Copies = 1
	}

	return printer.RawPrint(req.Printer, data, req.Copies)
}
