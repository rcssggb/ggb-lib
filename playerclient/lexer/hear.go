package lexer

import (
	"fmt"
	"strconv"
	"strings"
)

// HearSymbols contains
type HearSymbols struct {
	Time    int
	Sender  string
	Message string
}

// Hear parses (see 0 ((f r t) 55.7 3) ...
func Hear(m string) (data *HearSymbols, err error) {
	trimmedMsg := m
	trimmedMsg = strings.TrimPrefix(trimmedMsg, "(hear ")
	trimmedMsg = strings.TrimSuffix(trimmedMsg, ")\x00")

	timeEnd := strings.Index(trimmedMsg, " ")
	if timeEnd == -1 {
		// Fail-safe carried from lexer.Sight, possibly useless but it's better to be safe than sorry
		timeEnd = strings.Index(trimmedMsg, ")")
	}
	timeStr := string(trimmedMsg[0:timeEnd])
	time, err := strconv.ParseInt(timeStr, 10, 64)
	if err != nil {
		err = fmt.Errorf("error parsing string %s: %s", m, err)
		return
	}

	data = &HearSymbols{
		Time: int(time),
	}

	trimmedMsg = trimmedMsg[timeEnd+1:]

	msgParts := strings.Split(trimmedMsg, " ")

	// If message was sent by referee
	if msgParts[0] == "referee" {
		data.Sender = msgParts[0]

		if len(msgParts) != 2 {
			err = fmt.Errorf("invalid referee message format: %s", m)
			return
		}

		data.Message = msgParts[1]
		return
	}

	// If message was sent by the player itself
	if msgParts[0] == "self" {
		data.Sender = msgParts[0]

		if len(msgParts) < 2 {
			err = fmt.Errorf("invalid self message format: %s", m)
			return
		}

		data.Message = strings.Join(msgParts[1:len(msgParts)], " ")
		return
	}

	// If message was sent by another player
	if trimmedMsg[0] == '(' {
		closeIdx := strings.Index(trimmedMsg, ")")
		if closeIdx == -1 {
			err = fmt.Errorf("invalid player message format: %s", m)
		}

		data.Sender = trimmedMsg[1:closeIdx]
		data.Message = trimmedMsg[closeIdx+2 : len(trimmedMsg)]
		return
	}

	// TODO: If message sender was a coach

	// TODO: If message sender is just a direction

	return
}
