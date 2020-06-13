package lexer

import (
	"reflect"
	"testing"
)

func TestTeamNames2(t *testing.T) {
	tPos, err := TeamNames("(ok team_names (team l ggb-lib-test) (team r ggb-lib-test-2))\x00")
	if err != nil {
		t.Fail()
	}
	if tPos == nil {
		t.FailNow()
	}

	expectedTpos := TeamPositions{
		"left":  "ggb-lib-test",
		"right": "ggb-lib-test-2",
	}

	if !reflect.DeepEqual(expectedTpos, tPos) {
		t.Fail()
	}
}

func TestTeamNames1(t *testing.T) {
	tPos, err := TeamNames("(ok team_names (team l ggb-lib-test))\x00")
	if err != nil {
		t.Fail()
	}
	if tPos == nil {
		t.FailNow()
	}

	expectedTpos := TeamPositions{
		"left": "ggb-lib-test",
	}

	if !reflect.DeepEqual(expectedTpos, tPos) {
		t.Fail()
	}
}

func TestTeamNames0(t *testing.T) {
	tPos, err := TeamNames("(ok team_names)\x00")
	if err != nil {
		t.Fail()
	}
	if tPos == nil {
		t.FailNow()
	}

	expectedTpos := TeamPositions{}

	if !reflect.DeepEqual(expectedTpos, tPos) {
		t.Fail()
	}
}
