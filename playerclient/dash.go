package playerclient

import "github.com/rcssggb/ggb-lib/playerclient/actions"

/*
Dash accelerates the player in the direction specified. The Power is between
minpower (used value: −100) and maxpower (used value: 100).
*/
func (c *Client) Dash(power, direction float64) string {
	dashStr := actions.Dash(power, direction)
	c.cmdChannel <- dashStr
	return dashStr
}
