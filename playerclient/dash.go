package playerclient

import "github.com/rcssggb/ggb-lib/playerclient/actions"

/*
Dash accelerates the player in the direction of its body (not direction of
the current speed). The Power is between minpower (used value: −100) and
maxpower (used value: 100).
*/
func (c *Client) Dash(time int, power, direction float64) {
	dashStr := actions.Dash(power, direction)
	c.cmdChannel <- command{
		time:      time,
		cmdString: dashStr,
	}
}
