package rcsscommon

// PlayModeID is the type representing each game mode
type PlayModeID byte

const (
	PlayModeBeforeKickOff PlayModeID = iota + 0 // before_kick_off
	PlayModePlayOn                              // play_on
	PlayModeTimeOver                            // time_over
	PlayModeKickOffR                            // kick_off_r
	PlayModeKickOffL                            // kick_off_l
	PlayModeKickInR                             // kick_in_r
	PlayModeKickInL                             // kick_in_l
	PlayModeFreeKickR                           // free_kick_r
	PlayModeFreeKickL                           // free_kick_l
	PlayModeCornerKickR                         // corner_kick_r
	PlayModeCornerKickL                         // corner_kick_l
	PlayModeGoalKickR                           // goal_kick_r
	PlayModeGoalKickL                           // goal_kick_l
	PlayModeGoalR                               // goal_r
	PlayModeGoalL                               // goal_l
	PlayModeDropBall                            // drop_ball
	PlayModeOffsideR                            // offside_r
	PlayModeOffsideL                            // offside_l
)

func (id PlayModeID) String() string {
	switch id {
	case PlayModeBeforeKickOff:
		return "before_kick_off"
	case PlayModePlayOn:
		return "play_on"
	case PlayModeTimeOver:
		return "time_over"
	case PlayModeKickOffR:
		return "kick_off_r"
	case PlayModeKickOffL:
		return "kick_off_l"
	case PlayModeKickInR:
		return "kick_in_r"
	case PlayModeKickInL:
		return "kick_in_l"
	case PlayModeFreeKickR:
		return "free_kick_r"
	case PlayModeFreeKickL:
		return "free_kick_l"
	case PlayModeCornerKickR:
		return "corner_kick_r"
	case PlayModeCornerKickL:
		return "corner_kick_l"
	case PlayModeGoalKickR:
		return "goal_kick_r"
	case PlayModeGoalKickL:
		return "goal_kick_l"
	case PlayModeGoalR:
		return "goal_r"
	case PlayModeGoalL:
		return "goal_l"
	case PlayModeDropBall:
		return "drop_ball"
	case PlayModeOffsideR:
		return "offside_r"
	case PlayModeOffsideL:
		return "offside_l"
	default:
		return ""
	}
}
