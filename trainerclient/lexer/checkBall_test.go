package lexer

import (
	"testing"
)

func TestCheckBallInField(t *testing.T) {
	time, ballInfo, err := CheckBall("(ok check_ball 85 in_field)\x00")
	if err != nil {
		t.Fail()
	}

	if time != 85 {
		t.Fail()
	}

	if ballInfo != "in_field" {
		t.Fail()
	}
}

// (ok check_ball 85 in_field)
// (ok check_ball 32 goal_r)
// (ok check_ball 26 out_of_field)
