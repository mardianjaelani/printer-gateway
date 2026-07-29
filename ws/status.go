package ws

func (c *Client) Status() {

	c.Send(Response{
		Success: true,
		Action:  "status",
		Message: "Gateway Online",
	})

}
