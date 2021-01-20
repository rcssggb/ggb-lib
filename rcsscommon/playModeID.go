package rcsscommon

// PlayModeID is the type representing each game mode
type PlayModeID byte

const (
	PlayModeUnknown             PlayModeID = iota + 0 // unknown play mode
	PlayModeBeforeKickOff                             // before_kick_off
	PlayModePlayOn                                    // play_on
	PlayModeTimeOver                                  // time_over
	PlayModeKickOffR                                  // kick_off_r
	PlayModeKickOffL                                  // kick_off_l
	PlayModeKickInR                                   // kick_in_r
	PlayModeKickInL                                   // kick_in_l
	PlayModeFreeKickR                                 // free_kick_r
	PlayModeFreeKickL                                 // free_kick_l
	PlayModeCornerKickR                               // corner_kick_r
	PlayModeCornerKickL                               // corner_kick_l
	PlayModeGoalKickR                                 // goal_kick_r
	PlayModeGoalKickL                                 // goal_kick_l
	PlayModeGoalR                                     // goal_r
	PlayModeGoalL                                     // goal_l
	PlayModeDropBall                                  // drop_ball
	PlayModeOffsideR                                  // offside_r
	PlayModeOffsideL                                  // offside_l
	PlayModePenaltyKickL                              // penalty_kick_l
	PlayModePenaltyKickR                              // penalty_kick_r
	PlayModeFirstHalfOver                             // first_half_over
	PlayModePause                                     // pause
	PlayModeHumanJudge                                // human_judge
	PlayModeFoulChargeL                               // foul_charge_l
	PlayModeFoulChargeR                               // foul_charge_r
	PlayModeFoulPushL                                 // foul_push_l
	PlayModeFoulPushR                                 // foul_push_r
	PlayModeFoulMultipleAttackL                       // foul_multiple_attack_l
	PlayModeFoulMultipleAttackR                       // foul_multiple_attack_r
	PlayModeFoulBalloutL                              // foul_ballout_l
	PlayModeFoulBalloutR                              // foul_ballout_r
	PlayModeBackPassL                                 // back_pass_l
	PlayModeBackPassR                                 // back_pass_r
	PlayModeFreeKickFaultL                            // free_kick_fault_l
	PlayModeFreeKickFaultR                            // free_kick_fault_r
	PlayModeCatchFaultL                               // catch_fault_l
	PlayModeCatchFaultR                               // catch_fault_r
	PlayModeIndirectFreeKickL                         // indirect_free_kick_l
	PlayModeIndirectFreeKickR                         // indirect_free_kick_r
	PlayModePenaltySetupL                             // penalty_setup_l
	PlayModePenaltySetupR                             // penalty_setup_r
	PlayModePenaltyReadyL                             // penalty_ready_l
	PlayModePenaltyReadyR                             // penalty_ready_r
	PlayModePenaltyTakenL                             // penalty_taken_l
	PlayModePenaltyTakenR                             // penalty_taken_r
	PlayModePenaltyMissL                              // penalty_miss_l
	PlayModePenaltyMissR                              // penalty_miss_r
	PlayModePenaltyScoreL                             // penalty_score_l
	PlayModePenaltyScoreR                             // penalty_score_r
	PlayModeIllegalDefenseL                           // illegal_defense_l
	PlayModeIllegalDefenseR                           // illegal_defense_r
)

// NewPlayModeID parses a string and returns a PlayModeID
func NewPlayModeID(playMode string) PlayModeID {
	switch playMode {
	case "before_kick_off":
		return PlayModeBeforeKickOff
	case "play_on":
		return PlayModePlayOn
	case "time_over":
		return PlayModeTimeOver
	case "kick_off_r":
		return PlayModeKickOffR
	case "kick_off_l":
		return PlayModeKickOffL
	case "kick_in_r":
		return PlayModeKickInR
	case "kick_in_l":
		return PlayModeKickInL
	case "free_kick_r":
		return PlayModeFreeKickR
	case "free_kick_l":
		return PlayModeFreeKickL
	case "corner_kick_r":
		return PlayModeCornerKickR
	case "corner_kick_l":
		return PlayModeCornerKickL
	case "goal_kick_r":
		return PlayModeGoalKickR
	case "goal_kick_l":
		return PlayModeGoalKickL
	case "goal_r":
		return PlayModeGoalR
	case "goal_l":
		return PlayModeGoalL
	case "drop_ball":
		return PlayModeDropBall
	case "offside_r":
		return PlayModeOffsideR
	case "offside_l":
		return PlayModeOffsideL
	case "penalty_kick_l":
		return PlayModePenaltyKickL
	case "penalty_kick_r":
		return PlayModePenaltyKickR
	case "first_half_over":
		return PlayModeFirstHalfOver
	case "pause":
		return PlayModePause
	case "human_judge":
		return PlayModeHumanJudge
	case "foul_charge_l":
		return PlayModeFoulChargeL
	case "foul_charge_r":
		return PlayModeFoulChargeR
	case "foul_push_l":
		return PlayModeFoulPushL
	case "foul_push_r":
		return PlayModeFoulPushR
	case "foul_multiple_attack_l":
		return PlayModeFoulMultipleAttackL
	case "foul_multiple_attack_r":
		return PlayModeFoulMultipleAttackR
	case "foul_ballout_l":
		return PlayModeFoulBalloutL
	case "foul_ballout_r":
		return PlayModeFoulBalloutR
	case "back_pass_l":
		return PlayModeBackPassL
	case "back_pass_r":
		return PlayModeBackPassR
	case "free_kick_fault_l":
		return PlayModeFreeKickFaultL
	case "free_kick_fault_r":
		return PlayModeFreeKickFaultR
	case "catch_fault_l":
		return PlayModeCatchFaultL
	case "catch_fault_r":
		return PlayModeCatchFaultR
	case "indirect_free_kick_l":
		return PlayModeIndirectFreeKickL
	case "indirect_free_kick_r":
		return PlayModeIndirectFreeKickR
	case "penalty_setup_l":
		return PlayModePenaltySetupL
	case "penalty_setup_r":
		return PlayModePenaltySetupR
	case "penalty_ready_l":
		return PlayModePenaltyReadyL
	case "penalty_ready_r":
		return PlayModePenaltyReadyR
	case "penalty_taken_l":
		return PlayModePenaltyTakenL
	case "penalty_taken_r":
		return PlayModePenaltyTakenR
	case "penalty_miss_l":
		return PlayModePenaltyMissL
	case "penalty_miss_r":
		return PlayModePenaltyMissR
	case "penalty_score_l":
		return PlayModePenaltyScoreL
	case "penalty_score_r":
		return PlayModePenaltyScoreR
	case "illegal_defense_l":
		return PlayModeIllegalDefenseL
	case "illegal_defense_r":
		return PlayModeIllegalDefenseR
	default:
		return PlayModeUnknown
	}
}

// OneHot returns play mode encoded in a onehot vector
func (id PlayModeID) OneHot() []float64 {
	ret := make([]float64, 51)
	switch id {
	case PlayModeBeforeKickOff:
		ret[0] = 1
	case PlayModePlayOn:
		ret[1] = 1
	case PlayModeTimeOver:
		ret[2] = 1
	case PlayModeKickOffR:
		ret[3] = 1
	case PlayModeKickOffL:
		ret[4] = 1
	case PlayModeKickInR:
		ret[5] = 1
	case PlayModeKickInL:
		ret[6] = 1
	case PlayModeFreeKickR:
		ret[7] = 1
	case PlayModeFreeKickL:
		ret[8] = 1
	case PlayModeCornerKickR:
		ret[9] = 1
	case PlayModeCornerKickL:
		ret[10] = 1
	case PlayModeGoalKickR:
		ret[11] = 1
	case PlayModeGoalKickL:
		ret[12] = 1
	case PlayModeGoalR:
		ret[13] = 1
	case PlayModeGoalL:
		ret[14] = 1
	case PlayModeDropBall:
		ret[15] = 1
	case PlayModeOffsideR:
		ret[16] = 1
	case PlayModeOffsideL:
		ret[17] = 1
	case PlayModePenaltyKickL:
		ret[18] = 1
	case PlayModePenaltyKickR:
		ret[19] = 1
	case PlayModeFirstHalfOver:
		ret[20] = 1
	case PlayModePause:
		ret[21] = 1
	case PlayModeHumanJudge:
		ret[22] = 1
	case PlayModeFoulChargeL:
		ret[23] = 1
	case PlayModeFoulChargeR:
		ret[24] = 1
	case PlayModeFoulPushL:
		ret[25] = 1
	case PlayModeFoulPushR:
		ret[26] = 1
	case PlayModeFoulMultipleAttackL:
		ret[27] = 1
	case PlayModeFoulMultipleAttackR:
		ret[28] = 1
	case PlayModeFoulBalloutL:
		ret[29] = 1
	case PlayModeFoulBalloutR:
		ret[30] = 1
	case PlayModeBackPassL:
		ret[31] = 1
	case PlayModeBackPassR:
		ret[32] = 1
	case PlayModeFreeKickFaultL:
		ret[33] = 1
	case PlayModeFreeKickFaultR:
		ret[34] = 1
	case PlayModeCatchFaultL:
		ret[35] = 1
	case PlayModeCatchFaultR:
		ret[36] = 1
	case PlayModeIndirectFreeKickL:
		ret[37] = 1
	case PlayModeIndirectFreeKickR:
		ret[38] = 1
	case PlayModePenaltySetupL:
		ret[39] = 1
	case PlayModePenaltySetupR:
		ret[40] = 1
	case PlayModePenaltyReadyL:
		ret[41] = 1
	case PlayModePenaltyReadyR:
		ret[42] = 1
	case PlayModePenaltyTakenL:
		ret[43] = 1
	case PlayModePenaltyTakenR:
		ret[44] = 1
	case PlayModePenaltyMissL:
		ret[45] = 1
	case PlayModePenaltyMissR:
		ret[46] = 1
	case PlayModePenaltyScoreL:
		ret[47] = 1
	case PlayModePenaltyScoreR:
		ret[48] = 1
	case PlayModeIllegalDefenseL:
		ret[49] = 1
	case PlayModeIllegalDefenseR:
		ret[50] = 1
	}
	return ret
}

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
	case PlayModePenaltyKickL:
		return "penalty_kick_l"
	case PlayModePenaltyKickR:
		return "penalty_kick_r"
	case PlayModeFirstHalfOver:
		return "first_half_over"
	case PlayModePause:
		return "pause"
	case PlayModeHumanJudge:
		return "human_judge"
	case PlayModeFoulChargeL:
		return "foul_charge_l"
	case PlayModeFoulChargeR:
		return "foul_charge_r"
	case PlayModeFoulPushL:
		return "foul_push_l"
	case PlayModeFoulPushR:
		return "foul_push_r"
	case PlayModeFoulMultipleAttackL:
		return "foul_multiple_attack_l"
	case PlayModeFoulMultipleAttackR:
		return "foul_multiple_attack_r"
	case PlayModeFoulBalloutL:
		return "foul_ballout_l"
	case PlayModeFoulBalloutR:
		return "foul_ballout_r"
	case PlayModeBackPassL:
		return "back_pass_l"
	case PlayModeBackPassR:
		return "back_pass_r"
	case PlayModeFreeKickFaultL:
		return "free_kick_fault_l"
	case PlayModeFreeKickFaultR:
		return "free_kick_fault_r"
	case PlayModeCatchFaultL:
		return "catch_fault_l"
	case PlayModeCatchFaultR:
		return "catch_fault_r"
	case PlayModeIndirectFreeKickL:
		return "indirect_free_kick_l"
	case PlayModeIndirectFreeKickR:
		return "indirect_free_kick_r"
	case PlayModePenaltySetupL:
		return "penalty_setup_l"
	case PlayModePenaltySetupR:
		return "penalty_setup_r"
	case PlayModePenaltyReadyL:
		return "penalty_ready_l"
	case PlayModePenaltyReadyR:
		return "penalty_ready_r"
	case PlayModePenaltyTakenL:
		return "penalty_taken_l"
	case PlayModePenaltyTakenR:
		return "penalty_taken_r"
	case PlayModePenaltyMissL:
		return "penalty_miss_l"
	case PlayModePenaltyMissR:
		return "penalty_miss_r"
	case PlayModePenaltyScoreL:
		return "penalty_score_l"
	case PlayModePenaltyScoreR:
		return "penalty_score_r"
	case PlayModeIllegalDefenseL:
		return "illegal_defense_l"
	case PlayModeIllegalDefenseR:
		return "illegal_defense_r"
	default:
		return ""
	}
}
