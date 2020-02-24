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

// Hear parses (hear TIME SENDER MESSAGE)
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

	// If message sender was a coach
	if strings.HasPrefix(msgParts[0], "online_coach") {
		data.Sender = msgParts[0]

		if len(msgParts) < 2 {
			err = fmt.Errorf("invalid online coach message: %s", m)
			return
		}

		data.Message = strings.Join(msgParts[1:len(msgParts)], " ")
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

	// At this point, message must be from unknown player (hear TIME DIR [our|opp] [UNUM] MESSAGE)
	if len(msgParts) < 2 {
		err = fmt.Errorf("invalid hear message: %s", m)
		return
	}

	// If unknown player is friendly, message is (hear TIME DIR our UNUM MESSAGE)
	if msgParts[1] == "our" {
		if len(msgParts) < 3 {
			err = fmt.Errorf("invalid unknown player message format: %s", m)
			return
		}
		data.Sender = strings.Join(msgParts[0:3], " ")
		data.Message = strings.Join(msgParts[3:len(msgParts)], " ")
		return
	}

	// If unknown player is an opponent, message is (hear TIME DIR opp MESSAGE)
	if msgParts[1] == "opp" {
		data.Sender = strings.Join(msgParts[0:2], " ")
		data.Message = strings.Join(msgParts[2:len(msgParts)], " ")
		return
	}

	err = fmt.Errorf("unknown hear format")
	return
}
