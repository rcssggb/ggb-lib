package actions

import (
	"fmt"

	"github.com/rcssggb/ggb-lib/rcsscommon/field"
)

// Move constructs the string to send a "move" command to the server
func Move(x, y float64) string {
	// Normalize x value
	if x < field.MinX {
		x = field.MinX
	}
	if x > field.MaxX {
		x = field.MaxX
	}

	// Normalize y value
	if y < field.MinY {
		y = field.MinY
	}
	if y > field.MaxY {
		y = field.MaxY
	}

	return fmt.Sprintf("(move %.1f %.1f)", x, y)
}
