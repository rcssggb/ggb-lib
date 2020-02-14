package lexer

import (
	"reflect"
	"testing"
)

func TestSightShort(t *testing.T) {
	sightData, err := Sight("(see 37 ((f c b) 16.6 -1 -0 -0.8) ((f b 0) 21.5 2 -0 -0.6) ((f b r 10) 25 -21 -0 -0.6) ((f b l 10) 22.2 28) ((p \"HELIOS_B\" 10) 7.4 1 -0.148 -1.8 167 -153) ((l b) 16.6 -79))\x00")
	if err != nil {
		t.Fail()
	}
	if sightData == nil {
		t.FailNow()
	}

	if sightData.Time != 37 {
		t.Fail()
	}

	correctObjData := map[string][]string{
		"f b r 10": []string{"25", "-21", "-0", "-0.6"},
		"f b l 10": []string{"22.2", "28"},
		// "p \"HELIOS_B\" 10": []string{"7.4", "1", "-0.148", "-1.8", "167", "-153"},
		"l b":   []string{"16.6", "-79"},
		"f c b": []string{"16.6", "-1", "-0", "-0.8"},
		"f b 0": []string{"21.5", "2", "-0", "-0.6"},
	}

	if !reflect.DeepEqual(correctObjData, sightData.ObjMap) {
		t.Fail()
	}

	correctPlayersData := SightPlayersSymbols{
		KnownPlayersMap: map[string][]string{
			"p \"HELIOS_B\" 10": []string{"7.4", "1", "-0.148", "-1.8", "167", "-153"},
		},
		KnownTeamPlayersMap: map[string][][]string{},
		UnknownPlayers:      [][]string{},
	}

	if !reflect.DeepEqual(correctPlayersData, sightData.PlayersMap) {
		t.Fail()
	}
}

func TestSightWithBallAndPlayers(t *testing.T) {
	sightData, err := Sight("(see 3000 ((f c) 14.6 42 0.292 0.9) ((f l t) 76.7 36) ((f l b) 68 -20) ((f g l b) 63.4 3) ((g l) 64.1 10) ((f g l t) 65.4 16) ((f p l b) 47.9 -12) ((f p l c) 47.9 13) ((f p l t) 55.7 33) ((f t l 20) 58 58) ((f t l 30) 63.4 51) ((f t l 40) 70.8 45) ((f t l 50) 78.3 40) ((f b l 10) 36.2 -53) ((f b l 20) 42.5 -42) ((f b l 30) 50.4 -35) ((f b l 40) 59.1 -29) ((f b l 50) 68 -25) ((f l 0) 69.4 9) ((f l t 10) 71.5 17) ((f l t 20) 74.4 24) ((f l t 30) 79 31) ((f l b 10) 68.7 1) ((f l b 20) 69.4 -8) ((f l b 30) 71.5 -16) ((b) 14.9 42 0.298 0.9) ((p) 60.3 7) ((p \"HELIOS_B\") 30 13) ((p \"HELIOS_B\" 3) 27.1 -12 0 0.2 -52 -122) ((p \"HELIOS_B\") 33.1 44) ((p \"HELIOS_B\") 27.1 -31) ((p \"HELIOS_B\" 6) 20.1 -3 0.402 0.5 -35 50) ((p \"HELIOS_B\" 8) 13.5 -16 0 1 -53 -113) ((p) 60.3 -6) ((p) 44.7 6) ((p) 66.7 17) ((p \"HELIOS_A\" 7) 10 -34 -0 1.4 -52 38) ((l l) 63.4 -89))\x00")
	if err != nil {
		t.Fail()
	}
	if sightData == nil {
		t.FailNow()
	}

	if sightData.Time != 3000 {
		t.Fail()
	}

	correctObjMap := map[string][]string{
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
	}

	if !reflect.DeepEqual(correctObjMap, sightData.ObjMap) {
		t.Fail()
	}

	correctPlayersData := SightPlayersSymbols{
		KnownPlayersMap: map[string][]string{
			"p \"HELIOS_B\" 6": []string{"20.1", "-3", "0.402", "0.5", "-35", "50"},
			"p \"HELIOS_B\" 3": []string{"27.1", "-12", "0", "0.2", "-52", "-122"},
			"p \"HELIOS_B\" 8": []string{"13.5", "-16", "0", "1", "-53", "-113"},
			"p \"HELIOS_A\" 7": []string{"10", "-34", "-0", "1.4", "-52", "38"},
		},
		KnownTeamPlayersMap: map[string][][]string{
			"\"HELIOS_B\"": [][]string{
				[]string{"30", "13"},
				[]string{"33.1", "44"},
				[]string{"27.1", "-31"},
			},
		},
		UnknownPlayers: [][]string{
			[]string{"60.3", "7"},
			[]string{"60.3", "-6"},
			[]string{"44.7", "6"},
			[]string{"66.7", "17"},
		},
	}

	if !reflect.DeepEqual(correctPlayersData, sightData.PlayersMap) {
		t.Fail()
	}
}

// (see 0 ((f r t) 49.4 3) ((f r b) 86.5 55) ((f g r b) 66 42) ((g r) 61.6 37) ((f g r t) 58 31) ((f p r c) 49.4 48) ((f p r t) 37 27) ((f t r 10) 7.3 -16 0 0) ((f t r 20) 17.1 -7 0 0) ((f t r 30) 27.1 -4) ((f t r 40) 37 -3) ((f t r 50) 47 -2) ((f b r 50) 89.1 58) ((f r 0) 66 34) ((f r t 10) 60.9 26) ((f r t 20) 56.8 17) ((f r t 30) 55.1 7) ((f r b 10) 72.2 41) ((f r b 20) 79 46) ((f r b 30) 86.5 51))

// (see 2372 ((f c) 47.9 24) ((f c b) 61.6 -10) ((f l t) 104.6 46) ((f l b) 107.8 8) ((f g l b) 101.5 23) ((g l) 100.5 27) ((f g l t) 100.5 31) ((f p r b) 27.9 -35 -0 0.3) ((f p r c) 12.9 6 0 0.5) ((f p l b) 87.4 13) ((f p l c) 83.9 26) ((f p l t) 84.8 40) ((f t l 20) 75.9 56) ((f t l 30) 84.8 53) ((f t l 40) 93.7 51) ((f t l 50) 103.5 49) ((f b 0) 65.4 -13) ((f b r 10) 58 -20) ((f b r 20) 51.9 -28) ((f b r 30) 47.5 -38) ((f b r 40) 44.7 -50) ((f b l 10) 73 -8) ((f b l 20) 80.6 -3) ((f b l 30) 89.1 0) ((f b l 40) 98.5 3) ((f b l 50) 107.8 5) ((f l 0) 105.6 27) ((f l t 10) 105.6 32) ((f l t 20) 106.7 38) ((f l t 30) 107.8 43) ((f l b 10) 106.7 22) ((f l b 20) 107.8 16) ((f l b 30) 111.1 11) ((b) 24.5 42 1.96 -0.6) ((p) 99.5 27) ((p) 49.4 33) ((p "HELIOS_B") 49.4 16) ((p "HELIOS_B") 44.7 44) ((p "HELIOS_B") 49.4 5) ((p "HELIOS_B") 33.1 26) ((p "HELIOS_B" 7) 33.1 42 -0 0 -118 -180) ((p "HELIOS_B") 36.6 3) ((p "HELIOS_B") 33.1 -13) ((p "HELIOS_B" 11) 22.2 25 -0 0.8 134 93) ((p "HELIOS_A" 2) 18.2 -3 -0 0.5 68 40) ((p "HELIOS_A" 3) 18.2 44 0.364 -0.1 27 53) ((p "HELIOS_A") 30 -29) ((p "HELIOS_A" 6) 16.4 24 0.328 0.7 60 -29) ((p "HELIOS_A") 30 0) ((p "HELIOS_A" 8) 27.1 43 -0 -0.2 -81 -17) ((p) 44.7 -12) ((p "HELIOS_A" 11) 33.1 24 0 0.2 65 -24) ((l b) 79 30))

// (see 2372 ((f c) 38.1 -26) ((f c t) 62.8 2) ((f r t) 56.8 54) ((f l t) 100.5 -24) ((f g l b) 86.5 -48) ((g l) 87.4 -44) ((f g l t) 89.1 -39) ((f p r c) 19.5 43 -0 0.2) ((f p r t) 39.6 38) ((f p l b) 68.7 -57) ((f p l c) 71.5 -41) ((f p l t) 79.8 -26) ((f t 0) 66.7 4) ((f t r 10) 62.8 12) ((f t r 20) 59.7 21) ((f t r 30) 58.6 31) ((f t r 40) 58.6 41) ((f t r 50) 60.9 50) ((f t l 10) 72.2 -3) ((f t l 20) 79 -9) ((f t l 30) 85.6 -13) ((f t l 40) 93.7 -18) ((f t l 50) 101.5 -21) ((f l 0) 92.8 -44) ((f l t 10) 95.6 -38) ((f l t 20) 98.5 -33) ((f l t 30) 103.5 -28) ((f l b 10) 90.9 -50) ((f l b 20) 90.9 -57) ((b) 30 16) ((p) 81.5 -43) ((p "HELIOS_B") 44.7 -17) ((p "HELIOS_B") 36.6 -34) ((p) 44.7 -5) ((p "HELIOS_B" 5) 30 -50 -0 -0.3 73 112) ((p "HELIOS_B") 30 -6) ((p "HELIOS_B") 36.6 6) ((p "HELIOS_B" 8) 18.2 -29 -0 0.2 52 3) ((p "HELIOS_B") 40.4 22) ((p "HELIOS_B" 10) 10 -47 -0 -0.8 41 -49) ((p "HELIOS_B") 24.5 16) ((p "HELIOS_A" 2) 14.9 34 -0 0.3 -18 -46) ((p "HELIOS_A" 3) 30 27 -0 -0.6 -59 -33) ((p "HELIOS_A") 40.4 33) ((p "HELIOS_A" 6) 22.2 29 -0 -0.2 -26 -115) ((p "HELIOS_A" 7) 14.9 -16 -0.298 -0.3 -33 -89) ((p "HELIOS_A" 8) 33.1 12 -0.662 0 -167 -103) ((p) 49.4 10) ((p "HELIOS_A") 27.1 -8) ((l t) 64.1 -56))

// (see 4950 ((f c) 20.7 -73 0 -0.1) ((f c b) 39.6 -14) ((f r b) 80.6 -47) ((f l b) 46.5 61) ((f g r b) 73.7 -67) ((g r) 73 -72) ((f g r t) 73.7 -78) ((f p r b) 60.3 -53) ((f p r c) 56.8 -72) ((f p l b) 25 55 -0 -0.1) ((f b 0) 43.8 -10) ((f b r 10) 49.4 -20) ((f b r 20) 56.3 -28) ((f b r 30) 64.1 -35) ((f b r 40) 72.2 -39) ((f b r 50) 80.6 -43) ((f b l 10) 40 3) ((f b l 20) 38.9 17) ((f b l 30) 40 31) ((f b l 40) 43.4 44) ((f b l 50) 48.4 55) ((f r 0) 78.3 -72) ((f r t 10) 79 -79) ((f r t 20) 80.6 -87) ((f r b 10) 79 -65) ((f r b 20) 80.6 -58) ((f r b 30) 83.9 -51) ((f l b 20) 41.7 80) ((f l b 30) 47.5 69) ((p "HELIOS_B" 5) 11 24 -0 1 107 127) ((p "HELIOS_B" 6) 27.1 -83 0 -0.1 121 31) ((p "HELIOS_B" 8) 18.2 -46 -0 0.2 105 106) ((p "HELIOS_B") 40.4 -43) ((p "HELIOS_B") 33.1 -86) ((p) 73.7 -74) ((p "HELIOS_A" 2) 27.1 -69 -0 0.1 126 36) ((p "HELIOS_A" 4) 27.1 -41 -0 0.1 149 60) ((p "HELIOS_A") 36.6 -85) ((p "HELIOS_A" 7) 18.2 -57 -0 0 87 21) ((p "HELIOS_A" 9) 22.2 7 0 0.3 68 158) ((l b) 35.5 -72))

// (see 3001 ((f c) 18 41 -0 0) ((f c t) 46.5 78) ((f c b) 28.2 -51) ((f l t) 80.6 40) ((f l b) 71.5 -12) ((f g l b) 67.4 10) ((g l) 68 16) ((f g l t) 69.4 21) ((f p l b) 51.9 -4) ((f p l c) 51.9 18) ((f p l t) 59.1 38) ((f t 0) 51.4 80) ((f t l 10) 55.1 70) ((f t l 20) 60.3 62) ((f t l 30) 66.7 55) ((f t l 40) 73.7 49) ((f t l 50) 81.5 44) ((f b 0) 32.8 -55) ((f b r 10) 29.4 -73) ((f b l 10) 38.5 -42) ((f b l 20) 45.6 -32) ((f b l 30) 53.5 -26) ((f b l 40) 62.2 -21) ((f b l 50) 71.5 -17) ((f l 0) 73 15) ((f l t 10) 75.2 23) ((f l t 20) 78.3 30) ((f l t 30) 83.1 36) ((f l b 10) 72.2 7) ((f l b 20) 73 -1) ((f l b 30) 75.2 -8) ((b) 18.2 41 -0 0) ((p) 66.7 16) ((p) 44.7 28) ((p "HELIOS_B") 40.4 14) ((p "HELIOS_B") 44.7 34) ((p "HELIOS_B") 40.4 7) ((p "HELIOS_B") 40.4 21) ((p "HELIOS_B") 33.1 34) ((p "HELIOS_B" 8) 30 17 -0 0 -154 -154) ((p "HELIOS_B") 36.6 41) ((p "HELIOS_B" 10) 30 7 0 0 -33 -33) ((p "HELIOS_B" 11) 33.1 26 -0 0 -149 -149) ((l l) 68 -83))
