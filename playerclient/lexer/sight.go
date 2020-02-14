package lexer

import (
	"fmt"
	"strconv"
	"strings"
)

// SightSymbols contains
type SightSymbols struct {
	Time       int
	ObjMap     map[string][]string
	PlayersMap SightPlayersSymbols
}

type SightPlayersSymbols struct {
	KnownPlayersMap     map[string][]string
	KnownTeamPlayersMap map[string][][]string
	UnknownPlayers      [][]string
}

// Sight parses (see 0 ((f r t) 55.7 3) ...
func Sight(m string) (data *SightSymbols, err error) {
	trimmedMsg := m
	trimmedMsg = strings.TrimPrefix(trimmedMsg, "(see ")
	trimmedMsg = strings.TrimSuffix(trimmedMsg, ")")

	timeEnd := strings.Index(trimmedMsg, " ")
	if timeEnd == -1 {
		// If player sees nothing, message will be `(see <time>)`
		timeEnd = strings.Index(trimmedMsg, ")")
	}
	timeStr := string(trimmedMsg[0:timeEnd])
	time, err := strconv.ParseInt(timeStr, 10, 64)
	if err != nil {
		err = fmt.Errorf("error parsing string %s: %s", m, err)
		return
	}

	data = &SightSymbols{
		Time:   int(time),
		ObjMap: make(map[string][]string),
		PlayersMap: SightPlayersSymbols{
			KnownPlayersMap:     make(map[string][]string),
			KnownTeamPlayersMap: make(map[string][][]string),
			UnknownPlayers:      [][]string{},
		},
	}

	trimmedMsg = trimmedMsg[timeEnd+1:]

	for openIdx := strings.Index(trimmedMsg, "(("); openIdx != -1; openIdx = strings.Index(trimmedMsg, "((") {
		closeIdx := strings.Index(trimmedMsg, ")")
		objName := trimmedMsg[openIdx+2 : closeIdx]
		trimmedMsg = trimmedMsg[closeIdx+2 : len(trimmedMsg)]

		closeIdx = strings.Index(trimmedMsg, ")")
		params := trimmedMsg[:closeIdx]
		trimmedMsg = trimmedMsg[closeIdx+1 : len(trimmedMsg)]

		splitParam := strings.Split(params, " ")

		if len(splitParam) != 2 &&
			len(splitParam) != 4 &&
			len(splitParam) != 6 {
			err = fmt.Errorf("invalid number of sight values in object %s seen in %s", objName, m)
			return
		}

		// If seen object is a player
		if objName[0] == 'p' {
			objNameParts := strings.Split(objName, " ")
			switch len(objNameParts) {
			case 3:
				data.PlayersMap.KnownPlayersMap[objName] = splitParam
			case 2:
				data.PlayersMap.KnownTeamPlayersMap[objNameParts[1]] = append(data.PlayersMap.KnownTeamPlayersMap[objNameParts[1]], splitParam)
			case 1:
				data.PlayersMap.UnknownPlayers = append(data.PlayersMap.UnknownPlayers, splitParam)
			default:
				err = fmt.Errorf("invalid object name %s seen in %s", objName, m)
				return
			}
			continue
		}

		data.ObjMap[objName] = splitParam
	}

	return
}
