package trainerclient

import "github.com/rcssggb/ggb-lib/rcsscommon"

/*
ChangeMode changes the game mode
*/
func (c *Client) ChangeMode(mode rcsscommon.ModeID) string {
	changeModeString := "(change_mode "
	switch mode {
	case rcsscommon.BeforeKickOff:
		changeModeString = changeModeString + "before_kick_off)"
	case rcsscommon.PlayOn:
		changeModeString = changeModeString + "play_on)"
	case rcsscommon.KickOffR:
		changeModeString = changeModeString + "kick_off_r)"
	case rcsscommon.KickOffL:
		changeModeString = changeModeString + "kick_off_l)"
	case rcsscommon.KickInR:
		changeModeString = changeModeString + "kick_in_r)"
	case rcsscommon.KickInL:
		changeModeString = changeModeString + "kick_in_l)"
	case rcsscommon.FreeKickR:
		changeModeString = changeModeString + "free_kick_r)"
	case rcsscommon.FreeKickL:
		changeModeString = changeModeString + "free_kick_l)"
	case rcsscommon.CornerKickR:
		changeModeString = changeModeString + "corner_kick_r)"
	case rcsscommon.CornerKickL:
		changeModeString = changeModeString + "corner_kick_l)"
	case rcsscommon.GoalKickR:
		changeModeString = changeModeString + "goal_kick_r)"
	case rcsscommon.GoalKickL:
		changeModeString = changeModeString + "goal_kick_l)"
	case rcsscommon.GoalR:
		changeModeString = changeModeString + "goal_r)"
	case rcsscommon.GoalL:
		changeModeString = changeModeString + "goal_l)"
	case rcsscommon.DropBall:
		changeModeString = changeModeString + "drop_ball)"
	case rcsscommon.OffsideR:
		changeModeString = changeModeString + "offside_r)"
	case rcsscommon.OffsideL:
		changeModeString = changeModeString + "offside_l)"

	}
	c.cmdChannel <- command{
		cmdString: changeModeString,
	}
	return changeModeString
}
