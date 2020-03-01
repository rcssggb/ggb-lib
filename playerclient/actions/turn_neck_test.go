package actions

import (
	"testing"
)

func TestTurnNeck(t *testing.T) {
	kickStr := TurnNeck(20)
	if kickStr != "(turn_neck 20.000)" {
		t.Fail()
	}
}
