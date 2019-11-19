package parser

import (
	"fmt"

	"github.com/rcssggb/ggb-lib/playerclient/lexer"
	body "github.com/rcssggb/ggb-lib/rcsscommon/senseBody"
)

// ViewModeData specifies the settings of the sight sensor
type ViewModeData struct {
	Quality body.Quality
	Width   body.Width
}

// StaminaData specifies the current stamina levels
type StaminaData struct {
	Value           float64
	Effort          float64
	StaminaCapacity float64
}

// SpeedData specifies the current speed of the player
type SpeedData struct {
	Magnitude float64
	Direction float64 // this is relative to face angle
}

// TargetData specifies distance and direction the player is pointing to
type TargetData struct {
	Distance  float64
	Direction float64
}

// ArmData specifies if arm is movable, when the pointing will expire and the target
type ArmData struct {
	Movable int64
	Expires int64
	Target  TargetData
	Count   int64
}

// FocusData specifies where the player is focusing
type FocusData struct {
	Side  body.SideID
	Count int64
}

// TackleData specifies the number of cycles the current tackle will take
type TackleData struct {
	Expires int64
	Count   int64
}

// FoulData specifies the foul expire cycles and card
type FoulData struct {
	Charged int64
	Card    body.CardID
}

// SenseBodyData is the final format of player information
type SenseBodyData struct {
	Time       int
	ViewMode   ViewModeData
	Stamina    StaminaData
	Speed      SpeedData
	HeadAngle  float64
	Kick       int64
	Dash       int64
	Turn       int64
	Say        int64
	TurnNeck   int64
	Catch      int64
	Move       int64
	ChangeView int64
	Arm        ArmData
	Focus      FocusData
	Tackle     TackleData
	Collision  string
	Foul       FoulData
}

// SenseBody parses sense body data coming from lexer
func SenseBody(symbols lexer.SenseBodySymbols) (senseBodyData *SenseBodyData, err error) {
	senseBodyData = &SenseBodyData{
		Time: symbols.Time,
	}

	for objName, _ := range symbols.SenseBodyMap {
		switch objName {

		default:
			err = fmt.Errorf("Parse not implemented for %s", objName)
			// fmt.Println(err.Error())
		}
	}

	return
}
