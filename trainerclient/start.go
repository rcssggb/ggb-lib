package trainerclient

/*
Start starts the match with play-mode 'kick_off_l'
*/
func (c *Client) Start() string {
	startString := "(start)"
	c.cmdChannel <- command{
		time:      c.currentTime,
		cmdString: startString,
	}
	return startString
}
