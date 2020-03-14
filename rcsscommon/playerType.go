package rcsscommon

import (
	"fmt"
	"strconv"
	"strings"
)

// PlayerType is the type containing all player type parameters
type PlayerType struct {
	ID                    int64
	PlayerSpeedMax        float64
	StaminaIncMax         float64
	PlayerDecay           float64
	InertiaMoment         float64
	DashPowerRate         float64
	PlayerSize            float64
	KickableMargin        float64
	KickRand              float64
	ExtraStamina          float64
	EffortMax             float64
	EffortMin             float64
	KickPowerRate         float64
	FoulDetectProbability float64
	CatchableAreaLStretch float64
}

// ParsePlayerType parses player type message and returns
func ParsePlayerType(m string, errCh chan string) PlayerType {
	playerType := PlayerType{
		ID: -1,
	}
	trimmedMsg := m
	trimmedMsg = strings.TrimPrefix(trimmedMsg, "(player_type (")
	trimmedMsg = strings.TrimSuffix(trimmedMsg, "))\x00")

	params := strings.Split(trimmedMsg, ")(")

	for _, param := range params {
		paramParts := strings.Split(param, " ")
		if len(paramParts) != 2 {
			errCh <- fmt.Sprintf("invalid server param format: %s", param)
			continue
		}
		paramName := paramParts[0]
		paramVal := paramParts[1]

		var err error
		switch paramName {
		case "id":
			playerType.ID, err = strconv.ParseInt(paramVal, 10, 64)
		case "player_speed_max":
			playerType.PlayerSpeedMax, err = strconv.ParseFloat(paramVal, 64)
		case "stamina_inc_max":
			playerType.StaminaIncMax, err = strconv.ParseFloat(paramVal, 64)
		case "player_decay":
			playerType.PlayerDecay, err = strconv.ParseFloat(paramVal, 64)
		case "inertia_moment":
			playerType.InertiaMoment, err = strconv.ParseFloat(paramVal, 64)
		case "dash_power_rate":
			playerType.DashPowerRate, err = strconv.ParseFloat(paramVal, 64)
		case "player_size":
			playerType.PlayerSize, err = strconv.ParseFloat(paramVal, 64)
		case "kickable_margin":
			playerType.KickableMargin, err = strconv.ParseFloat(paramVal, 64)
		case "kick_rand":
			playerType.KickRand, err = strconv.ParseFloat(paramVal, 64)
		case "extra_stamina":
			playerType.ExtraStamina, err = strconv.ParseFloat(paramVal, 64)
		case "effort_max":
			playerType.EffortMax, err = strconv.ParseFloat(paramVal, 64)
		case "effort_min":
			playerType.EffortMin, err = strconv.ParseFloat(paramVal, 64)
		case "kick_power_rate":
			playerType.KickPowerRate, err = strconv.ParseFloat(paramVal, 64)
		case "foul_detect_probability":
			playerType.FoulDetectProbability, err = strconv.ParseFloat(paramVal, 64)
		case "catchable_area_l_stretch":
			playerType.CatchableAreaLStretch, err = strconv.ParseFloat(paramVal, 64)
		default:
			errCh <- fmt.Sprintf("unsupported player_type param (%s)", param)
		}
		if err != nil {
			errCh <- fmt.Sprintf("could not parse player type param (%s): %s", param, err)
			err = nil
		}
	}

	return playerType
}
