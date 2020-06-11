package parser

import (
	"reflect"
	"testing"

	"github.com/rcssggb/ggb-lib/playerclient/lexer"
	"github.com/rcssggb/ggb-lib/rcsscommon"
)

func TestSightShort(t *testing.T) {
	sightSymbols := lexer.SightSymbols{
		Time: 37,
		ObjMap: map[string][]string{
			"f b r 10": {"25", "-21", "-0", "-0.6"},
			"f b l 10": {"22.2", "28"},
			"l b":      {"16.6", "-79"},
			"f c b":    {"16.6", "-1", "-0", "-0.8"},
			"f b 0":    {"21.5", "2", "-0", "-0.6"},
		},
		Players: lexer.SightPlayersSymbols{
			Known: map[string][]string{
				"p \"HELIOS_B\" 10": {"7.4", "1", "-0.148", "-1.8", "167", "-153"},
			},
			KnownTeam: map[string][][]string{},
			Unknown:   [][]string{},
		},
	}

	errCh := make(chan string, 32)
	sightData := Sight(sightSymbols, errCh)

	if len(errCh) != 0 {
		t.Fail()
	}

	if sightData == nil {
		t.FailNow()
	}

	if sightData.Time != 37 {
		t.Fail()
	}

	if sightData.Ball != nil {
		t.Fail()
	}

	expectedFlags := FlagArray{
		{
			ID:        rcsscommon.FlagCenterBot,
			Distance:  16.6,
			Direction: -1,
		},
		{
			ID:        rcsscommon.FlagBot0,
			Distance:  21.5,
			Direction: 2,
		},
		{
			ID:        rcsscommon.FlagBotLeft10,
			Distance:  22.2,
			Direction: 28,
		},
		{
			ID:        rcsscommon.FlagBotRight10,
			Distance:  25,
			Direction: -21,
		},
	}

	if !reflect.DeepEqual(sightData.Flags, expectedFlags) {
		t.Fail()
	}

	expectedPlayers := PlayerArray{
		PlayerData{
			Team:       "HELIOS_B",
			Unum:       10,
			Distance:   7.4,
			Direction:  1,
			DistChange: -0.148,
			DirChange:  -1.8,
			BodyDir:    167,
			HeadDir:    -153,
		},
	}

	if !reflect.DeepEqual(sightData.Players, expectedPlayers) {
		t.Fail()
	}
}

func TestSightWithBall(t *testing.T) {
	sightSymbols := lexer.SightSymbols{
		Time: 3000,
		ObjMap: map[string][]string{
			"f c":      {"14.6", "42", "0.292", "0.9"},
			"f l t":    {"76.7", "36"},
			"f l b":    {"68", "-20"},
			"f g l b":  {"63.4", "3"},
			"g l":      {"64.1", "10"},
			"f g l t":  {"65.4", "16"},
			"f p l b":  {"47.9", "-12"},
			"f p l c":  {"47.9", "13"},
			"f p l t":  {"55.7", "33"},
			"f t l 20": {"58", "58"},
			"f t l 30": {"63.4", "51"},
			"f t l 40": {"70.8", "45"},
			"f t l 50": {"78.3", "40"},
			"f b l 10": {"36.2", "-53"},
			"f b l 20": {"42.5", "-42"},
			"f b l 30": {"50.4", "-35"},
			"f b l 40": {"59.1", "-29"},
			"f b l 50": {"68", "-25"},
			"f l 0":    {"69.4", "9"},
			"f l t 10": {"71.5", "17"},
			"f l t 20": {"74.4", "24"},
			"f l t 30": {"79", "31"},
			"f l b 10": {"68.7", "1"},
			"f l b 20": {"69.4", "-8"},
			"f l b 30": {"71.5", "-16"},
			"b":        {"14.9", "42", "0.298", "0.9"},
			"l l":      {"63.4", "-89"},
		},
		Players: lexer.SightPlayersSymbols{
			Known: map[string][]string{
				"p \"HELIOS_B\" 6": {"20.1", "-3", "0.402", "0.5", "-35", "50"},
				"p \"HELIOS_B\" 3": {"27.1", "-12", "0", "0.2", "-52", "-122"},
				"p \"HELIOS_B\" 8": {"13.5", "-16", "0", "1", "-53", "-113"},
				"p \"HELIOS_A\" 7": {"10", "-34", "-0", "1.4", "-52", "38"},
			},
			KnownTeam: map[string][][]string{
				"\"HELIOS_B\"": {
					{"30", "13"},
					{"33.1", "44"},
					{"27.1", "-31"},
				},
			},
			Unknown: [][]string{
				{"60.3", "7"},
				{"60.3", "-6"},
				{"44.7", "6"},
				{"66.7", "17"},
			},
		},
	}

	errCh := make(chan string, 32)
	sightData := Sight(sightSymbols, errCh)

	if len(errCh) != 0 {
		t.Fail()
	}

	if sightData == nil {
		t.FailNow()
	}

	if sightData.Time != 3000 {
		t.Fail()
	}

	expectedFlags := FlagArray{
		{
			ID:        rcsscommon.FlagCenter,
			Distance:  14.6,
			Direction: 42,
		},
		{
			ID:        rcsscommon.FlagBotLeft10,
			Distance:  36.2,
			Direction: -53,
		},
		{
			ID:        rcsscommon.FlagBotLeft20,
			Distance:  42.5,
			Direction: -42,
		},
		{
			ID:        rcsscommon.FlagLeftPenaltyBot,
			Distance:  47.9,
			Direction: -12,
		},
		{
			ID:        rcsscommon.FlagLeftPenaltyCenter,
			Distance:  47.9,
			Direction: 13,
		},
		{
			ID:        rcsscommon.FlagBotLeft30,
			Distance:  50.4,
			Direction: -35,
		},
		{
			ID:        rcsscommon.FlagLeftPenaltyTop,
			Distance:  55.7,
			Direction: 33,
		},
		{
			ID:        rcsscommon.FlagTopLeft20,
			Distance:  58,
			Direction: 58,
		},
		{
			ID:        rcsscommon.FlagBotLeft40,
			Distance:  59.1,
			Direction: -29,
		},
		{
			ID:        rcsscommon.FlagLeftGoalBot,
			Distance:  63.4,
			Direction: 3,
		},
		{
			ID:        rcsscommon.FlagTopLeft30,
			Distance:  63.4,
			Direction: 51,
		},
		{
			ID:        rcsscommon.FlagLeftGoal,
			Distance:  64.1,
			Direction: 10,
		},
		{
			ID:        rcsscommon.FlagLeftGoalTop,
			Distance:  65.4,
			Direction: 16,
		},
		{
			ID:        rcsscommon.FlagLeftBot,
			Distance:  68,
			Direction: -20,
		},
		{
			ID:        rcsscommon.FlagBotLeft50,
			Distance:  68,
			Direction: -25,
		},
		{
			ID:        rcsscommon.FlagLeftBot10,
			Distance:  68.7,
			Direction: 1,
		},
		{
			ID:        rcsscommon.FlagLeftBot20,
			Distance:  69.4,
			Direction: -8,
		},
		{
			ID:        rcsscommon.FlagLeft0,
			Distance:  69.4,
			Direction: 9,
		},
		{
			ID:        rcsscommon.FlagTopLeft40,
			Distance:  70.8,
			Direction: 45,
		},
		{
			ID:        rcsscommon.FlagLeftBot30,
			Distance:  71.5,
			Direction: -16,
		},
		{
			ID:        rcsscommon.FlagLeftTop10,
			Distance:  71.5,
			Direction: 17,
		},
		{
			ID:        rcsscommon.FlagLeftTop20,
			Distance:  74.4,
			Direction: 24,
		},
		{
			ID:        rcsscommon.FlagLeftTop,
			Distance:  76.7,
			Direction: 36,
		},
		{
			ID:        rcsscommon.FlagTopLeft50,
			Distance:  78.3,
			Direction: 40,
		},
		{
			ID:        rcsscommon.FlagLeftTop30,
			Distance:  79,
			Direction: 31,
		},
	}

	if !reflect.DeepEqual(sightData.Flags, expectedFlags) {
		t.Fail()
	}

	expectedBall := BallData{
		Distance:   14.9,
		Direction:  42,
		DistChange: 0.298,
		DirChange:  0.9,
	}

	if sightData.Ball == nil || *sightData.Ball != expectedBall {
		t.Fail()
	}

	expectedPlayers := PlayerArray{
		{
			Team:       "HELIOS_A",
			Unum:       7,
			Distance:   10,
			Direction:  -34,
			DistChange: -0,
			DirChange:  1.4,
			BodyDir:    -52,
			HeadDir:    38,
		},
		{
			Team:       "HELIOS_B",
			Unum:       8,
			Distance:   13.5,
			Direction:  -16,
			DistChange: 0,
			DirChange:  1,
			BodyDir:    -53,
			HeadDir:    -113,
		},
		{
			Team:       "HELIOS_B",
			Unum:       6,
			Distance:   20.1,
			Direction:  -3,
			DistChange: 0.402,
			DirChange:  0.5,
			BodyDir:    -35,
			HeadDir:    50,
		},
		{
			Team:       "HELIOS_B",
			Unum:       3,
			Distance:   27.1,
			Direction:  -12,
			DistChange: 0,
			DirChange:  0.2,
			BodyDir:    -52,
			HeadDir:    -122,
		},
		{
			Team:      "HELIOS_B",
			Distance:  27.1,
			Direction: -31,
		},
		{
			Team:      "HELIOS_B",
			Distance:  30,
			Direction: 13,
		},
		{
			Team:      "HELIOS_B",
			Distance:  33.1,
			Direction: 44,
		},
		{
			Distance:  44.7,
			Direction: 6,
		},
		{
			Distance:  60.3,
			Direction: -6,
		},
		{
			Distance:  60.3,
			Direction: 7,
		},
		{
			Distance:  66.7,
			Direction: 17,
		},
	}

	if !reflect.DeepEqual(sightData.Players, expectedPlayers) {
		t.Fail()
	}
}
