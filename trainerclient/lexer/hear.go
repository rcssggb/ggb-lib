package lexer

import (
	"fmt"
	"strconv"
	"strings"
)

// HearSymbols contains
type HearSymbols struct {
	Sender  string
	Time    int
	Message string
}

// Hear parses (hear SENDER TIME MESSAGE)
func Hear(m string) (data *HearSymbols, err error) {
	trimmedMsg := m
	trimmedMsg = strings.TrimPrefix(trimmedMsg, "(hear ")
	trimmedMsg = strings.TrimSuffix(trimmedMsg, ")\x00")

	data = &HearSymbols{}

	msgParts := strings.Split(trimmedMsg, " ")

	// If message was sent by referee
	if msgParts[0] == "referee" {
		if len(msgParts) != 3 {
			err = fmt.Errorf("invalid referee message format: %s", m)
			return
		}

		data.Sender = msgParts[0]

		var time int64
		time, err = strconv.ParseInt(msgParts[1], 10, 64)
		if err != nil {
			err = fmt.Errorf("could not parse time in message %s: %s", m, err)
			return
		}
		data.Time = (int)(time)

		data.Message = msgParts[2]
		return
	}

	err = fmt.Errorf("unsupported hear message: %s", m)
	return
}
