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
		"foul":        []string{"charged 0", "card none"},
	}

	if !reflect.DeepEqual(correctData, senseBodyData.SenseBodyMap) {
		t.Fail()
	}

}

// (sense_body 0 (view_mode high normal) (stamina 8000 1 130600) (speed 0 0) (head_angle 0) (kick 0) (dash 0) (turn 1) (say 0) (turn_neck 1) (catch 0) (move 0) (change_view 0) (arm (movable 0) (expires 0) (target 0 0) (count 0)) (focus (target none) (count 0)) (tackle (expires 0) (count 0)) (collision none) (foul  (charged 0) (card none)))
