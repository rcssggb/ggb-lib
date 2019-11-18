package actions

import (
	"fmt"
	"testing"

	"github.com/rcssggb/ggb-lib/common/field"
)

func TestMoveInt(t *testing.T) {
	moveStr := Move(2, 2)
	if moveStr != "(move 2.0 2.0)" {
		t.Fail()
	}
}

func TestMoveNegInt(t *testing.T) {
	moveStr := Move(-17, -17)
	if moveStr != "(move -17.0 -17.0)" {
		t.Fail()
	}
}

func TestMovePrec(t *testing.T) {
	moveStr := Move(2.42, 2.42)
	if moveStr != "(move 2.4 2.4)" {
		t.Fail()
	}
}

func TestMoveNegFloat(t *testing.T) {
	moveStr := Move(-22.7, -13.4)
	if moveStr != "(move -22.7 -13.4)" {
		t.Fail()
	}
}

func TestMoveOutOfBounds(t *testing.T) {
	moveStr := Move(field.MinX-1, 0)
	if moveStr != fmt.Sprintf("(move %.1f 0.0)", field.MinX) {
		t.Fail()
	}

	moveStr = Move(field.MaxX+7.2, 0)
	if moveStr != fmt.Sprintf("(move %.1f 0.0)", field.MaxX) {
		t.Fail()
	}

	moveStr = Move(7.2, field.MinY-2.4)
	if moveStr != fmt.Sprintf("(move 7.2 %.1f)", field.MinY) {
		t.Fail()
	}

	moveStr = Move(7.2, field.MaxY+271.2)
	if moveStr != fmt.Sprintf("(move 7.2 %.1f)", field.MaxY) {
		t.Fail()
	}

	moveStr = Move(field.MinX-12, field.MaxY+271.2)
	if moveStr != fmt.Sprintf("(move %.1f %.1f)", field.MinX, field.MaxY) {
		t.Fail()
	}
}
