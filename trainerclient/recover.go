package trainerclient

/*
Recover resets stamina, recovery, effort and hear capacity to the values
at the beggining of the game
*/
func (c *Client) Recover() string {
	recoverString := "(recover)"
	c.cmdChannel <- command{
		cmdString: recoverString,
	}
	return recoverString
}
