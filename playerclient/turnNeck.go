package playerclient

import "github.com/rcssggb/ggb-lib/playerclient/actions"

/*
TurnNeck can be sent (and will be executed) each cycle independently, along with
other action commands. The neck will rotate with the given Angle relative to previous
Angle. Note that the resulting neck angle will be between minneckang (default: −90)
and maxneckang (default: 90) relative to the player’s body direction.
*/
func (c *Client) TurnNeck(time int, moment float64) {
	turnNeckStr := actions.TurnNeck(moment)
	c.cmdChannel <- command{
		time:      time,
		cmdString: turnNeckStr,
	}
}
