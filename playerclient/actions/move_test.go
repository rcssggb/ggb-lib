package actions

import (
	"fmt"
	"testing"

	"github.com/rcssggb/ggb-lib/rcsscommon"
)

func TestMoveInt(t *testing.T) {
	moveStr := Move(2, 2)
	if moveStr != "(move 2.0 2.0)\x00" {
		t.Fail()
	}
}

func TestMoveNegInt(t *testing.T) {
	moveStr := Move(-17, -17)
	if moveStr != "(move -17.0 -17.0)\x00" {
		t.Fail()
	}
}

func TestMovePrec(t *testing.T) {
	moveStr := Move(2.42, 2.42)
	if moveStr != "(move 2.4 2.4)\x00" {
		t.Fail()
	}
}

func TestMoveNegFloat(t *testing.T) {
	moveStr := Move(-22.7, -13.4)
	if moveStr != "(move -22.7 -13.4)\x00" {
		t.Fail()
	}
}

func TestMoveOutOfBounds(t *testing.T) {
	moveStr := Move(rcsscommon.FieldMinX-1, 0)
	if moveStr != fmt.Sprintf("(move %.1f 0.0)\x00", rcsscommon.FieldMinX) {
		t.Fail()
	}

	moveStr = Move(rcsscommon.FieldMaxX+7.2, 0)
	if moveStr != fmt.Sprintf("(move %.1f 0.0)\x00", rcsscommon.FieldMaxX) {
		t.Fail()
	}

	moveStr = Move(7.2, rcsscommon.FieldMinY-2.4)
	if moveStr != fmt.Sprintf("(move 7.2 %.1f)\x00", rcsscommon.FieldMinY) {
		t.Fail()
	}

	moveStr = Move(7.2, rcsscommon.FieldMaxY+271.2)
	if moveStr != fmt.Sprintf("(move 7.2 %.1f)\x00", rcsscommon.FieldMaxY) {
		t.Fail()
	}

	moveStr = Move(rcsscommon.FieldMinX-12, rcsscommon.FieldMaxY+271.2)
	if moveStr != fmt.Sprintf("(move %.1f %.1f)\x00", rcsscommon.FieldMinX, rcsscommon.FieldMaxY) {
		t.Fail()
	}
}
