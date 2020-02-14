package parser

import (
	"sort"
	"testing"

	"github.com/rcssggb/ggb-lib/playerclient/lexer"
)

func TestSightShort(t *testing.T) {
	sightSymbols := lexer.SightSymbols{
		Time: 37,
		ObjMap: map[string][]string{
			"f b r 10":          []string{"25", "-21", "-0", "-0.6"},
			"f b l 10":          []string{"22.2", "28"},
			"p \"HELIOS_B\" 10": []string{"7.4", "1", "-0.148", "-1.8", "167", "-153"},
			"l b":               []string{"16.6", "-79"},
			"f c b":             []string{"16.6", "-1", "-0", "-0.8"},
			"f b 0":             []string{"21.5", "2", "-0", "-0.6"},
		},
	}
	sightData, err := Sight(sightSymbols)
	if err != nil {
		t.Fail()
	}
	if sightData == nil {
		t.FailNow()
	}

	if sightData.Time != 37 {
		t.Fail()
	}

	if !sort.IsSorted(sightData.Flags) {
		t.Fail()
	}

	if sightData.Flags.Len() != 4 {
		t.Fail()
	}
}

func TestSightWithBall(t *testing.T) {
	sightSymbols := lexer.SightSymbols{
		Time: 3000,
		ObjMap: map[string][]string{
			"f c":      []string{"14.6", "42", "0.292", "0.9"},
			"f l t":    []string{"76.7", "36"},
			"f l b":    []string{"68", "-20"},
			"f g l b":  []string{"63.4", "3"},
			"g l":      []string{"64.1", "10"},
			"f g l t":  []string{"65.4", "16"},
			"f p l b":  []string{"47.9", "-12"},
			"f p l c":  []string{"47.9", "13"},
			"f p l t":  []string{"55.7", "33"},
			"f t l 20": []string{"58", "58"},
			"f t l 30": []string{"63.4", "51"},
			"f t l 40": []string{"70.8", "45"},
			"f t l 50": []string{"78.3", "40"},
			"f b l 10": []string{"36.2", "-53"},
			"f b l 20": []string{"42.5", "-42"},
			"f b l 30": []string{"50.4", "-35"},
			"f b l 40": []string{"59.1", "-29"},
			"f b l 50": []string{"68", "-25"},
			"f l 0":    []string{"69.4", "9"},
			"f l t 10": []string{"71.5", "17"},
			"f l t 20": []string{"74.4", "24"},
			"f l t 30": []string{"79", "31"},
			"f l b 10": []string{"68.7", "1"},
			"f l b 20": []string{"69.4", "-8"},
			"f l b 30": []string{"71.5", "-16"},
			"b":        []string{"14.9", "42", "0.298", "0.9"},
			"l l":      []string{"63.4", "-89"},
		},
	}

	sightData, err := Sight(sightSymbols)
	if err != nil {
		t.Fail()
	}
	if sightData == nil {
		t.FailNow()
	}

	if sightData.Time != 3000 {
		t.Fail()
	}

	if !sort.IsSorted(sightData.Flags) {
		t.Fail()
	}

	if sightData.Flags.Len() != 25 {
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
}
