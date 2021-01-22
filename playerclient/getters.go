package playerclient

import (
	"github.com/rcssggb/ggb-lib/playerclient/parser"
	"github.com/rcssggb/ggb-lib/playerclient/types"
	"github.com/rcssggb/ggb-lib/rcsscommon"
)

// See returns current sightData object
func (c *Client) See() parser.SightData {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.sightData
}

// SenseBody returns current senseBodyData object
func (c *Client) SenseBody() parser.SenseBodyData {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.bodyData
}

// TeamSide returns player's team side on the field
func (c *Client) TeamSide() rcsscommon.SideType {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.teamSide
}

// TeamName returns player's team name
func (c *Client) TeamName() string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.teamName
}

// Shirt returns player's shirt number
func (c *Client) Shirt() int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.shirtNum
}

// PlayMode returns current play mode
func (c *Client) PlayMode() rcsscommon.PlayModeID {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.playMode
}

// Time returns current simulation time step
func (c *Client) Time() int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.currentTime
}

// ServerParams returns rcssserver params object
func (c *Client) ServerParams() rcsscommon.ServerParams {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.serverParams
}

// PlayerParams returns player params object
func (c *Client) PlayerParams() rcsscommon.PlayerParams {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.playerParams
}

// TeammateTypes returns teammateTypes map
func (c *Client) TeammateTypes() types.TeammateTypes {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	ret := make(types.TeammateTypes)
	for k, v := range c.teammateTypes {
		ret[k] = v
	}
	return ret
}

// TeamGoals returns player's own team goals
func (c *Client) TeamGoals() int64 {
	if c.teamSide == rcsscommon.LeftSide {
		return c.goalsL
	} else if c.teamSide == rcsscommon.RightSide {
		return c.goalsR
	}
	return -1
}

// OpponentGoals returns opponent's team goals
func (c *Client) OpponentGoals() int64 {
	if c.teamSide == rcsscommon.RightSide {
		return c.goalsL
	} else if c.teamSide == rcsscommon.LeftSide {
		return c.goalsR
	}
	return -1
}
