package lexer

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rcssggb/ggb-lib/trainerclient/types"
)

// Look parses (see_global ... and (ok look ...
func Look(m string) (data *types.GlobalPositionsSymbols, err error) {
	trimmedMsg := m
	trimmedMsg = strings.TrimPrefix(trimmedMsg, "(see_global ")
	trimmedMsg = strings.TrimPrefix(trimmedMsg, "(ok look ")
	trimmedMsg = strings.TrimSuffix(trimmedMsg, "))\x00")

	timeEnd := strings.Index(trimmedMsg, " ")
	if timeEnd == -1 {
		// As a safety measure
		timeEnd = strings.Index(trimmedMsg, ")")
	}
	timeStr := string(trimmedMsg[0:timeEnd])
	time, err := strconv.ParseInt(timeStr, 10, 64)
	if err != nil {
		err = fmt.Errorf("error parsing string %s: %s", m, err)
		return
	}

	data = &types.GlobalPositionsSymbols{
		Time:    int(time),
		Objects: make(map[string][]string),
	}

	trimmedMsg = trimmedMsg[timeEnd+1:]
	trimmedMsg = strings.TrimPrefix(trimmedMsg, "(")

	objects := strings.Split(trimmedMsg, ") (")

	for _, obj := range objects {
		open := strings.Index(obj, "(")
		close := strings.Index(obj, ")")
		objName := obj[open+1 : close]
		obj = obj[close+2:]

		objData := strings.Split(obj, " ")

		data.Objects[objName] = objData
	}

	return
}
