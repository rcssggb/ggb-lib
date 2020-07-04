package playerclient

import "github.com/rcssggb/ggb-lib/playerclient/actions"

/*
Turn will turn the player’s body direction Moment degrees relative to the
current direction. The moment is in degrees from −180 to 180.
*/
func (c *Client) Turn(moment float64) string {
	turnStr := actions.Turn(moment)
	c.cmdChannel <- turnStr
	return turnStr
}
