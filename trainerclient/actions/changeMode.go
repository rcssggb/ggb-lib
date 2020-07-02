package actions

import "github.com/rcssggb/ggb-lib/rcsscommon"

// ChangeMode constructs the string to send a change_mode command to the server
func ChangeMode(mode rcsscommon.PlayModeID) string {
	changeModeString := "(change_mode "

	switch mode {
	case rcsscommon.PlayModeBeforeKickOff:
		changeModeString = changeModeString + "before_kick_off)"
	case rcsscommon.PlayModePlayOn:
		changeModeString = changeModeString + "play_on)"
	case rcsscommon.PlayModeKickOffR:
		changeModeString = changeModeString + "kick_off_r)"
	case rcsscommon.PlayModeKickOffL:
		changeModeString = changeModeString + "kick_off_l)"
	case rcsscommon.PlayModeKickInR:
		changeModeString = changeModeString + "kick_in_r)"
	case rcsscommon.PlayModeKickInL:
		changeModeString = changeModeString + "kick_in_l)"
	case rcsscommon.PlayModeFreeKickR:
		changeModeString = changeModeString + "free_kick_r)"
	case rcsscommon.PlayModeFreeKickL:
		changeModeString = changeModeString + "free_kick_l)"
	case rcsscommon.PlayModeCornerKickR:
		changeModeString = changeModeString + "corner_kick_r)"
	case rcsscommon.PlayModeCornerKickL:
		changeModeString = changeModeString + "corner_kick_l)"
	case rcsscommon.PlayModeGoalKickR:
		changeModeString = changeModeString + "goal_kick_r)"
	case rcsscommon.PlayModeGoalKickL:
		changeModeString = changeModeString + "goal_kick_l)"
	case rcsscommon.PlayModeGoalR:
		changeModeString = changeModeString + "goal_r)"
	case rcsscommon.PlayModeGoalL:
		changeModeString = changeModeString + "goal_l)"
	case rcsscommon.PlayModeDropBall:
		changeModeString = changeModeString + "drop_ball)"
	case rcsscommon.PlayModeOffsideR:
		changeModeString = changeModeString + "offside_r)"
	case rcsscommon.PlayModeOffsideL:
		changeModeString = changeModeString + "offside_l)"
	}

	return changeModeString
}
