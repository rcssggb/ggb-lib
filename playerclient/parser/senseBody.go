package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rcssggb/ggb-lib/playerclient/lexer"
	"github.com/rcssggb/ggb-lib/rcsscommon"
)

// ViewModeData specifies the settings of the sight sensor, which can be
// adjusted during the match (not currently supported)
type ViewModeData struct {
	// Defines ViewQualityFactor, which is a factor in defining visual sensor
	// period. Also affects stamina decay.
	// If set to ViewQualityLow, factor is 0.5.
	// If set to ViewQualityHigh, factor is 1.
	Quality rcsscommon.ViewQuality

	// Defines ViewWidthFactor which is a factor in defining visual cone angle.
	// If set to ViewWidthNarrow, factor is 0.5
	// If set to ViewWidthNormal, factor is 1
	// If set to ViewWidthWide, factor is 2
	Width rcsscommon.ViewWidth
}

// StaminaData specifies the current stamina levels
type StaminaData struct {
	// The actual stamina value for the player
	Value float64

	// The dash efficiency factor. Decreases when stamina is low.
	Effort float64

	// Max stamina (?)
	Capacity float64
}

// SpeedData specifies the current speed of the player
type SpeedData struct {
	// Magnitude of player's speed
	Magnitude float64

	// Direction of player's speed relative to face angle
	Direction float64
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

// TackleData specifies the number of cycles the current tackle will take
type TackleData struct {
	Expires int64
	Count   int64
}

// SenseBodyData is the final format of player information
type SenseBodyData struct {
	Time          int
	ViewMode      ViewModeData
	Stamina       StaminaData
	Speed         SpeedData
	HeadAngle     float64
	KickCount     uint64
	DashCount     uint64
	TurnCount     uint64
	TurnNeckCount uint64
	CatchCount    uint64
	MoveCount     uint64

	// Say        int64
	// ChangeView int64
	// Arm        ArmData
	// Focus      FocusData
	// Tackle     TackleData
	// Collision  string
	// Foul       FoulData
}

// SenseBody parses sense body data coming from lexer
func SenseBody(symbols lexer.SenseBodySymbols) (senseBodyData *SenseBodyData, err error) {
	senseBodyData = &SenseBodyData{
		Time: symbols.Time,
	}

	for fieldName, fieldData := range symbols.SenseBodyMap {
		switch fieldName {
		case "view_mode":
			if len(fieldData) != 2 {
				err = fmt.Errorf("invalid format: (%s %s)", fieldName, strings.Join(fieldData, " "))
				return
			}

			viewQuality := fieldData[0]
			switch viewQuality {
			case "high":
				(*senseBodyData).ViewMode.Quality = rcsscommon.ViewQualityHigh
			case "low":
				(*senseBodyData).ViewMode.Quality = rcsscommon.ViewQualityLow
			}

			viewWidth := fieldData[1]
			switch viewWidth {
			case "narrow":
				(*senseBodyData).ViewMode.Width = rcsscommon.ViewWidthNarrow
			case "normal":
				(*senseBodyData).ViewMode.Width = rcsscommon.ViewWidthNormal
			case "wide":
				(*senseBodyData).ViewMode.Width = rcsscommon.ViewWidthWide
			}
		case "stamina":
			if len(fieldData) != 3 {
				err = fmt.Errorf("invalid format: (%s %s)", fieldName, strings.Join(fieldData, " "))
				return
			}

			var val, eff, cap float64
			val, err = strconv.ParseFloat(fieldData[0], 64)
			if err != nil {
				err = fmt.Errorf("failed to parse body stamina value %s to float64: %s", fieldData[0], err)
				return
			}

			eff, err = strconv.ParseFloat(fieldData[1], 64)
			if err != nil {
				err = fmt.Errorf("failed to parse body stamina effort %s to float64: %s", fieldData[1], err)
				return
			}

			cap, err = strconv.ParseFloat(fieldData[2], 64)
			if err != nil {
				err = fmt.Errorf("failed to parse body stamina capacity %s to float64: %s", fieldData[2], err)
				return
			}

			(*senseBodyData).Stamina.Value = val
			(*senseBodyData).Stamina.Effort = eff
			(*senseBodyData).Stamina.Capacity = cap
		case "speed":
			if len(fieldData) != 2 {
				err = fmt.Errorf("invalid format: (%s %s)", fieldName, strings.Join(fieldData, " "))
				return
			}

			var mag, dir float64
			mag, err = strconv.ParseFloat(fieldData[0], 64)
			if err != nil {
				err = fmt.Errorf("failed to parse body speed magnitude %s to float64: %s", fieldData[0], err)
				return
			}

			dir, err = strconv.ParseFloat(fieldData[1], 64)
			if err != nil {
				err = fmt.Errorf("failed to parse body speed direction %s to float64: %s", fieldData[0], err)
				return
			}

			(*senseBodyData).Speed.Magnitude = mag
			(*senseBodyData).Speed.Direction = dir
		case "head_angle":
			if len(fieldData) != 1 {
				err = fmt.Errorf("invalid format: (%s %s)", fieldName, strings.Join(fieldData, " "))
				return
			}

			var headAngle float64
			headAngle, err = strconv.ParseFloat(fieldData[0], 64)
			if err != nil {
				err = fmt.Errorf("failed to parse body head angle %s to float64: %s", fieldData[0], err)
				return
			}

			(*senseBodyData).HeadAngle = headAngle
		case "kick":
			if len(fieldData) != 1 {
				err = fmt.Errorf("invalid format: (%s %s)", fieldName, strings.Join(fieldData, " "))
				return
			}

			var count uint64
			count, err = strconv.ParseUint(fieldData[0], 10, 64)
			if err != nil {
				err = fmt.Errorf("failed to parse body kick count %s to uint64: %s", fieldData[0], err)
				return
			}

			(*senseBodyData).KickCount = count
		case "dash":
			if len(fieldData) != 1 {
				err = fmt.Errorf("invalid format: (%s %s)", fieldName, strings.Join(fieldData, " "))
				return
			}

			var count uint64
			count, err = strconv.ParseUint(fieldData[0], 10, 64)
			if err != nil {
				err = fmt.Errorf("failed to parse body dash count %s to uint64: %s", fieldData[0], err)
				return
			}

			(*senseBodyData).DashCount = count
		case "turn":
			if len(fieldData) != 1 {
				err = fmt.Errorf("invalid format: (%s %s)", fieldName, strings.Join(fieldData, " "))
				return
			}

			var count uint64
			count, err = strconv.ParseUint(fieldData[0], 10, 64)
			if err != nil {
				err = fmt.Errorf("failed to parse body turn count %s to uint64: %s", fieldData[0], err)
				return
			}

			(*senseBodyData).TurnCount = count
		case "turn_neck":
			if len(fieldData) != 1 {
				err = fmt.Errorf("invalid format: (%s %s)", fieldName, strings.Join(fieldData, " "))
				return
			}

			var count uint64
			count, err = strconv.ParseUint(fieldData[0], 10, 64)
			if err != nil {
				err = fmt.Errorf("failed to parse body turn_neck count %s to uint64: %s", fieldData[0], err)
				return
			}

			(*senseBodyData).TurnNeckCount = count
		case "catch":
			if len(fieldData) != 1 {
				err = fmt.Errorf("invalid format: (%s %s)", fieldName, strings.Join(fieldData, " "))
				return
			}

			var count uint64
			count, err = strconv.ParseUint(fieldData[0], 10, 64)
			if err != nil {
				err = fmt.Errorf("failed to parse body catch count %s to uint64: %s", fieldData[0], err)
				return
			}

			(*senseBodyData).CatchCount = count
		case "move":
			if len(fieldData) != 1 {
				err = fmt.Errorf("invalid format: (%s %s)", fieldName, strings.Join(fieldData, " "))
				return
			}

			var count uint64
			count, err = strconv.ParseUint(fieldData[0], 10, 64)
			if err != nil {
				err = fmt.Errorf("failed to parse body move count %s to uint64: %s", fieldData[0], err)
				return
			}

			(*senseBodyData).MoveCount = count
		}
	}

	return
}
