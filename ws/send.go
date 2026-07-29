package ws

import (
	"log"
)

// Send mengirim response ke client melalui WebSocket
func (c *Client) Send(resp Response) error {

	if err := c.Conn.WriteJSON(resp); err != nil {
		log.Println("WebSocket Send Error:", err)
		return err
	}

	return nil
}

// Error mengirim response error
func (c *Client) Error(message string) {

	_ = c.Send(Response{
		Success: false,
		Action:  "error",
		Message: message,
	})
}

// Success mengirim response sukses
func (c *Client) Success(action string, data interface{}, message string) {

	_ = c.Send(Response{
		Success: true,
		Action:  action,
		Message: message,
		Data:    data,
	})
}
