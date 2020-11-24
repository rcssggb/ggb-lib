package rcsscommon

// PlayModeID is the type representing each game mode
type PlayModeID byte

const (
	PlayModeBeforeKickOff       PlayModeID = iota + 0 // before_kick_off
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
