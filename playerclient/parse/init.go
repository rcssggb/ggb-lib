package parse

import (
	"fmt"
)

// InitData is the format return by parse.Init
type InitData struct {
	Side     rune
	Unum     int
	PlayMode string
}

// Init parses (init Side Unum PlayMode)
func Init(m string) (data *InitData, err error) {
	var side rune
	var unum int
	var playMode string

	_, err = fmt.Sscanf(m, "(init %c %d %s", &side, &unum, &playMode)
	if err != nil {
		return
	}
	playMode = playMode[0 : len(playMode)-1] // trim out last )

	data = &InitData{
		Side:     side,
		Unum:     unum,
		PlayMode: playMode,
	}

	return
}
