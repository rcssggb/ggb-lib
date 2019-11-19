package actions

import (
	"fmt"
)

// Dash constructs the string to send a "Dash" command to the server
func Dash(power float64) string {
	return fmt.Sprintf("(dash %.3f)", power)
}
