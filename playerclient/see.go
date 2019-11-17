package playerclient

import "github.com/rcssggb/ggb-lib/playerclient/parser"

// See returns current sightData object
func (c *Client) See() parser.SightData {
	return c.sightData
}
