package playerclient

import "github.com/rcssggb/ggb-lib/playerclient/actions"

/*
Kick accelerates the ball with the given Power in the given Direction. The
direction is relative to the the Direction of the body of the player and the
power is again between minpower and maxpower.
*/
func (c *Client) Kick(time int, power, direction float64) string {
	kickStr := actions.Kick(power, direction)
	c.cmdChannel <- command{
		time:      time,
		cmdString: kickStr,
	}
	return kickStr
}
