package lexer

import (
	"fmt"
	"strings"
)

func PlayerType(m string) (data map[string]string, err error) {
	trimmedMsg := m
	trimmedMsg = strings.TrimPrefix(trimmedMsg, "(player_type ")
	trimmedMsg = strings.TrimSuffix(trimmedMsg, ")\x00")

	data = make(map[string]string)

	for openIdx := strings.Index(trimmedMsg, "("); openIdx != -1; openIdx = strings.Index(trimmedMsg, "(") {
		closeIdx := strings.Index(trimmedMsg, ")")
		obj := trimmedMsg[openIdx+1 : closeIdx]

		s := strings.Split(obj, " ")
		if len(s) != 2 {
			err = fmt.Errorf("invalid number of player_type values in object %s seen in %s", obj, m)
		}

		data[s[0]] = s[1]
		trimmedMsg = trimmedMsg[closeIdx+1:]
	}

	return
}
