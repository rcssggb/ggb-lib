package lexer

import (
	"fmt"
	"strconv"
	"strings"
)

// SightSymbols contains
type SightSymbols struct {
	Time     int
	SightMap map[string][]string
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
		Time:     int(time),
		SightMap: make(map[string][]string),
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
		}

		if objName[0] == 'p' {
			// for now we're skipping players, format will need to change a bit
			continue
		}
		data.SightMap[objName] = splitParam
	}

	return
}
