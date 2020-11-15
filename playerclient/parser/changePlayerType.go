package parser

import (
	"fmt"
	"strconv"
	"strings"
)

// ChangePlayerTypeData contains the parsed change player type instructions
type ChangePlayerTypeData struct {
	Unum       int
	PlayerType int
}

// ChangePlayerType parses (change_player_type UNUM [NEW_TYPE])
// If new player type is not shown, it means an opponent changed its player type
func ChangePlayerType(m string) (data ChangePlayerTypeData, err error) {
	trimmedMsg := m
	trimmedMsg = strings.TrimPrefix(trimmedMsg, "(change_player_type ")
	trimmedMsg = strings.TrimSuffix(trimmedMsg, ")\x00")

	msgParts := strings.Split(trimmedMsg, " ")

	var unum int64
	var pType int64
	if len(msgParts) == 2 { // If a teammate changed player type
		unum, err = strconv.ParseInt(msgParts[0], 10, 64)
		if err != nil {
			err = fmt.Errorf("failed to parse unum: %s", msgParts[0])
		}

		pType, err = strconv.ParseInt(msgParts[1], 10, 64)
		if err != nil {
			err = fmt.Errorf("failed to parse player type: %s", msgParts[1])
		}
	} else if len(msgParts) == 1 { // If an opponent changed player type
		unum, err = strconv.ParseInt(msgParts[0], 10, 64)
		if err != nil {
			err = fmt.Errorf("failed to parse unum: %s", msgParts[0])
		}

		pType = -1 // this indicates an opponent player changed player type
	} else {
		err = fmt.Errorf("change_player_type must have 1 or 2 arguments")
		return
	}
	data.Unum = (int)(unum)
	data.PlayerType = (int)(pType)
	return
}
