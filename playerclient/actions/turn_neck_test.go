package actions

import (
	"testing"
)

func TestTurnNeck(t *testing.T) {
	turnNeckStr := TurnNeck(20)
	if turnNeckStr != "(turn_neck 20.000)" {
		t.Fail()
	}
}
