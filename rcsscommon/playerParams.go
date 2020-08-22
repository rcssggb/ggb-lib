package rcsscommon

import (
	"fmt"
	"strconv"
	"strings"
)

// PlayerParams is an object containing all player.conf parameters.
type PlayerParams struct {
	// TODO: explain each player param behavior
	AllowMultDefaultType             bool
	CatchableAreaLStretchMax         float64
	CatchableAreaLStretchMin         float64
	DashPowerRateDeltaMax            float64
	DashPowerRateDeltaMin            float64
	EffortMaxDeltaFactor             float64
	EffortMinDeltaFactor             float64
	ExtraStaminaDeltaMax             float64
	ExtraStaminaDeltaMin             float64
	FoulDetectProbabilityDeltaFactor float64
	InertiaMomentDeltaFactor         float64
	KickPowerRateDeltaMax            float64
	KickPowerRateDeltaMin            float64
	KickRandDeltaFactor              float64
	KickableMarginDeltaMax           float64
	KickableMarginDeltaMin           float64
	NewDashPowerRateDeltaMax         float64
	NewDashPowerRateDeltaMin         float64
	NewStaminaIncMaxDeltaFactor      float64
	PlayerDecayDeltaMax              float64
	PlayerDecayDeltaMin              float64
	PlayerSizeDeltaFactor            float64
	PlayerSpeedMaxDeltaMax           float64
	PlayerSpeedMaxDeltaMin           float64
	PlayerTypes                      int64
	PtMax                            int64
	RandomSeed                       int64
	StaminaIncMaxDeltaFactor         float64
	SubsMax                          int64
}

// DefaultPlayerParams initializes player params with default version 16.0.0 values
func DefaultPlayerParams() PlayerParams {
	return PlayerParams{
		AllowMultDefaultType:             false,
		CatchableAreaLStretchMax:         1.3,
		CatchableAreaLStretchMin:         1,
		DashPowerRateDeltaMax:            0,
		DashPowerRateDeltaMin:            0,
		EffortMaxDeltaFactor:             -0.004,
		EffortMinDeltaFactor:             -0.004,
		ExtraStaminaDeltaMax:             50,
		ExtraStaminaDeltaMin:             0,
		FoulDetectProbabilityDeltaFactor: 0,
		InertiaMomentDeltaFactor:         25,
		KickPowerRateDeltaMax:            0,
		KickPowerRateDeltaMin:            0,
		KickRandDeltaFactor:              1,
		KickableMarginDeltaMax:           0.1,
		KickableMarginDeltaMin:           -0.1,
		NewDashPowerRateDeltaMax:         0.0008,
		NewDashPowerRateDeltaMin:         -0.0012,
		NewStaminaIncMaxDeltaFactor:      -6000,
		PlayerDecayDeltaMax:              0.1,
		PlayerDecayDeltaMin:              -0.1,
		PlayerSizeDeltaFactor:            -100,
		PlayerSpeedMaxDeltaMax:           0,
		PlayerSpeedMaxDeltaMin:           0,
		PlayerTypes:                      18,
		PtMax:                            1,
		StaminaIncMaxDeltaFactor:         0,
		SubsMax:                          3,
	}
}

// Parse receives a string from server and updates all player params
func (p *PlayerParams) Parse(m string, errCh chan string) {
	trimmedMsg := m
	trimmedMsg = strings.TrimPrefix(trimmedMsg, "(player_param (")
	trimmedMsg = strings.TrimSuffix(trimmedMsg, "))\x00")

	params := strings.Split(trimmedMsg, ")(")

	for _, param := range params {
		paramParts := strings.Split(param, " ")
		if len(paramParts) < 2 {
			errCh <- fmt.Sprintf("invalid server param format: %s", param)
			continue
		}
		paramName := paramParts[0]
		paramVal := strings.Join(paramParts[1:], " ")

		paramVal = strings.TrimPrefix(paramVal, "\"")
		paramVal = strings.TrimSuffix(paramVal, "\"")

		var err error
		switch paramName {
		case "allow_mult_default_type":
			p.AllowMultDefaultType, err = strconv.ParseBool(paramVal)
		case "catchable_area_l_stretch_max":
			p.CatchableAreaLStretchMax, err = strconv.ParseFloat(paramVal, 64)
		case "catchable_area_l_stretch_min":
			p.CatchableAreaLStretchMin, err = strconv.ParseFloat(paramVal, 64)
		case "dash_power_rate_delta_max":
			p.DashPowerRateDeltaMax, err = strconv.ParseFloat(paramVal, 64)
		case "dash_power_rate_delta_min":
			p.DashPowerRateDeltaMin, err = strconv.ParseFloat(paramVal, 64)
		case "effort_max_delta_factor":
			p.EffortMaxDeltaFactor, err = strconv.ParseFloat(paramVal, 64)
		case "effort_min_delta_factor":
			p.EffortMinDeltaFactor, err = strconv.ParseFloat(paramVal, 64)
		case "extra_stamina_delta_max":
			p.ExtraStaminaDeltaMax, err = strconv.ParseFloat(paramVal, 64)
		case "extra_stamina_delta_min":
			p.ExtraStaminaDeltaMin, err = strconv.ParseFloat(paramVal, 64)
		case "foul_detect_probability_delta_factor":
			p.FoulDetectProbabilityDeltaFactor, err = strconv.ParseFloat(paramVal, 64)
		case "inertia_moment_delta_factor":
			p.InertiaMomentDeltaFactor, err = strconv.ParseFloat(paramVal, 64)
		case "kick_power_rate_delta_max":
			p.KickPowerRateDeltaMax, err = strconv.ParseFloat(paramVal, 64)
		case "kick_power_rate_delta_min":
			p.KickPowerRateDeltaMin, err = strconv.ParseFloat(paramVal, 64)
		case "kick_rand_delta_factor":
			p.KickRandDeltaFactor, err = strconv.ParseFloat(paramVal, 64)
		case "kickable_margin_delta_max":
			p.KickableMarginDeltaMax, err = strconv.ParseFloat(paramVal, 64)
		case "kickable_margin_delta_min":
			p.KickableMarginDeltaMin, err = strconv.ParseFloat(paramVal, 64)
		case "new_dash_power_rate_delta_max":
			p.NewDashPowerRateDeltaMax, err = strconv.ParseFloat(paramVal, 64)
		case "new_dash_power_rate_delta_min":
			p.NewDashPowerRateDeltaMin, err = strconv.ParseFloat(paramVal, 64)
		case "new_stamina_inc_max_delta_factor":
			p.NewStaminaIncMaxDeltaFactor, err = strconv.ParseFloat(paramVal, 64)
		case "player_decay_delta_max":
			p.PlayerDecayDeltaMax, err = strconv.ParseFloat(paramVal, 64)
		case "player_decay_delta_min":
			p.PlayerDecayDeltaMin, err = strconv.ParseFloat(paramVal, 64)
		case "player_size_delta_factor":
			p.PlayerSizeDeltaFactor, err = strconv.ParseFloat(paramVal, 64)
		case "player_speed_max_delta_max":
			p.PlayerSpeedMaxDeltaMax, err = strconv.ParseFloat(paramVal, 64)
		case "player_speed_max_delta_min":
			p.PlayerSpeedMaxDeltaMin, err = strconv.ParseFloat(paramVal, 64)
		case "player_types":
			p.PlayerTypes, err = strconv.ParseInt(paramVal, 10, 64)
		case "pt_max":
			p.PtMax, err = strconv.ParseInt(paramVal, 10, 64)
		case "random_seed":
			p.RandomSeed, err = strconv.ParseInt(paramVal, 10, 64)
		case "stamina_inc_max_delta_factor":
			p.StaminaIncMaxDeltaFactor, err = strconv.ParseFloat(paramVal, 64)
		case "subs_max":
			p.SubsMax, err = strconv.ParseInt(paramVal, 10, 64)
		default:
			errCh <- fmt.Sprintf("unsupported player param (%s)", param)
		}
		if err != nil {
			errCh <- fmt.Sprintf("could not parse player param (%s): %s", param, err)
			err = nil
		}
	}
}
