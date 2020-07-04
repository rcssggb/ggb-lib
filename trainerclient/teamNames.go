package trainerclient

/*
TeamNames provides information about the name of both teams
and which side they're playing on
*/
func (c *Client) TeamNames() string {
	teamNamesString := "(team_names)"
	c.cmdChannel <- teamNamesString
	return teamNamesString
}
