package models

type PrintRequest struct {
	Printer string `json:"printer"`
	Data    string `json:"data"`
}

type PrintPdfRequest struct {
	Printer string `json:"printer"`
	Data    string `json:"data"`
	Copies  int    `json:"copies"`
}

type ApiResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
