package trainerclient

import (
	"fmt"

	"github.com/rcssggb/ggb-lib/rcsscommon"
)

/*
MovePlayer moves the player identified by teamName and unum to a (x,y) location
with direction 'dir' and velocity (dx, dy)
*/
func (c *Client) MovePlayer(teamName string, unum int, x, y, dir, dx, dy float64) string {
	playerObject := fmt.Sprintf("(player %s %d)", teamName, unum)

	return c.move(playerObject, false, x, y, dir, dx, dy)
}

/*
MoveBall moves the ball to a (x,y) location
with direction 'dir' and velocity (dx, dy)
*/
func (c *Client) MoveBall(x, y, dx, dy float64) string {
	ballObject := fmt.Sprint("(ball)")

	return c.move(ballObject, true, x, y, 0, dx, dy)
}

func (c *Client) move(object string, isBall bool, x, y, dir, dx, dy float64) string {
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

	moveStr := fmt.Sprintf("(move %s %.1f %.1f %.1f %.1f %.1f)", object, x, y, dir, dx, dy)
	c.cmdChannel <- command{
		time:      c.currentTime,
		cmdString: moveStr,
	}
	return moveStr
}
