package parser

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/rcssggb/ggb-lib/playerclient/lexer"
)

// PlayerArray is an array of PlayerData
type PlayerArray []PlayerData

// PlayerData contains data about seen players.
// Team == "" means you don't know the player's team.
// Unum == -1 means you don't know the player's shirt num.
type PlayerData struct {
	Team       string
	Unum       int
	Distance   float64
	Direction  float64
	DistChange float64
	DirChange  float64
	BodyDir    float64
	HeadDir    float64
}

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

		dist, dir, distChange, dirChange, bodyDir, headDir, err := parsePlayerVals(pData)
		if err != nil {
			errCh <- fmt.Sprintf("unable to parse sight data for seen player %s: %s", pName, err)
			continue
		}

		players = append(players, PlayerData{
			Team:       teamName,
			Unum:       int(unum),
			Distance:   dist,
			Direction:  dir,
			DistChange: distChange,
			DirChange:  dirChange,
			BodyDir:    bodyDir,
			HeadDir:    headDir,
		})
	}

	// Add known team players to player array
	for teamName, pList := range playersSymbols.KnownTeam {
		teamName = strings.ReplaceAll(teamName, "\"", "")
		for _, pData := range pList {
			dist, dir, distChange, dirChange, bodyDir, headDir, err := parsePlayerVals(pData)
			if err != nil {
				errCh <- fmt.Sprintf("unable to parse sight data for seen player of team %s: %s", teamName, err)
				continue
			}

			players = append(players, PlayerData{
				Team:       teamName,
				Distance:   dist,
				Direction:  dir,
				DistChange: distChange,
				DirChange:  dirChange,
				BodyDir:    bodyDir,
				HeadDir:    headDir,
			})
		}
	}

	// Add unknown players to player array
	for _, pData := range playersSymbols.Unknown {
		dist, dir, distChange, dirChange, bodyDir, headDir, err := parsePlayerVals(pData)
		if err != nil {
			errCh <- fmt.Sprintf("unable to parse sight data for seen player: %s", err)
			continue
		}

		players = append(players, PlayerData{
			Distance:   dist,
			Direction:  dir,
			DistChange: distChange,
			DirChange:  dirChange,
			BodyDir:    bodyDir,
			HeadDir:    headDir,
		})
	}

	return
}

func parsePlayerVals(data []string) (dist, dir, distChange, dirChange, bodyDir, headDir float64, err error) {
	if len(data) < 2 {
		err = fmt.Errorf("too few values in player data, got: %s", strings.Join(data, " "))
		return
	}

	distStr := data[0]
	dist, err = strconv.ParseFloat(distStr, 64)
	if err != nil {
		err = fmt.Errorf("unable to parse float \"%s\": %s", distStr, err)
		return
	}

	dirStr := data[1]
	dir, err = strconv.ParseFloat(dirStr, 64)
	if err != nil {
		err = fmt.Errorf("unable to parse float \"%s\": %s", dirStr, err)
		return
	}

	// If player polar velocities were measured
	if len(data) >= 4 {
		distChngStr := data[2]
		distChange, err = strconv.ParseFloat(distChngStr, 64)
		if err != nil {
			err = fmt.Errorf("unable to parse float \"%s\": %s", distChngStr, err)
			return
		}

		dirChngStr := data[3]
		dirChange, err = strconv.ParseFloat(dirChngStr, 64)
		if err != nil {
			err = fmt.Errorf("unable to parse float \"%s\": %s", dirChngStr, err)
			return
		}

		// If player facing was measured, parse it
		if len(data) >= 6 {
			bodyDirStr := data[4]
			bodyDir, err = strconv.ParseFloat(bodyDirStr, 64)
			if err != nil {
				err = fmt.Errorf("unable to parse float \"%s\": %s", bodyDirStr, err)
				return
			}

			headDirStr := data[5]
			headDir, err = strconv.ParseFloat(headDirStr, 64)
			if err != nil {
				err = fmt.Errorf("unable to parse float \"%s\": %s", headDirStr, err)
				return
			}
		} else {
			// If player facing was not measured, assume he is facing towards it's velocity
			dirChngRad := 3.14159 * dirChange / 180.0
			tanVel := dist * dirChngRad
			bodyDir = 180.0 * math.Atan2(tanVel, distChange) / 3.14159
			headDir = bodyDir
		}
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
