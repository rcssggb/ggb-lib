package playerclient

import "github.com/rcssggb/ggb-lib/playerclient/parser"

import "github.com/rcssggb/ggb-lib/rcsscommon"

// See returns current sightData object
func (c *Client) See() parser.SightData {
	return c.sightData
}

// TeamSide returns player's team side on the field
func (c *Client) TeamSide() rcsscommon.SideType {
	return c.teamSide
}

// TeamName returns player's team name
func (c *Client) TeamName() string {
	return c.teamName
}

// Shirt returns player's shirt number
func (c *Client) Shirt() int {
	return c.shirtNum
}

// PlayMode returns current play mode
func (c *Client) PlayMode() string {
	return c.playMode
}