package parser

import (
	"testing"
)

func TestChangePlayerTypeTeammate(t *testing.T) {
	msg := "(change_player_type 2 1)\x00"

	cpt, err := ChangePlayerType(msg)
	if err != nil {
		t.Fail()
	}

	if cpt.Unum != 2 {
		t.Fail()
	}

	if cpt.PlayerType != 1 {
		t.Fail()
	}
}

func TestChangePlayerTypeOpponent(t *testing.T) {
	msg := "(change_player_type 4)\x00"

	cpt, err := ChangePlayerType(msg)
	if err != nil {
		t.Fail()
	}

	if cpt.Unum != 4 {
		t.Fail()
	}

	if cpt.PlayerType != -1 {
		t.Fail()
	}
}
