package playerclient

import "github.com/rcssggb/ggb-lib/playerclient/actions"

// Move converts commands into
func (c *Client) Move(time int, x, y float64) {
	moveStr := actions.Move(x, y)
	c.cmdChannel <- command{
		time:      time,
		cmdString: moveStr,
	}
}
