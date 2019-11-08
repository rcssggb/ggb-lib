package lexer

import "testing"

func TestInitBasic(t *testing.T) {
	initData, err := Init("(init l 1 before_kick_off)")
	if err != nil {
		t.Fail()
	}
	if initData == nil {
		t.FailNow()
	}
	if initData.PlayMode != "before_kick_off" {
		t.Fail()
	}
	if initData.Unum != 1 {
		t.Fail()
	}
	if initData.Side != 'l' {
		t.Fail()
	}
}

func TestInitIncomplete(t *testing.T) {
	initData, err := Init("(init l ")
	if err == nil {
		t.Fail()
	}
	if initData != nil {
		t.Fail()
	}
}
