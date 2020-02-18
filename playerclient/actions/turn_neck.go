package actions

import (
	"fmt"
)

// TurnNeck constructs the string to send a "turn_neck" command to the server
func TurnNeck(moment float64) string {
	return fmt.Sprintf("(turn_neck %.3f)\x00", moment)
}
