package playerclient

// WaitSynch blocks until (think) message is received.
func (c *Client) WaitSynch() {
	<-c.thinkChan
}

// DoneSynch send a "(done)" message to the server
func (c *Client) DoneSynch() {
	c.cmdChannel <- "(done)"
}
