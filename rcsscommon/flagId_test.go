package rcsscommon

import "testing"

func TestFlagCenterPos(t *testing.T) {
	x, y := FlagCenter.Position()
	if x != 0 {
		t.Fail()
	}
	if y != 0 {
		t.Fail()
	}
}

func TestFlagTopLeft40Pos(t *testing.T) {
	x, y := FlagTopLeft40.Position()
	if x != -40 {
		t.Fail()
	}
	if y != -39 {
		t.Fail()
	}
}
