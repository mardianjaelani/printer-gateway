package ws

// Request dari Client (Vue) ke Gateway
type Message struct {
	Action string `json:"action"`

	// Print
	Printer string `json:"printer,omitempty"`
	Copies  int    `json:"copies,omitempty"`
	Job     string `json:"job,omitempty"`
	Type    string `json:"type,omitempty"`
	Data    string `json:"data,omitempty"`

	// Authentication (opsional)
	Token string `json:"token,omitempty"`
}

// Response dari Gateway ke Client
type Response struct {
	Success bool `json:"success"`

	Action  string `json:"action"`
	Message string `json:"message,omitempty"`

	Data interface{} `json:"data,omitempty"`
}
