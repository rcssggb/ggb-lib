package rcsscommon

// FlagID is the type representing each marker flag in the field
type FlagID byte

// TODO: add comment with each flag coordinate
const (
	FlagCenter FlagID = iota + 0

	FlagCenterTop

	FlagCenterBot

	FlagLeftGoal

	FlagLeftGoalTop

	FlagLeftGoalBot

	FlagRightGoal

	FlagRightGoalTop

	FlagRightGoalBot

	FlagLeftPenaltyCenter

	FlagLeftPenaltyTop

	FlagLeftPenaltyBot

	FlagRightPenaltyCenter

	FlagRightPenaltyTop

	FlagRightPenaltyBot

	FlagLeftTop

	FlagLeftBot

	FlagRightTop

	FlagRightBot

	FlagTopLeft50

	FlagTopLeft40

	FlagTopLeft30

	FlagTopLeft20

	FlagTopLeft10

	FlagTop0

	FlagTopRight10

	FlagTopRight20

	FlagTopRight30

	FlagTopRight40

	FlagTopRight50

	FlagBotLeft50

	FlagBotLeft40

	FlagBotLeft30

	FlagBotLeft20

	FlagBotLeft10

	FlagBot0

	FlagBotRight10

	FlagBotRight20

	FlagBotRight30

	FlagBotRight40

	FlagBotRight50

	FlagLeftTop30

	FlagLeftTop20

	FlagLeftTop10

	FlagLeft0

	FlagLeftBot10

	FlagLeftBot20

	FlagLeftBot30

	FlagRightTop30

	FlagRightTop20

	FlagRightTop10

	FlagRight0

	FlagRightBot10

	FlagRightBot20

	FlagRightBot30
)
