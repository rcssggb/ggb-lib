package rcsscommon

// FlagID is the type representing each marker flag in the field
type FlagID byte

const (
	// FlagCenter 00.00 00.00
	FlagCenter FlagID = iota + 0

	// FlagCenterTop 00.00 -34.00
	FlagCenterTop

	// FlagCenterBot 00.00 34.00
	FlagCenterBot

	// FlagLeftGoal -52.50 00.00
	FlagLeftGoal

	// FlagLeftGoalTop -52.50 -07.01
	FlagLeftGoalTop

	// FlagLeftGoalBot -52.50 07.01
	FlagLeftGoalBot

	// FlagRightGoal 52.50 00.00
	FlagRightGoal

	// FlagRightGoalTop 52.50 -07.01
	FlagRightGoalTop

	// FlagRightGoalBot 52.50 07.01
	FlagRightGoalBot

	// FlagLeftPenaltyCenter -36.00 00.00
	FlagLeftPenaltyCenter

	// FlagLeftPenaltyTop -36.00 -20.20
	FlagLeftPenaltyTop

	// FlagLeftPenaltyBot -36.00 20.20
	FlagLeftPenaltyBot

	// FlagRightPenaltyCenter 36.00 00.00
	FlagRightPenaltyCenter

	// FlagRightPenaltyTop 36.00 -20.20
	FlagRightPenaltyTop

	// FlagRightPenaltyBot 36.00 20.20
	FlagRightPenaltyBot

	// FlagLeftTop -52.50 -34.00
	FlagLeftTop

	// FlagLeftBot -52.50 34.00
	FlagLeftBot

	// FlagRightTop 52.50 -34.00
	FlagRightTop

	// FlagRightBot 52.50 34.00
	FlagRightBot

	// FlagTopLeft50 -50.00 -39.00
	FlagTopLeft50

	// FlagTopLeft40 -40.00 -39.00
	FlagTopLeft40

	// FlagTopLeft30 -30.00 -39.00
	FlagTopLeft30

	// FlagTopLeft20 -20.00 -39.00
	FlagTopLeft20

	// FlagTopLeft10 -10.00 -39.00
	FlagTopLeft10

	// FlagTop0 00.00 -39.00
	FlagTop0

	// FlagTopRight10 10.00 -39.00
	FlagTopRight10

	// FlagTopRight20 20.00 -39.00
	FlagTopRight20

	// FlagTopRight30 30.00 -39.00
	FlagTopRight30

	// FlagTopRight40 40.00 -39.00
	FlagTopRight40

	// FlagTopRight50 50.00 -39.00
	FlagTopRight50

	// FlagBotLeft50 -50.00 39.00
	FlagBotLeft50

	// FlagBotLeft40 -40.00 39.00
	FlagBotLeft40

	// FlagBotLeft30 -30.00 39.00
	FlagBotLeft30

	// FlagBotLeft20 -20.00 39.00
	FlagBotLeft20

	// FlagBotLeft10 -10.00 39.00
	FlagBotLeft10

	// FlagBot0 00.00 39.00
	FlagBot0

	// FlagBotRight10 10.00 39.00
	FlagBotRight10

	// FlagBotRight20 20.00 39.00
	FlagBotRight20

	// FlagBotRight30 30.00 39.00
	FlagBotRight30

	// FlagBotRight40 40.00 39.00
	FlagBotRight40

	// FlagBotRight50 50.00 39.00
	FlagBotRight50

	// FlagLeftTop30 -57.50 -30.00
	FlagLeftTop30

	// FlagLeftTop20 -57.50 -20.00
	FlagLeftTop20

	// FlagLeftTop10 -57.50 -10.00
	FlagLeftTop10

	// FlagLeft0 -57.50 00.00
	FlagLeft0

	// FlagLeftBot10 -57.50 10.00
	FlagLeftBot10

	// FlagLeftBot20 -57.50 20.00
	FlagLeftBot20

	// FlagLeftBot30 -57.50 30.00
	FlagLeftBot30

	// FlagRightTop30 57.50 -30.00
	FlagRightTop30

	// FlagRightTop20 57.50 -20.00
	FlagRightTop20

	// FlagRightTop10 57.50 -10.00
	FlagRightTop10

	// FlagRight0 57.50 00.00
	FlagRight0

	// FlagRightBot10 57.50 10.00
	FlagRightBot10

	// FlagRightBot20 57.50 20.00
	FlagRightBot20

	// FlagRightBot30 57.50 30.00
	FlagRightBot30
)

// Position returns the flag's x and y coordinates
func (f FlagID) Position() (x, y float64) {
	switch f {
	case FlagCenter:
		return 00.00, 00.00
	case FlagCenterTop:
		return 00.00, -34.00
	case FlagCenterBot:
		return 00.00, 34.00
	case FlagLeftGoal:
		return -52.50, 00.00
	case FlagLeftGoalTop:
		return -52.50, -07.01
	case FlagLeftGoalBot:
		return -52.50, 07.01
	case FlagRightGoal:
		return 52.50, 00.00
	case FlagRightGoalTop:
		return 52.50, -07.01
	case FlagRightGoalBot:
		return 52.50, 07.01
	case FlagLeftPenaltyCenter:
		return -36.00, 00.00
	case FlagLeftPenaltyTop:
		return -36.00, -20.20
	case FlagLeftPenaltyBot:
		return -36.00, 20.20
	case FlagRightPenaltyCenter:
		return 36.00, 00.00
	case FlagRightPenaltyTop:
		return 36.00, -20.20
	case FlagRightPenaltyBot:
		return 36.00, 20.20
	case FlagLeftTop:
		return -52.50, -34.00
	case FlagLeftBot:
		return -52.50, 34.00
	case FlagRightTop:
		return 52.50, -34.00
	case FlagRightBot:
		return 52.50, 34.00
	case FlagTopLeft50:
		return -50.00, -39.00
	case FlagTopLeft40:
		return -40.00, -39.00
	case FlagTopLeft30:
		return -30.00, -39.00
	case FlagTopLeft20:
		return -20.00, -39.00
	case FlagTopLeft10:
		return -10.00, -39.00
	case FlagTop0:
		return 00.00, -39.00
	case FlagTopRight10:
		return 10.00, -39.00
	case FlagTopRight20:
		return 20.00, -39.00
	case FlagTopRight30:
		return 30.00, -39.00
	case FlagTopRight40:
		return 40.00, -39.00
	case FlagTopRight50:
		return 50.00, -39.00
	case FlagBotLeft50:
		return -50.00, 39.00
	case FlagBotLeft40:
		return -40.00, 39.00
	case FlagBotLeft30:
		return -30.00, 39.00
	case FlagBotLeft20:
		return -20.00, 39.00
	case FlagBotLeft10:
		return -10.00, 39.00
	case FlagBot0:
		return 00.00, 39.00
	case FlagBotRight10:
		return 10.00, 39.00
	case FlagBotRight20:
		return 20.00, 39.00
	case FlagBotRight30:
		return 30.00, 39.00
	case FlagBotRight40:
		return 40.00, 39.00
	case FlagBotRight50:
		return 50.00, 39.00
	case FlagLeftTop30:
		return -57.50, -30.00
	case FlagLeftTop20:
		return -57.50, -20.00
	case FlagLeftTop10:
		return -57.50, -10.00
	case FlagLeft0:
		return -57.50, 00.00
	case FlagLeftBot10:
		return -57.50, 10.00
	case FlagLeftBot20:
		return -57.50, 20.00
	case FlagLeftBot30:
		return -57.50, 30.00
	case FlagRightTop30:
		return 57.50, -30.00
	case FlagRightTop20:
		return 57.50, -20.00
	case FlagRightTop10:
		return 57.50, -10.00
	case FlagRight0:
		return 57.50, 00.00
	case FlagRightBot10:
		return 57.50, 10.00
	case FlagRightBot20:
		return 57.50, 20.00
	case FlagRightBot30:
		return 57.50, 30.00
	}
	return 0, 0
}
