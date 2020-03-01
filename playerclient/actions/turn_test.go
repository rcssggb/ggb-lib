package actions

import (
	"testing"
)

func TestTurn(t *testing.T) {
	turnStr := Turn(20)
	if turnStr != "(turn 20.000)" {
		t.Fail()
	}
}
