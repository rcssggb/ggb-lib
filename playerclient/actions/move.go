package actions

import (
	"fmt"

	"github.com/rcssggb/ggb-lib/rcsscommon"
)

// Move constructs the string to send a "move" command to the server
func Move(x, y float64) string {
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

	return fmt.Sprintf("(move %.1f %.1f)\x00", x, y)
}
