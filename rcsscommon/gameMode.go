package rcsscommon

// ModeID is the type representing each game mode
type ModeID byte

const (
	BeforeKickOff ModeID = iota + 0 // before_kick_off
	PlayOn                          // play_on
	TimeOver                        // time_over
	KickOffR                        // kick_off_r
	KickOffL                        // kick_off_l
	KickInR                         // kick_in_r
	KickInL                         // kick_in_l
	FreeKickR                       // free_kick_r
	FreeKickL                       // free_kick_l
	CornerKickR                     // corner_kick_r
	CornerKickL                     // corner_kick_l
	GoalKickR                       // goal_kick_r
	GoalKickL                       // goal_kick_l
	GoalR                           // goal_r
	GoalL                           // goal_l
	DropBall                        // drop_ball
	OffsideR                        // offside_r
	OffsideL                        // offside_l
)
