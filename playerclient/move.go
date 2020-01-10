package playerclient

import "github.com/rcssggb/ggb-lib/playerclient/actions"

/*
Move can be executed only before kick off and after a goal. It moves the player
to the exact position of X (between −54 and 54) and Y (between −32 and 32) in
one simulation cycle. This is useful for before kick off arrangements.
*/
func (c *Client) Move(time int, x, y float64) string {
	moveStr := actions.Move(x, y)
	c.cmdChannel <- command{
		time:      time,
		cmdString: moveStr,
	}
	return moveStr
}
