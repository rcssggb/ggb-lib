package parser

import (
	"testing"

	"github.com/rcssggb/ggb-lib/playerclient/lexer"
	"github.com/rcssggb/ggb-lib/rcsscommon"
)

func TestSenseBody(t *testing.T) {
	senseBodySymbols := lexer.SenseBodySymbols{
		Time: 0,
		SenseBodyMap: map[string][]string{
			"view_mode":   {"high", "normal"},
			"stamina":     {"8000", "1", "130600"},
			"speed":       {"0", "0"},
			"head_angle":  {"0"},
			"kick":        {"0"},
			"dash":        {"0"},
			"turn":        {"1"},
			"say":         {"0"},
			"turn_neck":   {"1"},
			"catch":       {"0"},
			"move":        {"0"},
			"change_view": {"0"},
			"arm":         {"movable 0", "expires 0", "target 0 0", "count 0"},
			"focus":       {"target none", "count 0"},
			"tackle":      {"expires 0", "count 0"},
			"collision":   {"none"},
			"foul":        {"charged 0", "card none"}},
	}
	senseBodyData, err := SenseBody(senseBodySymbols)
	if err != nil {
		t.Fail()
	}
	if senseBodyData == nil {
		t.FailNow()
	}

	expectedData := SenseBodyData{
		Time: 0,
		ViewMode: ViewModeData{
			Quality: rcsscommon.ViewQualityHigh,
			Width:   rcsscommon.ViewWidthNormal,
		},
		Stamina: StaminaData{
			Value:    8000,
			Effort:   1,
			Capacity: 130600,
		},
		Speed: SpeedData{
			Magnitude: 0,
			Direction: 0,
		},
		HeadAngle:     0,
		KickCount:     0,
		DashCount:     0,
		TurnCount:     1,
		TurnNeckCount: 1,
		CatchCount:    0,
		MoveCount:     0,
	}

	if *senseBodyData != expectedData {
		t.Fail()
	}
}
