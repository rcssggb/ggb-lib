package actions

import (
	"fmt"
)

// Dash constructs the string to send a "Dash" command to the server
func Dash(power, direction float64) string {
	// Normalize direction between -180º and 180º
	normDir := direction
	for normDir > 180.0 {
		normDir -= 360.0
	}
	for normDir < -180 {
		normDir += 360.0
	}

	return fmt.Sprintf("(dash %.3f %.3f)", power, normDir)
}
