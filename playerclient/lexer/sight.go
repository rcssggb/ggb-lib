package lexer

import (
	"fmt"
	"strconv"
	"strings"
)

// SightSymbols contains
type SightSymbols struct {
	Time    int
	ObjMap  map[string][]string
	Players SightPlayersSymbols
}

type SightPlayersSymbols struct {
	Known     map[string][]string
	KnownTeam map[string][][]string
	Unknown   [][]string
}

// Sight parses (see 0 ((f r t) 55.7 3) ...
func Sight(m string) (data *SightSymbols, err error) {
	trimmedMsg := m
	trimmedMsg = strings.TrimPrefix(trimmedMsg, "(see ")
	trimmedMsg = strings.TrimSuffix(trimmedMsg, ")\x00")

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
		Players: SightPlayersSymbols{
			Known:     make(map[string][]string),
			KnownTeam: make(map[string][][]string),
			Unknown:   [][]string{},
		},
	}

	trimmedMsg = trimmedMsg[timeEnd+1:]

	for openIdx := strings.Index(trimmedMsg, "(("); openIdx != -1; openIdx = strings.Index(trimmedMsg, "((") {
		closeIdx := strings.Index(trimmedMsg, ")")
		objName := trimmedMsg[openIdx+2 : closeIdx]
		trimmedMsg = trimmedMsg[closeIdx+2:]

		closeIdx = strings.Index(trimmedMsg, ")")
		params := trimmedMsg[:closeIdx]
		trimmedMsg = trimmedMsg[closeIdx+1:]

		splitParam := strings.Split(params, " ")

		if len(splitParam) > 7 {
			err = fmt.Errorf("invalid number of sight values in object %s seen in %s", objName, m)
			return
		}

		// If seen object is a player
		if objName[0] == 'p' {
			objNameParts := strings.Split(objName, " ")
			switch len(objNameParts) {
			case 3:
				data.Players.Known[objName] = splitParam
			case 2:
				data.Players.KnownTeam[objNameParts[1]] = append(data.Players.KnownTeam[objNameParts[1]], splitParam)
			case 1:
				data.Players.Unknown = append(data.Players.Unknown, splitParam)
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
