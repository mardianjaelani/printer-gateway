package models

type PrintPdfRequest struct {
	Printer string `json:"printer"`
	Copies  int    `json:"copies"`
	Job     string `json:"job"`
	Data    string `json:"data"` // base64 pdf
}
