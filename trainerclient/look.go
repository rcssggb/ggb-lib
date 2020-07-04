package trainerclient

/*
Look provides information about the positions of the left and right
goals, the ball and all active players
*/
func (c *Client) Look() string {
	lookStr := "(look)"
	c.cmdChannel <- command{
		time:      c.currentTime,
		cmdString: lookStr,
	}
	return lookStr
}
