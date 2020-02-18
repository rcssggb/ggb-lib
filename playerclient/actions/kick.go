package actions

import (
	"fmt"
)

// Kick constructs the string to send a "Kick" command to the server
func Kick(pwr, dir float64) string {
	return fmt.Sprintf("(kick %.3f %f)\x00", pwr, dir)
}
