package lexer

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// CheckBall parses (ok check_ball ...
func CheckBall(m string) (time int, ballInfo string, err error) {
	trimmedMsg := m
	trimmedMsg = strings.TrimPrefix(trimmedMsg, "(ok check_ball ")
	trimmedMsg = strings.TrimSuffix(trimmedMsg, ")\x00")

	splitCheckBall := strings.Split(trimmedMsg, " ")
	if len(splitCheckBall) != 2 {
		err = errors.New("something went wrong with check_ball parsing")
		return
	}

	timeStr := splitCheckBall[0]
	t, err := strconv.ParseInt(timeStr, 10, 64)
	if err != nil {
		err = fmt.Errorf("error parsing string %s: %s", m, err)
		return
	}

	time = int(t)
	ballInfo = splitCheckBall[1]

	return
}
