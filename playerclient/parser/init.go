package parser

import (
	"fmt"
	"strings"

	"github.com/rcssggb/ggb-lib/rcsscommon"
)

// InitData is the format return by parse.Init
type InitData struct {
	Side     rcsscommon.SideType
	Unum     int
	PlayMode string
}

// Init parses (init Side Unum PlayMode)
func Init(m string) (data *InitData, err error) {
	var side rcsscommon.SideType
	var unum int
	var playMode string

	_, err = fmt.Sscanf(m, "(init %c %d %s", &side, &unum, &playMode)
	if err != nil {
		return
	}
	playModeEnd := strings.Index(playMode, ")")
	playMode = playMode[0:playModeEnd] // trim out last )

	data = &InitData{
		Side:     side,
		Unum:     unum,
		PlayMode: playMode,
	}

	return
}
