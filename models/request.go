package models

type PrintTextRequest struct {
	Printer string `json:"printer"`
	Text    string `json:"text"`
}
