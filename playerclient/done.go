package playerclient

// DoneSynch send a "(done)" message to the server
func (c *Client) DoneSynch() {
	c.cmdChannel <- "(done)"
}
