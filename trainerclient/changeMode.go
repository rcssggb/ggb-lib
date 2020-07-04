package trainerclient

import (
	"github.com/rcssggb/ggb-lib/rcsscommon"
	"github.com/rcssggb/ggb-lib/trainerclient/actions"
)

/*
ChangeMode changes the game mode
*/
func (c *Client) ChangeMode(mode rcsscommon.PlayModeID) string {
	changeModeString := actions.ChangeMode(mode)
	c.cmdChannel <- command{
		time:      c.currentTime,
		cmdString: changeModeString,
	}
	return changeModeString
}
