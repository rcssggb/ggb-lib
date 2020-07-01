package trainerclient

import (
	"fmt"

	"github.com/rcssggb/ggb-lib/rcsscommon"
)

// TODO: implement direction and velocity (which are optional)

/*
MovePlayer moves the player identified by teamName and unum to a (x,y) location
with direction 'dir' and velocity (dx, dy)
*/
func (c *Client) MovePlayer(teamName string, unum int, x, y float64) string {
	playerObject := fmt.Sprintf("(player %s %d)", teamName, unum)

	return c.move(playerObject, x, y)
}

/*
MoveBall moves the ball to a (x,y) location
with direction 'dir' and velocity (dx, dy)
*/
func (c *Client) MoveBall(x, y float64) string {
	ballObject := fmt.Sprint("(ball)")

	return c.move(ballObject, x, y)
}

func (c *Client) move(object string, x, y float64) string {
	// Normalize x value
	if x < rcsscommon.FieldMinX {
		x = rcsscommon.FieldMinX
	}
	if x > rcsscommon.FieldMaxX {
		x = rcsscommon.FieldMaxX
	}

	// Normalize y value
	if y < rcsscommon.FieldMinY {
		y = rcsscommon.FieldMinY
	}
	if y > rcsscommon.FieldMaxY {
		y = rcsscommon.FieldMaxY
	}

	moveStr := fmt.Sprintf("(move %s %.1f %.1f)", object, x, y)
	c.cmdChannel <- command{
		cmdString: moveStr,
	}
	return moveStr
}
