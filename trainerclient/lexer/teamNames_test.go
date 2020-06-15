package lexer

import (
	"testing"
)

func TestTeamNames2(t *testing.T) {
	lTeam, rTeam, err := TeamNames("(ok team_names (team l ggb-lib-test) (team r ggb-lib-test-2))\x00")
	if err != nil {
		t.Fail()
	}

	if lTeam != "ggb-lib-test" {
		t.Fail()
	}

	if rTeam != "ggb-lib-test-2" {
		t.Fail()
	}
}

func TestTeamNames1(t *testing.T) {
	lTeam, rTeam, err := TeamNames("(ok team_names (team l ggb-lib-test))\x00")
	if err != nil {
		t.Fail()
	}

	if lTeam != "ggb-lib-test" {
		t.Fail()
	}

	if rTeam != "" {
		t.Fail()
	}
}

func TestTeamNames0(t *testing.T) {
	lTeam, rTeam, err := TeamNames("(ok team_names)\x00")
	if err != nil {
		t.Fail()
	}

	if lTeam != "" {
		t.Fail()
	}

	if rTeam != "" {
		t.Fail()
	}
}
