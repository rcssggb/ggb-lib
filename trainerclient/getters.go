package trainerclient

import (
	"github.com/rcssggb/ggb-lib/rcsscommon"
	"github.com/rcssggb/ggb-lib/trainerclient/types"
)

// GlobalPositions returns latest global positions object
func (c *Client) GlobalPositions() types.GlobalPositions {
	return c.globalPositions
}

// LeftTeamName returns left team name
func (c *Client) LeftTeamName() string {
	return c.lTeamName
}

// RightTeamName returns right team name
func (c *Client) RightTeamName() string {
	return c.rTeamName
}

// PlayMode returns current play mode
func (c *Client) PlayMode() rcsscommon.PlayModeID {
	return c.playMode
}

// Time returns current simulation time step
func (c *Client) Time() int {
	return c.currentTime
}

// ServerParams returns rcssserver params object
func (c *Client) ServerParams() rcsscommon.ServerParams {
	return c.serverParams
}
