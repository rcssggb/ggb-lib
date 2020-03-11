package rcsscommon

import (
	"fmt"
	"strconv"
	"strings"
)

// ServerParams is an object containing all supported server parameters.
type ServerParams struct {
	// TODO: explain each server param behavior
	AudioCutDist              float64
	AutoMode                  bool
	BackDashRate              float64
	BackPasses                bool
	BallAccelMax              float64
	BallDecay                 float64
	BallRand                  float64
	BallSize                  float64
	BallSpeedMax              float64
	BallStuckArea             float64
	BallWeight                float64
	CatchBanCycle             int64
	CatchProbability          float64
	CatchableAreaL            float64
	CatchableAreaW            float64
	CKickMargin               float64
	CLangAdviceWin            int64
	CLangDefineWin            int64
	CLangDelWin               int64
	CLangInfoWin              int64
	CLangMessDelay            int64
	CLangMessPerCycle         int64
	CLangMetaWin              int64
	CLangRuleWin              int64
	CLangWinSize              int64
	Coach                     bool
	CoachPort                 int64
	CoachWReferee             bool
	ConnectWait               int64
	ControlRadius             float64
	DashAngleStep             float64
	DashPowerRate             float64
	DropBallTime              int64
	EffortDec                 float64
	EffortDecThr              float64
	EffortInc                 float64
	EffortIncThr              float64
	EffortInit                float64
	EffortMin                 float64
	ExtraHalfTime             int64
	ExtraStamina              float64
	FixedTeamnameL            string
	FixedTeamnameR            string
	ForbidKickOffOffside      bool
	FoulCycles                int64
	FoulDetectProbability     float64
	FoulExponent              float64
	FreeKickFaults            int64
	FreeformSendPeriod        int64
	FreeformWaitPeriod        int64
	FullstateL                bool
	FullstateR                bool
	GameLogCompression        bool
	GameLogDated              bool
	GameLogDir                string
	GameLogFixed              bool
	GameLogFixedName          string
	GameLogVersion            int64
	GameLogging               bool
	GameOverWait              int64
	GoalWidth                 float64
	GoalieMaxMoves            int64
	GoldenGoal                bool
	HalfTime                  int64 // seconds
	HearDecay                 float64
	HearInc                   float64
	HearMax                   float64
	IllegalDefenseDistX       float64
	IllegalDefenseDuration    int64
	IllegalDefenseNumber      int64
	IllegalDefenseWidth       float64
	InertiaMoment             float64
	Keepaway                  bool
	KeepawayLength            float64
	KeepawayLogDated          bool
	KeepawayLogDir            string
	KeepawayLogFixed          bool
	KeepawayLogFixedName      string
	KeepawayLogging           bool
	KeepawayStart             int64
	KeepawayWidth             float64
	KickOffWait               int64
	KickPowerRate             float64
	KickRand                  float64
	KickRandFactorL           bool
	KickRandFactorR           bool
	KickableMargin            float64
	LandmarkFile              string
	LogDateFormat             string
	LogTimes                  bool
	MaxBackTacklePower        float64
	MaxDashAngle              float64
	MaxDashPower              float64
	MaxGoalKicks              int64
	MaxTacklePower            float64
	MaxMoment                 float64
	MaxNeckAng                float64
	MaxNeckMoment             float64
	MaxPower                  float64
	MinDashAngle              float64
	MinDashPower              float64
	MinMoment                 float64
	MinNeckAng                float64
	MinNeckMoment             float64
	MinPower                  float64
	NrExtraHalfs              int64
	NrNormalHalfs             int64
	OffsideActiveAreaSize     float64
	OffsideKickMargin         float64
	OnlineCoachPort           int64
	OldCoachHear              int64
	PenAllowMultKicks         bool
	PenBeforeSetupWait        int64
	PenCoachMovesPlayers      bool
	PenDistX                  float64
	PenMaxExtraKicks          int64
	PenMaxGoalieDistX         float64
	PenNrKicks                int64
	PenRandomWinner           bool
	PenReadyWait              int64
	PenSetupWait              int64
	PenTakenWait              int64
	PenaltyShootOuts          bool
	PlayerAccelMax            float64
	PlayerDecay               float64
	PlayerRand                float64
	PlayerSize                float64
	PlayerSpeedMax            float64
	PlayerSpeedMaxMin         float64
	PlayerWeight              float64
	PointToBan                int64
	PointToDuration           int64
	Port                      int64
	PRandFactorL              float64
	PRandFactorR              float64
	Profile                   int64
	ProperGoalKicks           int64
	QuantizeStep              float64
	QuantizeStepL             float64
	RecordMessages            bool
	RecoverDec                float64
	RecoverDecThr             float64
	RecoverInit               float64
	RecoverMin                float64
	RecvStep                  int64
	RedCardProbability        float64
	SayCoachCountMax          int64
	SayCoachMsgSize           int64
	SayMsgSize                int64
	SendComms                 int64
	SendStep                  int64
	SendViStep                int64
	SenseBodyStep             int64
	SideDashRate              float64
	SimulatorStep             int64 // in milliseconds
	SlowDownFactor            float64
	SlownessOnTopForLeftTeam  bool
	SlownessOnTopForRightTeam bool
	StaminaCapacity           float64
	StaminaIncMax             float64
	StaminaMax                float64
	StartGoalL                int64
	StartGoalR                int64
	StoppedBallVel            float64
	SynchMicroSleep           int64
	SynchMode                 bool
	SynchOffset               int64
	SynchSeeOffset            int64
	TackleBackDist            float64
	TackleCycles              int64
	TackleDist                float64
	TackleExponent            float64
	TacklePowerRate           float64
	TackleRandFactor          float64
	TackleWidth               float64
	TeamActuatorNoise         int64
	TeamLStart                string
	TeamRStart                string
	TextLogCompression        bool
	TextLogDated              bool
	TextLogDir                string
	TextLogFixed              bool
	TextLogFixedName          string
	TextLogging               bool
	UseOffside                bool
	Verbose                   bool
	VisibleAngle              float64
	VisibleDistance           float64
	WindAng                   float64
	WindDir                   float64
	WindForce                 float64
	WindNone                  bool
	WindRand                  float64
	WindRandom                float64
}

// DefaultServerParams initializes server params with default version 16.0.0 values
func DefaultServerParams() ServerParams {
	return ServerParams{
		AudioCutDist:              50,
		AutoMode:                  false,
		BackDashRate:              0.6,
		BackPasses:                true,
		BallAccelMax:              2.7,
		BallDecay:                 0.94,
		BallRand:                  0.05,
		BallSize:                  0.085,
		BallSpeedMax:              3,
		BallStuckArea:             3,
		BallWeight:                0.2,
		CatchBanCycle:             5,
		CatchProbability:          1,
		CatchableAreaL:            1.2,
		CatchableAreaW:            1,
		CKickMargin:               1,
		CLangAdviceWin:            1,
		CLangDefineWin:            1,
		CLangDelWin:               1,
		CLangInfoWin:              1,
		CLangMessDelay:            50,
		CLangMessPerCycle:         1,
		CLangMetaWin:              1,
		CLangRuleWin:              1,
		CLangWinSize:              300,
		Coach:                     false,
		CoachPort:                 6001,
		CoachWReferee:             false,
		ConnectWait:               300,
		ControlRadius:             2,
		DashAngleStep:             1,
		DashPowerRate:             0.006,
		DropBallTime:              100,
		EffortDec:                 0.005,
		EffortDecThr:              0.3,
		EffortInc:                 0.01,
		EffortIncThr:              0.6,
		EffortInit:                1,
		EffortMin:                 0.6,
		ExtraHalfTime:             100,
		ExtraStamina:              50,
		FixedTeamnameL:            "",
		FixedTeamnameR:            "",
		ForbidKickOffOffside:      true,
		FoulCycles:                5,
		FoulDetectProbability:     0.5,
		FoulExponent:              10,
		FreeKickFaults:            1,
		FreeformSendPeriod:        20,
		FreeformWaitPeriod:        600,
		FullstateL:                false,
		FullstateR:                false,
		GameLogCompression:        false,
		GameLogDated:              true,
		GameLogDir:                "/root/logs",
		GameLogFixed:              false,
		GameLogFixedName:          "rcssserver",
		GameLogVersion:            5,
		GameLogging:               true,
		GameOverWait:              100,
		GoalWidth:                 14.02,
		GoalieMaxMoves:            2,
		GoldenGoal:                false,
		HalfTime:                  300,
		HearDecay:                 1,
		HearInc:                   1,
		HearMax:                   1,
		IllegalDefenseDistX:       16.5,
		IllegalDefenseDuration:    20,
		IllegalDefenseNumber:      0,
		IllegalDefenseWidth:       40.32,
		InertiaMoment:             5,
		Keepaway:                  false,
		KeepawayLength:            20,
		KeepawayLogDated:          true,
		KeepawayLogDir:            "./",
		KeepawayLogFixed:          false,
		KeepawayLogFixedName:      "rcssserver",
		KeepawayLogging:           true,
		KeepawayStart:             -1,
		KeepawayWidth:             20,
		KickOffWait:               100,
		KickPowerRate:             0.027,
		KickRand:                  0.1,
		KickRandFactorL:           true,
		KickRandFactorR:           true,
		KickableMargin:            0.7,
		LandmarkFile:              "~/.rcssserver-landmark.xml",
		LogDateFormat:             "%Y%m%d%H%M%S-",
		LogTimes:                  false,
		MaxBackTacklePower:        100,
		MaxDashAngle:              180,
		MaxDashPower:              100,
		MaxGoalKicks:              3,
		MaxTacklePower:            100,
		MaxMoment:                 180,
		MaxNeckAng:                90,
		MaxNeckMoment:             180,
		MaxPower:                  100,
		MinDashAngle:              -180,
		MinDashPower:              -100,
		MinMoment:                 -180,
		MinNeckAng:                -90,
		MinNeckMoment:             -180,
		MinPower:                  -100,
		NrExtraHalfs:              2,
		NrNormalHalfs:             2,
		OffsideActiveAreaSize:     2.5,
		OffsideKickMargin:         9.15,
		OnlineCoachPort:           6002,
		OldCoachHear:              0,
		PenAllowMultKicks:         true,
		PenBeforeSetupWait:        10,
		PenCoachMovesPlayers:      true,
		PenDistX:                  42.5,
		PenMaxExtraKicks:          5,
		PenMaxGoalieDistX:         14,
		PenNrKicks:                5,
		PenRandomWinner:           false,
		PenReadyWait:              10,
		PenSetupWait:              70,
		PenTakenWait:              150,
		PenaltyShootOuts:          true,
		PlayerAccelMax:            1,
		PlayerDecay:               0.4,
		PlayerRand:                0.1,
		PlayerSize:                0.3,
		PlayerSpeedMax:            1.05,
		PlayerSpeedMaxMin:         0.75,
		PlayerWeight:              60,
		PointToBan:                5,
		PointToDuration:           20,
		Port:                      6000,
		PRandFactorL:              1,
		PRandFactorR:              1,
		Profile:                   0,
		ProperGoalKicks:           0,
		QuantizeStep:              0.1,
		QuantizeStepL:             0.01,
		RecordMessages:            false,
		RecoverDec:                0.002,
		RecoverDecThr:             0.3,
		RecoverInit:               1,
		RecoverMin:                0.5,
		RecvStep:                  10,
		RedCardProbability:        0,
		SayCoachCountMax:          128,
		SayCoachMsgSize:           128,
		SayMsgSize:                10,
		SendComms:                 0,
		SendStep:                  150,
		SendViStep:                100,
		SenseBodyStep:             100,
		SideDashRate:              0.4,
		SimulatorStep:             100,
		SlowDownFactor:            1,
		SlownessOnTopForLeftTeam:  true,
		SlownessOnTopForRightTeam: true,
		StaminaCapacity:           130600,
		StaminaIncMax:             45,
		StaminaMax:                8000,
		StartGoalL:                0,
		StartGoalR:                0,
		StoppedBallVel:            0.01,
		SynchMicroSleep:           1,
		SynchMode:                 false,
		SynchOffset:               60,
		SynchSeeOffset:            0,
		TackleBackDist:            0,
		TackleCycles:              10,
		TackleDist:                2,
		TackleExponent:            6,
		TacklePowerRate:           0.027,
		TackleRandFactor:          2,
		TackleWidth:               1.25,
		TeamActuatorNoise:         0,
		TeamLStart:                "",
		TeamRStart:                "",
		TextLogCompression:        false,
		TextLogDated:              true,
		TextLogDir:                "/root/logs",
		TextLogFixed:              false,
		TextLogFixedName:          "rcssserver",
		TextLogging:               true,
		UseOffside:                true,
		Verbose:                   false,
		VisibleAngle:              90,
		VisibleDistance:           3,
		WindAng:                   0,
		WindDir:                   0,
		WindForce:                 0,
		WindNone:                  false,
		WindRand:                  0,
		WindRandom:                0,
	}
}

// Parse receives a string from server and updates all server params
func (sp *ServerParams) Parse(m string, errCh chan string) {
	trimmedMsg := m
	trimmedMsg = strings.TrimPrefix(trimmedMsg, "(server_param (")
	trimmedMsg = strings.TrimSuffix(trimmedMsg, "))\x00")

	params := strings.Split(trimmedMsg, ")(")

	for _, param := range params {
		paramParts := strings.Split(param, " ")
		if len(paramParts) < 2 {
			errCh <- fmt.Sprintf("invalid server param format: %s", param)
			return
		}
		paramName := paramParts[0]
		paramVal := strings.Join(paramParts[1:len(paramParts)], " ")

		paramVal = strings.TrimPrefix(paramVal, "\"")
		paramVal = strings.TrimSuffix(paramVal, "\"")

		var err error
		switch paramName {
		case "audio_cut_dist":
			sp.AudioCutDist, err = strconv.ParseFloat(paramVal, 64)
		case "auto_mode":
			sp.AutoMode, err = strconv.ParseBool(paramVal)
		case "back_dash_rate":
			sp.BackDashRate, err = strconv.ParseFloat(paramVal, 64)
		case "back_passes":
			sp.BackPasses, err = strconv.ParseBool(paramVal)
		case "ball_accel_max":
			sp.BallAccelMax, err = strconv.ParseFloat(paramVal, 64)
		case "ball_decay":
			sp.BallDecay, err = strconv.ParseFloat(paramVal, 64)
		case "ball_rand":
			sp.BallRand, err = strconv.ParseFloat(paramVal, 64)
		case "ball_size":
			sp.BallSize, err = strconv.ParseFloat(paramVal, 64)
		case "ball_speed_max":
			sp.BallSpeedMax, err = strconv.ParseFloat(paramVal, 64)
		case "ball_stuck_area":
			sp.BallStuckArea, err = strconv.ParseFloat(paramVal, 64)
		case "ball_weight":
			sp.BallWeight, err = strconv.ParseFloat(paramVal, 64)
		case "catch_ban_cycle":
			sp.CatchBanCycle, err = strconv.ParseInt(paramVal, 10, 64)
		case "catch_probability":
			sp.CatchProbability, err = strconv.ParseFloat(paramVal, 64)
		case "catchable_area_l":
			sp.CatchableAreaL, err = strconv.ParseFloat(paramVal, 64)
		case "catchable_area_w":
			sp.CatchableAreaW, err = strconv.ParseFloat(paramVal, 64)
		case "ckick_margin":
			sp.CKickMargin, err = strconv.ParseFloat(paramVal, 64)
		case "clang_advice_win":
			sp.CLangAdviceWin, err = strconv.ParseInt(paramVal, 10, 64)
		case "clang_define_win":
			sp.CLangDefineWin, err = strconv.ParseInt(paramVal, 10, 64)
		case "clang_del_win":
			sp.CLangDelWin, err = strconv.ParseInt(paramVal, 10, 64)
		case "clang_info_win":
			sp.CLangInfoWin, err = strconv.ParseInt(paramVal, 10, 64)
		case "clang_mess_delay":
			sp.CLangMessDelay, err = strconv.ParseInt(paramVal, 10, 64)
		case "clang_mess_per_cycle":
			sp.CLangMessPerCycle, err = strconv.ParseInt(paramVal, 10, 64)
		case "clang_meta_win":
			sp.CLangMetaWin, err = strconv.ParseInt(paramVal, 10, 64)
		case "clang_rule_win":
			sp.CLangRuleWin, err = strconv.ParseInt(paramVal, 10, 64)
		case "clang_win_size":
			sp.CLangWinSize, err = strconv.ParseInt(paramVal, 10, 64)
		case "coach":
			sp.Coach, err = strconv.ParseBool(paramVal)
		case "coach_port":
			sp.CoachPort, err = strconv.ParseInt(paramVal, 10, 64)
		case "coach_w_referee":
			sp.CoachWReferee, err = strconv.ParseBool(paramVal)
		case "connect_wait":
			sp.ConnectWait, err = strconv.ParseInt(paramVal, 10, 64)
		case "control_radius":
			sp.ControlRadius, err = strconv.ParseFloat(paramVal, 64)
		case "dash_angle_step":
			sp.DashAngleStep, err = strconv.ParseFloat(paramVal, 64)
		case "dash_power_rate":
			sp.DashPowerRate, err = strconv.ParseFloat(paramVal, 64)
		case "drop_ball_time":
			sp.DropBallTime, err = strconv.ParseInt(paramVal, 10, 64)
		case "effort_dec":
			sp.EffortDec, err = strconv.ParseFloat(paramVal, 64)
		case "effort_dec_thr":
			sp.EffortDecThr, err = strconv.ParseFloat(paramVal, 64)
		case "effort_inc":
			sp.EffortInc, err = strconv.ParseFloat(paramVal, 64)
		case "effort_inc_thr":
			sp.EffortIncThr, err = strconv.ParseFloat(paramVal, 64)
		case "effort_init":
			sp.EffortInit, err = strconv.ParseFloat(paramVal, 64)
		case "effort_min":
			sp.EffortMin, err = strconv.ParseFloat(paramVal, 64)
		case "extra_half_time":
			sp.ExtraHalfTime, err = strconv.ParseInt(paramVal, 10, 64)
		case "extra_stamina":
			sp.ExtraStamina, err = strconv.ParseFloat(paramVal, 64)
		case "forbid_kick_off_offside":
			sp.ForbidKickOffOffside, err = strconv.ParseBool(paramVal)
		case "foul_cycles":
			sp.FoulCycles, err = strconv.ParseInt(paramVal, 10, 64)
		case "foul_detect_probability":
			sp.FoulDetectProbability, err = strconv.ParseFloat(paramVal, 64)
		case "foul_exponent":
			sp.FoulExponent, err = strconv.ParseFloat(paramVal, 64)
		case "free_kick_faults":
			sp.FreeKickFaults, err = strconv.ParseInt(paramVal, 10, 64)
		case "freeform_send_period":
			sp.FreeformSendPeriod, err = strconv.ParseInt(paramVal, 10, 64)
		case "freeform_wait_period":
			sp.FreeformWaitPeriod, err = strconv.ParseInt(paramVal, 10, 64)
		case "fullstate_l":
			sp.FullstateL, err = strconv.ParseBool(paramVal)
		case "fullstate_r":
			sp.FullstateR, err = strconv.ParseBool(paramVal)
		case "game_log_compression":
			sp.GameLogCompression, err = strconv.ParseBool(paramVal)
		case "game_log_dated":
			sp.GameLogDated, err = strconv.ParseBool(paramVal)
		case "game_log_dir":
			sp.GameLogDir = paramVal
		case "game_log_fixed":
			sp.GameLogFixed, err = strconv.ParseBool(paramVal)
		case "game_log_fixed_name":
			sp.GameLogFixedName = paramVal
		case "game_log_version":
			sp.GameLogVersion, err = strconv.ParseInt(paramVal, 10, 64)
		case "game_logging":
			sp.GameLogging, err = strconv.ParseBool(paramVal)
		case "game_over_wait":
			sp.GameOverWait, err = strconv.ParseInt(paramVal, 10, 64)
		case "goal_width":
			sp.GoalWidth, err = strconv.ParseFloat(paramVal, 64)
		case "goalie_max_moves":
			sp.GoalieMaxMoves, err = strconv.ParseInt(paramVal, 10, 64)
		case "golden_goal":
			sp.GoldenGoal, err = strconv.ParseBool(paramVal)
		case "half_time":
			sp.HalfTime, err = strconv.ParseInt(paramVal, 10, 64)
		case "hear_decay":
			sp.HearDecay, err = strconv.ParseFloat(paramVal, 64)
		case "hear_inc":
			sp.HearInc, err = strconv.ParseFloat(paramVal, 64)
		case "hear_max":
			sp.HearMax, err = strconv.ParseFloat(paramVal, 64)
		case "inertia_moment":
			sp.InertiaMoment, err = strconv.ParseFloat(paramVal, 64)
		case "keepaway":
			sp.Keepaway, err = strconv.ParseBool(paramVal)
		case "keepaway_length":
			sp.KeepawayLength, err = strconv.ParseFloat(paramVal, 64)
		case "keepaway_log_dated":
			sp.KeepawayLogDated, err = strconv.ParseBool(paramVal)
		case "keepaway_log_dir":
			sp.KeepawayLogDir = paramVal
		case "keepaway_log_fixed":
			sp.KeepawayLogFixed, err = strconv.ParseBool(paramVal)
		case "keepaway_log_fixed_name":
			sp.KeepawayLogFixedName = paramVal
		case "keepaway_logging":
			sp.KeepawayLogging, err = strconv.ParseBool(paramVal)
		case "keepaway_start":
			sp.KeepawayStart, err = strconv.ParseInt(paramVal, 10, 64)
		case "keepaway_width":
			sp.KeepawayWidth, err = strconv.ParseFloat(paramVal, 64)
		case "kick_off_wait":
			sp.KickOffWait, err = strconv.ParseInt(paramVal, 10, 64)
		case "kick_power_rate":
			sp.KickPowerRate, err = strconv.ParseFloat(paramVal, 64)
		case "kick_rand":
			sp.KickRand, err = strconv.ParseFloat(paramVal, 64)
		case "kick_rand_factor_l":
			sp.KickRandFactorL, err = strconv.ParseBool(paramVal)
		case "kick_rand_factor_r":
			sp.KickRandFactorR, err = strconv.ParseBool(paramVal)
		case "kickable_margin":
			sp.KickableMargin, err = strconv.ParseFloat(paramVal, 64)
		case "landmark_file":
			sp.LandmarkFile = paramVal
		case "log_date_format":
			sp.LogDateFormat = paramVal
		case "log_times":
			sp.LogTimes, err = strconv.ParseBool(paramVal)
		case "max_back_tackle_power":
			sp.MaxBackTacklePower, err = strconv.ParseFloat(paramVal, 64)
		case "max_dash_angle":
			sp.MaxDashAngle, err = strconv.ParseFloat(paramVal, 64)
		case "max_dash_power":
			sp.MaxDashPower, err = strconv.ParseFloat(paramVal, 64)
		case "max_goal_kicks":
			sp.MaxGoalKicks, err = strconv.ParseInt(paramVal, 10, 64)
		case "max_tackle_power":
			sp.MaxTacklePower, err = strconv.ParseFloat(paramVal, 64)
		case "maxmoment":
			sp.MaxMoment, err = strconv.ParseFloat(paramVal, 64)
		case "maxneckang":
			sp.MaxNeckAng, err = strconv.ParseFloat(paramVal, 64)
		case "maxneckmoment":
			sp.MaxNeckMoment, err = strconv.ParseFloat(paramVal, 64)
		case "maxpower":
			sp.MaxPower, err = strconv.ParseFloat(paramVal, 64)
		case "min_dash_angle":
			sp.MinDashAngle, err = strconv.ParseFloat(paramVal, 64)
		case "min_dash_power":
			sp.MinDashPower, err = strconv.ParseFloat(paramVal, 64)
		case "minmoment":
			sp.MinMoment, err = strconv.ParseFloat(paramVal, 64)
		case "minneckang":
			sp.MinNeckAng, err = strconv.ParseFloat(paramVal, 64)
		case "minneckmoment":
			sp.MinNeckMoment, err = strconv.ParseFloat(paramVal, 64)
		case "minpower":
			sp.MinPower, err = strconv.ParseFloat(paramVal, 64)
		case "nr_extra_halfs":
			sp.NrExtraHalfs, err = strconv.ParseInt(paramVal, 10, 64)
		case "nr_normal_halfs":
			sp.NrNormalHalfs, err = strconv.ParseInt(paramVal, 10, 64)
		case "offside_active_area_size":
			sp.OffsideActiveAreaSize, err = strconv.ParseFloat(paramVal, 64)
		case "offside_kick_margin":
			sp.OffsideKickMargin, err = strconv.ParseFloat(paramVal, 64)
		case "olcoach_port":
			sp.OnlineCoachPort, err = strconv.ParseInt(paramVal, 10, 64)
		case "old_coach_hear":
			sp.OldCoachHear, err = strconv.ParseInt(paramVal, 10, 64)
		case "pen_allow_mult_kicks":
			sp.PenAllowMultKicks, err = strconv.ParseBool(paramVal)
		case "pen_before_setup_wait":
			sp.PenBeforeSetupWait, err = strconv.ParseInt(paramVal, 10, 64)
		case "pen_coach_moves_players":
			sp.PenCoachMovesPlayers, err = strconv.ParseBool(paramVal)
		case "pen_dist_x":
			sp.PenDistX, err = strconv.ParseFloat(paramVal, 64)
		case "pen_max_extra_kicks":
			sp.PenMaxExtraKicks, err = strconv.ParseInt(paramVal, 10, 64)
		case "pen_max_goalie_dist_x":
			sp.PenMaxGoalieDistX, err = strconv.ParseFloat(paramVal, 64)
		case "pen_nr_kicks":
			sp.PenNrKicks, err = strconv.ParseInt(paramVal, 10, 64)
		case "pen_random_winner":
			sp.PenRandomWinner, err = strconv.ParseBool(paramVal)
		case "pen_ready_wait":
			sp.PenReadyWait, err = strconv.ParseInt(paramVal, 10, 64)
		case "pen_setup_wait":
			sp.PenSetupWait, err = strconv.ParseInt(paramVal, 10, 64)
		case "pen_taken_wait":
			sp.PenTakenWait, err = strconv.ParseInt(paramVal, 10, 64)
		case "penalty_shoot_outs":
			sp.PenaltyShootOuts, err = strconv.ParseBool(paramVal)
		case "player_accel_max":
			sp.PlayerAccelMax, err = strconv.ParseFloat(paramVal, 64)
		case "player_decay":
			sp.PlayerDecay, err = strconv.ParseFloat(paramVal, 64)
		case "player_rand":
			sp.PlayerRand, err = strconv.ParseFloat(paramVal, 64)
		case "player_size":
			sp.PlayerSize, err = strconv.ParseFloat(paramVal, 64)
		case "player_speed_max":
			sp.PlayerSpeedMax, err = strconv.ParseFloat(paramVal, 64)
		case "player_speed_max_min":
			sp.PlayerSpeedMaxMin, err = strconv.ParseFloat(paramVal, 64)
		case "player_weight":
			sp.PlayerWeight, err = strconv.ParseFloat(paramVal, 64)
		case "point_to_ban":
			sp.PointToBan, err = strconv.ParseInt(paramVal, 10, 64)
		case "point_to_duration":
			sp.PointToDuration, err = strconv.ParseInt(paramVal, 10, 64)
		case "port":
			sp.Port, err = strconv.ParseInt(paramVal, 10, 64)
		case "prand_factor_l":
			sp.PRandFactorL, err = strconv.ParseFloat(paramVal, 64)
		case "prand_factor_r":
			sp.PRandFactorR, err = strconv.ParseFloat(paramVal, 64)
		case "profile":
			sp.Profile, err = strconv.ParseInt(paramVal, 10, 64)
		case "proper_goal_kicks":
			sp.ProperGoalKicks, err = strconv.ParseInt(paramVal, 10, 64)
		case "quantize_step":
			sp.QuantizeStep, err = strconv.ParseFloat(paramVal, 64)
		case "quantize_step_l":
			sp.QuantizeStepL, err = strconv.ParseFloat(paramVal, 64)
		case "record_messages":
			sp.RecordMessages, err = strconv.ParseBool(paramVal)
		case "recover_dec":
			sp.RecoverDec, err = strconv.ParseFloat(paramVal, 64)
		case "recover_dec_thr":
			sp.RecoverDecThr, err = strconv.ParseFloat(paramVal, 64)
		case "recover_init":
			sp.RecoverInit, err = strconv.ParseFloat(paramVal, 64)
		case "recover_min":
			sp.RecoverMin, err = strconv.ParseFloat(paramVal, 64)
		case "recv_step":
			sp.RecvStep, err = strconv.ParseInt(paramVal, 10, 64)
		case "say_coach_cnt_max":
			sp.SayCoachCountMax, err = strconv.ParseInt(paramVal, 10, 64)
		case "say_coach_msg_size":
			sp.SayCoachMsgSize, err = strconv.ParseInt(paramVal, 10, 64)
		case "say_msg_size":
			sp.SayMsgSize, err = strconv.ParseInt(paramVal, 10, 64)
		case "send_comms":
			sp.SendComms, err = strconv.ParseInt(paramVal, 10, 64)
		case "send_step":
			sp.SendStep, err = strconv.ParseInt(paramVal, 10, 64)
		case "send_vi_step":
			sp.SendViStep, err = strconv.ParseInt(paramVal, 10, 64)
		case "sense_body_step":
			sp.SenseBodyStep, err = strconv.ParseInt(paramVal, 10, 64)
		case "side_dash_rate":
			sp.SideDashRate, err = strconv.ParseFloat(paramVal, 64)
		case "simulator_step":
			sp.SimulatorStep, err = strconv.ParseInt(paramVal, 10, 64)
		case "slow_down_factor":
			sp.SlowDownFactor, err = strconv.ParseFloat(paramVal, 64)
		case "slowness_on_top_for_left_team":
			sp.SlownessOnTopForLeftTeam, err = strconv.ParseBool(paramVal)
		case "slowness_on_top_for_right_team":
			sp.SlownessOnTopForRightTeam, err = strconv.ParseBool(paramVal)
		case "stamina_capacity":
			sp.StaminaCapacity, err = strconv.ParseFloat(paramVal, 64)
		case "stamina_inc_max":
			sp.StaminaIncMax, err = strconv.ParseFloat(paramVal, 64)
		case "stamina_max":
			sp.StaminaMax, err = strconv.ParseFloat(paramVal, 64)
		case "start_goal_l":
			sp.StartGoalL, err = strconv.ParseInt(paramVal, 10, 64)
		case "start_goal_r":
			sp.StartGoalR, err = strconv.ParseInt(paramVal, 10, 64)
		case "stopped_ball_vel":
			sp.StoppedBallVel, err = strconv.ParseFloat(paramVal, 64)
		case "synch_micro_sleep":
			sp.SynchMicroSleep, err = strconv.ParseInt(paramVal, 10, 64)
		case "synch_mode":
			sp.SynchMode, err = strconv.ParseBool(paramVal)
		case "synch_offset":
			sp.SynchOffset, err = strconv.ParseInt(paramVal, 10, 64)
		case "synch_see_offset":
			sp.SynchSeeOffset, err = strconv.ParseInt(paramVal, 10, 64)
		case "tackle_back_dist":
			sp.TackleBackDist, err = strconv.ParseFloat(paramVal, 64)
		case "tackle_cycles":
			sp.TackleCycles, err = strconv.ParseInt(paramVal, 10, 64)
		case "tackle_dist":
			sp.TackleDist, err = strconv.ParseFloat(paramVal, 64)
		case "tackle_exponent":
			sp.TackleExponent, err = strconv.ParseFloat(paramVal, 64)
		case "tackle_power_rate":
			sp.TacklePowerRate, err = strconv.ParseFloat(paramVal, 64)
		case "tackle_rand_factor":
			sp.TackleRandFactor, err = strconv.ParseFloat(paramVal, 64)
		case "tackle_width":
			sp.TackleWidth, err = strconv.ParseFloat(paramVal, 64)
		case "team_actuator_noise":
			sp.TeamActuatorNoise, err = strconv.ParseInt(paramVal, 10, 64)
		case "team_l_start":
			sp.TeamLStart = paramVal
		case "team_r_start":
			sp.TeamRStart = paramVal
		case "text_log_compression":
			sp.TextLogCompression, err = strconv.ParseBool(paramVal)
		case "text_log_dated":
			sp.TextLogDated, err = strconv.ParseBool(paramVal)
		case "text_log_dir":
			sp.TextLogDir = paramVal
		case "text_log_fixed":
			sp.TextLogFixed, err = strconv.ParseBool(paramVal)
		case "text_log_fixed_name":
			sp.TextLogFixedName = paramVal
		case "text_logging":
			sp.TextLogging, err = strconv.ParseBool(paramVal)
		case "use_offside":
			sp.UseOffside, err = strconv.ParseBool(paramVal)
		case "verbose":
			sp.Verbose, err = strconv.ParseBool(paramVal)
		case "visible_angle":
			sp.VisibleAngle, err = strconv.ParseFloat(paramVal, 64)
		case "visible_distance":
			sp.VisibleDistance, err = strconv.ParseFloat(paramVal, 64)
		case "wind_ang":
			sp.WindAng, err = strconv.ParseFloat(paramVal, 64)
		case "wind_dir":
			sp.WindDir, err = strconv.ParseFloat(paramVal, 64)
		case "wind_force":
			sp.WindForce, err = strconv.ParseFloat(paramVal, 64)
		case "wind_none":
			sp.WindNone, err = strconv.ParseBool(paramVal)
		case "wind_rand":
			sp.WindRand, err = strconv.ParseFloat(paramVal, 64)
		case "wind_random":
			sp.WindRandom, err = strconv.ParseFloat(paramVal, 64)
		default:
			errCh <- fmt.Sprintf("unsupported server param (%s)", param)
		}
		if err != nil {
			errCh <- fmt.Sprintf("could not parse server param (%s): %s", param, err)
			err = nil
		}
	}
}
