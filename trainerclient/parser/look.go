package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rcssggb/ggb-lib/trainerclient/lexer"
)

type GlobalPositions struct {
	Time  int
	Ball  AbsPosition
	lGoal AbsPosition
	rGoal AbsPosition
	Teams map[string]Team
}

type Team map[int]AbsPosition

// AbsPosition defines the generic absolute position coordinates definition
// TODO: maybe this should be in common and merged with the
// playerclient parser.SightData definition
type AbsPosition struct {
	X         float64
	Y         float64
	DeltaX    float64
	DeltaY    float64
	BodyAngle float64
	NeckAngle float64
}

// Look parses global positions info coming from lexer
func Look(gpSymbols lexer.GlobalPositions, errCh chan string) *GlobalPositions {
	gp := GlobalPositions{
		Time: gpSymbols.Time,
	}

	gp.Teams = make(map[string]Team)

	for objName, objData := range gpSymbols.Objects {
		var err error
		data := make([]float64, len(objData))
		for idx, val := range objData {
			data[idx], err = strconv.ParseFloat(val, 64)
			if err != nil {
				break
			}
		}
		if err != nil {
			errCh <- fmt.Sprintf("error parsing \"%s\" global position: %s", objName, err)
			continue
		}

		// If object is a goal
		if strings.HasPrefix(objName, "g ") {
			if len(objData) != 2 {
				errCh <- fmt.Sprintf("goal object \"%s\" has invalid number of data values (%d)", objName, len(objData))
				continue
			}

			if objName == "g r" {
				gp.rGoal.X = data[0]
				gp.rGoal.Y = data[1]
			} else if objName == "g l" {
				gp.lGoal.X = data[0]
				gp.lGoal.Y = data[1]
			} else {
				errCh <- fmt.Sprintf("invalid goal object \"%s\"", objName)
			}
		}

		// If object is the ball
		if objName == "b" {
			if len(objData) != 4 {
				errCh <- fmt.Sprintf("ball object \"%s\" has invalid number of data values (%d)", objName, len(objData))
				continue
			}

			gp.Ball.X = data[0]
			gp.Ball.Y = data[1]
			gp.Ball.DeltaX = data[2]
			gp.Ball.DeltaY = data[3]
		}

		// if object is a player
		if strings.HasPrefix(objName, "p ") {
			if len(objData) != 6 {
				errCh <- fmt.Sprintf("player object \"%s\" has invalid number of data values (%d)", objName, len(objData))
				continue
			}

			// Splits "p \"team-name\" 1" into ["p" "\"team-name\"" "1"]
			player := strings.Split(objName, " ")

			teamName := player[1]
			teamName = strings.TrimPrefix(teamName, "\"")
			teamName = strings.TrimSuffix(teamName, "\"")

			unum := strconv.Atoi(player[2])

			if gp.Teams[teamName] == nil {
				gp.Teams[teamName] = make(Team)
			}

			gp.Teams[teamName][unum] = AbsPosition{
				X:         data[0],
				Y:         data[1],
				DeltaX:    data[2],
				DeltaY:    data[3],
				BodyAngle: data[4],
				NeckAngle: data[5],
			}
		}

	}

	return &gp
}
