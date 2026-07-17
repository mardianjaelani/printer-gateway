package models

type PrintRequest struct {
	Printer string `json:"printer"`
	Data    string `json:"data"`
	Job     string `json:"job"`
	Copies  int    `json:"copies"`
}
