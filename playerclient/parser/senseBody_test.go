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
			"view_mode":   []string{"high", "normal"},
			"stamina":     []string{"8000", "1", "130600"},
			"speed":       []string{"0", "0"},
			"head_angle":  []string{"0"},
			"kick":        []string{"0"},
			"dash":        []string{"0"},
			"turn":        []string{"1"},
			"say":         []string{"0"},
			"turn_neck":   []string{"1"},
			"catch":       []string{"0"},
			"move":        []string{"0"},
			"change_view": []string{"0"},
			"arm":         []string{"movable 0", "expires 0", "target 0 0", "count 0"},
			"focus":       []string{"target none", "count 0"},
			"tackle":      []string{"expires 0", "count 0"},
			"collision":   []string{"none"},
			"foul":        []string{"charged 0", "card none"}},
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
