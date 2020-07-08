package lexer

import (
	"testing"
)

func TestChangePlayerType3(t *testing.T) {
	teamName, unum, ID, err := ChangePlayerType("(ok change_player_type ggb-team 8 3)\x00")
	if err != nil {
		t.Fail()
	}

	if teamName != "ggb-team" {
		t.Fail()
	}

	if unum != 8 {
		t.Fail()
	}
	if ID != 3 {
		t.Fail()
	}

}
