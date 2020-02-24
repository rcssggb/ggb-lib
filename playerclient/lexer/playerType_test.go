package lexer

import (
	"reflect"
	"testing"
)

func TestPlayerType(t *testing.T) {
	playerTypeData, err := PlayerType("(player_type (id 13)(player_speed_max 1.05)(stamina_inc_max 50.7145)(player_decay 0.403377)(inertia_moment 5.08442)(dash_power_rate 0.00504759)(player_size 0.3)(kickable_margin 0.635063)(kick_rand 0.0350625)(extra_stamina 59.3244)(effort_max 0.962702)(effort_min 0.562702)(kick_power_rate 0.027)(foul_detect_probability 0.5)(catchable_area_l_stretch 1.03477))\x00")
	if err != nil {
		t.Fail()
	}
	if playerTypeData == nil {
		t.FailNow()
	}

	correctData := map[string]string{
		"id":                       "13",
		"player_speed_max":         "1.05",
		"stamina_inc_max":          "50.7145",
		"player_decay":             "0.403377",
		"inertia_moment":           "5.08442",
		"dash_power_rate":          "0.00504759",
		"player_size":              "0.3",
		"kickable_margin":          "0.635063",
		"kick_rand":                "0.0350625",
		"extra_stamina":            "59.3244",
		"effort_max":               "0.962702",
		"effort_min":               "0.562702",
		"kick_power_rate":          "0.027",
		"foul_detect_probability":  "0.5",
		"catchable_area_l_stretch": "1.03477",
	}

	if !reflect.DeepEqual(correctData, playerTypeData) {
		t.Fail()
	}

}

// (player_type (id 13)(player_speed_max 1.05)(stamina_inc_max 50.7145)(player_decay 0.403377)(inertia_moment 5.08442)(dash_power_rate 0.00504759)(player_size 0.3)(kickable_margin 0.635063)(kick_rand 0.0350625)(extra_stamina 59.3244)(effort_max 0.962702)(effort_min 0.562702)(kick_power_rate 0.027)(foul_detect_probability 0.5)(catchable_area_l_stretch 1.03477))
