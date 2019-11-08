package lexer

import (
	"errors"
	"strings"
)

// ServerParamSymbols is the type to be returned by parse.ServerParamSymbols
type ServerParamSymbols map[string]string

// ServerParam parses (server_param ...
func ServerParam(m string) (sp ServerParamSymbols, err error) {
	sp = make(ServerParamSymbols)

	trimmedMsg := m
	trimmedMsg = strings.TrimPrefix(trimmedMsg, "(server_param ")
	trimmedMsg = strings.TrimSuffix(trimmedMsg, ")")

	for closeIdx := strings.Index(trimmedMsg, ")"); closeIdx != -1; closeIdx = strings.Index(trimmedMsg, ")") {
		currParam := trimmedMsg[1:closeIdx]
		trimmedMsg = trimmedMsg[closeIdx+1 : len(trimmedMsg)-1]
		splitParam := strings.Split(currParam, " ")
		if len(splitParam) != 2 {
			err = errors.New("something went wrong with server_param parsing")
			return
		}
		paramName := splitParam[0]
		paramValString := splitParam[1]

		sp[paramName] = paramValString
	}
	return
}
