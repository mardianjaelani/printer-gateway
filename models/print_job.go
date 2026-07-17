package models

import "time"

type PrintJob struct {
	ID        string
	Printer   string
	Data      []byte
	Copies    int
	CreatedAt time.Time
	Status    string
}
