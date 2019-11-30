package actions

import (
	"fmt"
)

// Turn constructs the string to send a "turn" command to the server
func Turn(moment float64) string {
	return fmt.Sprintf("(turn %.3f)", moment)
}
