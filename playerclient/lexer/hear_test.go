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

func TestHearKnownPlayer(t *testing.T) {
	hearSyms, err := Hear("(hear 1494 (p \"HELIOS_A\" 6) \"Bnraishl+k\")\x00")

	if err != nil {
		t.Fail()
	}

	if hearSyms == nil {
		t.FailNow()
	}

	if hearSyms.Time != 1494 {
		t.Fail()
	}

	if hearSyms.Sender != "p \"HELIOS_A\" 6" {
		t.Fail()
	}

	if hearSyms.Message != "\"Bnraishl+k\"" {
		t.Fail()
	}
}

func TestHearKnownPlayerWithSpace(t *testing.T) {
	hearSyms, err := Hear("(hear 1494 (p \"HELIOS_A\" 6) \"Bnrais l+k\")\x00")

	if err != nil {
		t.Fail()
	}

	if hearSyms == nil {
		t.FailNow()
	}

	if hearSyms.Time != 1494 {
		t.Fail()
	}

	if hearSyms.Sender != "p \"HELIOS_A\" 6" {
		t.Fail()
	}

	if hearSyms.Message != "\"Bnrais l+k\"" {
		t.Fail()
	}
}

func TestHearOnlineCoach(t *testing.T) {
	hearSyms, err := Hear("(hear 5749 online_coach_left (freeform \"(player_types (1 0)(2 5)(3 12)(4 15)(5 10)(6 1)(7 2)(8 6)(9 9)(10 3)(11 16))\"))\x00")

	if err != nil {
		t.Fail()
	}

	if hearSyms == nil {
		t.FailNow()
	}

	if hearSyms.Time != 5749 {
		t.Fail()
	}

	if hearSyms.Sender != "online_coach_left" {
		t.Fail()
	}

	if hearSyms.Message != "(freeform \"(player_types (1 0)(2 5)(3 12)(4 15)(5 10)(6 1)(7 2)(8 6)(9 9)(10 3)(11 16))\")" {
		t.Fail()
	}
}

func TestHearDirection(t *testing.T) {
	hearSyms, err := Hear("(hear 1495 82 our 6 \"BMD3UHL gu\")\x00")

	if err != nil {
		t.Fail()
	}

	if hearSyms == nil {
		t.FailNow()
	}

	if hearSyms.Time != 1495 {
		t.Fail()
	}

	if hearSyms.Sender != "82 our 6" {
		t.Fail()
	}

	if hearSyms.Message != "\"BMD3UHL gu\"" {
		t.Fail()
	}
}
