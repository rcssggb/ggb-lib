package lexer

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// ChangePlayerType parses (ok change_player_type TEAM UNUM TYPE)
func ChangePlayerType(m string) (teamName string, unum, ID int64, err error) {
	trimmedMsg := m
	trimmedMsg = strings.TrimPrefix(trimmedMsg, "(ok change_player_type ")
	trimmedMsg = strings.TrimSuffix(trimmedMsg, ")\x00")

	splitChangePlayerType := strings.Split(trimmedMsg, " ")
	if len(splitChangePlayerType) != 3 {
		err = errors.New("something went wrong with change_player_type parsing")
		return
	}

	teamName = splitChangePlayerType[0]

	unumStr := splitChangePlayerType[1]
	unum, err = strconv.ParseInt(unumStr, 10, 64)
	if err != nil {
		err = fmt.Errorf("error parsing string %s: %s", m, err)
		return
	}

	IDStr := splitChangePlayerType[2]

	ID, err = strconv.ParseInt(IDStr, 10, 64)
	if err != nil {
		err = fmt.Errorf("error parsing string %s: %s", m, err)
		return
	}

	return
}
