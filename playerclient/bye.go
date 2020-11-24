package playerclient

/*
Bye disconnects from the server.
*/
func (c *Client) Bye() string {
	c.cmdChannel <- "(bye)"
	return "(bye)"
}
