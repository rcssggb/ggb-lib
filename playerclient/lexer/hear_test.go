package lexer

import (
	"testing"
)

func TestHearReferee(t *testing.T) {
	hearSyms, err := Hear("(hear 3000 referee half_time)\x00")

	if err != nil {
		t.Fail()
	}

	if hearSyms == nil {
		t.FailNow()
	}

	if hearSyms.Time != 3000 {
		t.Fail()
	}

	if hearSyms.Sender != "referee" {
		t.Fail()
	}

	if hearSyms.Message != "half_time" {
		t.Fail()
	}
}

func TestHearSelf(t *testing.T) {
	hearSyms, err := Hear("(hear 3921 self \"R4-l7f3_F2\")\x00")

	if err != nil {
		t.Fail()
	}

	if hearSyms == nil {
		t.FailNow()
	}

	if hearSyms.Time != 3921 {
		t.Fail()
	}

	if hearSyms.Sender != "self" {
		t.Fail()
	}

	if hearSyms.Message != "\"R4-l7f3_F2\"" {
		t.Fail()
	}
}

func TestHearSelfWithSpaceInMessage(t *testing.T) {
	hearSyms, err := Hear("(hear 4 self \"RrNx?*Rg U\")\x00")

	if err != nil {
		t.Fail()
	}

	if hearSyms == nil {
		t.FailNow()
	}

	if hearSyms.Time != 4 {
		t.Fail()
	}

	if hearSyms.Sender != "self" {
		t.Fail()
	}

	if hearSyms.Message != "\"RrNx?*Rg U\"" {
		t.Fail()
	}
}

//

// (hear 4 self "RrNx?*Rg U")

// (hear 1494 (p "HELIOS_A" 6) "Bnraishl+k")

// (hear 1495 82 our 6 "BMD3UHL gu")

// (hear 0 referee kick_off_l)
