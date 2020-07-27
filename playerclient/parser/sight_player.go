package parser

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/rcssggb/ggb-lib/playerclient/lexer"
	"github.com/rcssggb/ggb-lib/playerclient/types"
)

// PlayerArray is an array of types.PlayerPosition
type PlayerArray []types.PlayerPosition

func parsePlayers(playersSymbols lexer.SightPlayersSymbols, errCh chan string) (players PlayerArray) {
	// Add known players to player array
	for pName, pData := range playersSymbols.Known {
		pNameParts := strings.Split(pName, " ")
		if len(pNameParts) != 3 {
			errCh <- fmt.Sprintf("invalid known player name (should be of form `p <team_name> <unum>`)")
			continue
		}

		teamName := strings.ReplaceAll(pNameParts[1], "\"", "")

		unum, parseErr := strconv.ParseInt(pNameParts[2], 10, 64)
		if parseErr != nil {
			errCh <- fmt.Sprintf("unable to parse unum integer for player %s: %s\n", pName, parseErr)
			continue
		}

		ppos, err := parsePlayerVals(pData)
		if err != nil {
			errCh <- fmt.Sprintf("unable to parse sight data for seen player %s: %s", pName, err)
			continue
		}

		ppos.Team = teamName
		ppos.Unum = int(unum)

		players = append(players, ppos)
	}

	// Add known team players to player array
	for teamName, pList := range playersSymbols.KnownTeam {
		teamName = strings.ReplaceAll(teamName, "\"", "")
		for _, pData := range pList {
			ppos, err := parsePlayerVals(pData)
			if err != nil {
				errCh <- fmt.Sprintf("unable to parse sight data for seen player of team %s: %s", teamName, err)
				continue
			}

			ppos.Team = teamName

			players = append(players, ppos)
		}
	}

	// Add unknown players to player array
	for _, pData := range playersSymbols.Unknown {
		ppos, err := parsePlayerVals(pData)
		if err != nil {
			errCh <- fmt.Sprintf("unable to parse sight data for seen player: %s", err)
			continue
		}

		players = append(players, ppos)
	}

	return
}

func parsePlayerVals(vals []string) (ppos types.PlayerPosition, err error) {
	data := make([]float64, 7)
	isPointing := false
	var actionData string
	for idx, val := range vals {
		if val != "k" && val != "t" {
			data[idx], err = strconv.ParseFloat(val, 64)
			if err != nil {
				break
			}
			if idx == 6 {
				isPointing = true // it's better to use a "p" marker in action if pointing and kicking at the same time is not allowed
			}
		} else {
			actionData = val
		}
	}
	if err != nil {
		err = fmt.Errorf("error parsing float64 val: %s", err)
		return
	}

	ppos = types.PlayerPosition{
		Distance:    data[0],
		Direction:   data[1],
		DistChange:  data[2],
		DirChange:   data[3],
		BodyDir:     data[4],
		NeckDir:     data[5],
		IsPointing:  isPointing,
		PointingDir: data[6],
		Action:      actionData,
	}

	return
}

// Len is the number of elements in the PlayerArray.
func (f PlayerArray) Len() int {
	return len(f)
}

// Swap swaps the elements with indexes i and j.
func (f PlayerArray) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

// Less reports whether the element with index i should sort before the element with index j.
func (f PlayerArray) Less(i, j int) bool {
	if f[i].Distance == f[j].Distance {
		if math.Abs(f[i].Direction) == math.Abs(f[j].Direction) {
			return f[i].Direction < f[j].Direction
		}
		return math.Abs(f[i].Direction) < math.Abs(f[j].Direction)
	}
	return f[i].Distance < f[j].Distance
}
