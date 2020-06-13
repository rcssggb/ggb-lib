package trainerclient

/*
TeamNames provides information about the name of both teams
and which side they're playing on
*/
func (c *Client) TeamNames() string {
	lookStr := "(team_names)"
	c.cmdChannel <- command{
		cmdString: lookStr,
	}
	return lookStr
}
