package actions

import (
	"testing"
)

func TestDash(t *testing.T) {
	dashStr := Dash(27.431, 0)
	if dashStr != "(dash 27.431 0.000)" {
		t.Fail()
	}
}

func TestDashPositiveDirNormalization(t *testing.T) {
	dashStr := Dash(70, 270)
	if dashStr != "(dash 70.000 -90.000)" {
		t.Fail()
	}
}

func TestDashNegativeDirNormalization(t *testing.T) {
	dashStr := Dash(100, -225)
	if dashStr != "(dash 100.000 135.000)" {
		t.Fail()
	}
}
