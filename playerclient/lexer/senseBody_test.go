package lexer

import (
	"reflect"
	"testing"
)

func TestSenseBody(t *testing.T) {
	senseBodyData, err := SenseBody("(sense_body 0 (view_mode high normal) (stamina 8000 1 130600) (speed 0 0) (head_angle 0) (kick 0) (dash 0) (turn 1) (say 0) (turn_neck 1) (catch 0) (move 0) (change_view 0) (arm (movable 0) (expires 0) (target 0 0) (count 0)) (focus (target none) (count 0)) (tackle (expires 0) (count 0)) (collision none) (foul (charged 0) (card none)))")
	if err != nil {
		t.Fail()
	}
	if senseBodyData == nil {
		t.FailNow()
	}
	if senseBodyData.Time != 0 {
		t.Fail()
	}

	correctData := map[string][]string{
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
		"foul":        {"charged 0", "card none"},
	}

	if !reflect.DeepEqual(correctData, senseBodyData.SenseBodyMap) {
		t.Fail()
	}
}
