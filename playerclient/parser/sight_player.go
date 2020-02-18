package parser

import (
	"fmt"
	"log"
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
}

func parsePlayers(playersSymbols lexer.SightPlayersSymbols) (players PlayerArray) {
	// Add known players to player array
	for pName, pData := range playersSymbols.Known {
		pNameParts := strings.Split(pName, " ")
		if len(pNameParts) != 3 {
			log.Println("internal warning: known player name should be of form (p <team_name> <unum>)")
			continue
		}

		teamName := strings.ReplaceAll(pNameParts[1], "\"", "")

		unum, parseErr := strconv.ParseInt(pNameParts[2], 10, 64)
		if parseErr != nil {
			log.Printf("warning: unable to parse unum integer \"%s\": %s\n", pNameParts[2], parseErr)
			continue
		}

		if len(pData) < 2 {
			log.Printf("warning: missing seen player data, got: ((%s)) %s)", pName, strings.Join(pData, " "))
			continue
		}

		dist, dir, distChange, dirChange, err := parsePlayerVals(pData)
		if err != nil {
			log.Println(err)
			continue
		}

		players = append(players, PlayerData{
			Team:       teamName,
			Unum:       int(unum),
			Distance:   dist,
			Direction:  dir,
			DistChange: distChange,
			DirChange:  dirChange,
		})
	}

	// Add known team players to player array
	for teamName, pList := range playersSymbols.KnownTeam {
		teamName = strings.ReplaceAll(teamName, "\"", "")
		for _, pData := range pList {
			dist, dir, distChange, dirChange, err := parsePlayerVals(pData)
			if err != nil {
				log.Println(err)
				continue
			}

			players = append(players, PlayerData{
				Team:       teamName,
				Distance:   dist,
				Direction:  dir,
				DistChange: distChange,
				DirChange:  dirChange,
			})
		}
	}

	// Add unknown players to player array
	for _, pData := range playersSymbols.Unknown {
		dist, dir, distChange, dirChange, err := parsePlayerVals(pData)
		if err != nil {
			log.Println(err)
			continue
		}

		players = append(players, PlayerData{
			Distance:   dist,
			Direction:  dir,
			DistChange: distChange,
			DirChange:  dirChange,
		})
	}

	return
}

func parsePlayerVals(data []string) (dist, dir, distChange, dirChange float64, err error) {
	distStr := data[0]
	dist, err = strconv.ParseFloat(distStr, 64)
	if err != nil {
		err = fmt.Errorf("warning: unable to parse float \"%s\": %s", distStr, err)
		return
	}

	dirStr := data[1]
	dir, err = strconv.ParseFloat(dirStr, 64)
	if err != nil {
		err = fmt.Errorf("warning: unable to parse float \"%s\": %s", dirStr, err)
		return
	}

	if len(data) >= 4 {
		distChngStr := data[2]
		distChange, err = strconv.ParseFloat(distChngStr, 64)
		if err != nil {
			err = fmt.Errorf("warning: unable to parse float \"%s\": %s", distChngStr, err)
			return
		}

		dirChngStr := data[3]
		dirChange, err = strconv.ParseFloat(dirChngStr, 64)
		if err != nil {
			err = fmt.Errorf("warning: unable to parse float \"%s\": %s", dirChngStr, err)
			return
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
