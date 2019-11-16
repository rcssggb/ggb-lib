package lexer

import (
	"strconv"
	"strings"
)

type SenseBodySymbols struct {
	Time         int
	SenseBodyMap map[string][]string
}

func SenseBody(m string) (data *SenseBodySymbols, err error) {
	trimmedMsg := m
	trimmedMsg = strings.TrimPrefix(trimmedMsg, "(sense_body ")
	trimmedMsg = strings.TrimSuffix(trimmedMsg, ")")

	timeStr := string(trimmedMsg[0])
	time, err := strconv.ParseInt(timeStr, 10, 64)

	data = &SenseBodySymbols{
		Time:         int(time),
		SenseBodyMap: make(map[string][]string),
	}

	trimmedMsg = trimmedMsg[2:]

	i := 0
	var openObj int
	var closeObj int
	for idx := strings.Index(trimmedMsg, "("); idx != -1 && idx < len(trimmedMsg); idx++ {
		if string(trimmedMsg[idx]) == "(" {
			if i == 0 {
				openObj = idx
			}
			i++
		}
		if string(trimmedMsg[idx]) == ")" {
			i--
			if i == 0 {
				closeObj = idx
			}
		}

		var obj string
		if i == 0 {
			obj = trimmedMsg[openObj+1 : closeObj]

			var key string
			var splitParams []string
			if strings.Contains(obj, ") (") {
				openIdx := strings.Index(obj, "(")

				key = obj[:openIdx-1]
				params := obj[len(key)+1:]

				splitParams = strings.Split(params[1:len(params)-1], ") (")
			} else {
				keyEnd := strings.Index(obj, " ")

				key = obj[:keyEnd]
				params := obj[keyEnd+1:]

				splitParams = strings.Split(params, " ")
			}
			data.SenseBodyMap[key] = splitParams
			idx++
		}

	}

	return
}
