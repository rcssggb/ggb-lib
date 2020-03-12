package rcsscommon

import (
	"testing"
)

func TestPlayerType(t *testing.T) {
	errCh := make(chan string, 32)
	playerTypeData := ParsePlayerType("(player_type (id 13)(player_speed_max 1.05)(stamina_inc_max 50.7145)(player_decay 0.403377)(inertia_moment 5.08442)(dash_power_rate 0.00504759)(player_size 0.3)(kickable_margin 0.635063)(kick_rand 0.0350625)(extra_stamina 59.3244)(effort_max 0.962702)(effort_min 0.562702)(kick_power_rate 0.027)(foul_detect_probability 0.5)(catchable_area_l_stretch 1.03477))\x00", errCh)

	if len(errCh) != 0 {
		t.Fail()
	}

	if playerTypeData == (PlayerType{}) {
		t.FailNow()
	}

	expectedData := PlayerType{
		ID:                    13,
		PlayerSpeedMax:        1.05,
		StaminaIncMax:         50.7145,
		PlayerDecay:           0.403377,
		InertiaMoment:         5.08442,
		DashPowerRate:         0.00504759,
		PlayerSize:            0.3,
		KickableMargin:        0.635063,
		KickRand:              0.0350625,
		ExtraStamina:          59.3244,
		EffortMax:             0.962702,
		EffortMin:             0.562702,
		KickPowerRate:         0.027,
		FoulDetectProbability: 0.5,
		CatchableAreaLStretch: 1.03477,
	}

	if playerTypeData != expectedData {
		t.Fail()
	}
}

// (player_type (id 13)(player_speed_max 1.05)(stamina_inc_max 50.7145)(player_decay 0.403377)(inertia_moment 5.08442)(dash_power_rate 0.00504759)(player_size 0.3)(kickable_margin 0.635063)(kick_rand 0.0350625)(extra_stamina 59.3244)(effort_max 0.962702)(effort_min 0.562702)(kick_power_rate 0.027)(foul_detect_probability 0.5)(catchable_area_l_stretch 1.03477))

// (player_type (id 0)(player_speed_max 1.05)(stamina_inc_max 45)(player_decay 0.4)(inertia_moment 5)(dash_power_rate 0.006)(player_size 0.3)(kickable_margin 0.7)(kick_rand 0.1)(extra_stamina 50)(effort_max 1)(effort_min 0.6)(kick_power_rate 0.027)(foul_detect_probability 0.5)(catchable_area_l_stretch 1))
