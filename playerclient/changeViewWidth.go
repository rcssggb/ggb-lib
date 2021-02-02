package playerclient

import (
	"fmt"

	"github.com/rcssggb/ggb-lib/rcsscommon"
)

// ChangeView changes the player's view mode. If synchronous view mode is on,
// Only the width parameter matters
func (c *Client) ChangeView(width rcsscommon.ViewWidth, quality rcsscommon.ViewQuality) string {
	changeViewStr := fmt.Sprintf("(change_view %s %s)", width, quality)
	c.cmdChannel <- changeViewStr
	return changeViewStr
}
