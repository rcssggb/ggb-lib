package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rcssggb/ggb-lib/trainerclient/types"
)

// Look parses global positions info coming from lexer
func Look(gpSymbols types.GlobalPositionsSymbols, errCh chan string) *types.GlobalPositions {
	gp := types.GlobalPositions{
		Time: gpSymbols.Time,
	}

	gp.Teams = make(map[string]types.Team)

	for objName, objData := range gpSymbols.Objects {
		var err error

		// length is 7 because it seems a "pointing direction" or "action" is sometimes provided although the docs are not provided
		data := make([]float64, 7)
		isPointing := false
		var actionData string
		for idx, val := range objData {
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
			errCh <- fmt.Sprintf("error parsing \"%s\" global position: %s", objName, err)
			continue
		}

		objPos := types.AbsPosition{
			X:           data[0],
			Y:           data[1],
			DeltaX:      data[2],
			DeltaY:      data[3],
			BodyAngle:   data[4],
			NeckAngle:   data[5],
			IsPointing:  isPointing,
			PointingDir: data[6],
			Action:      actionData,
		}

		// If object is a goal
		if strings.HasPrefix(objName, "g ") {
			if len(objData) != 2 {
				errCh <- fmt.Sprintf("goal object \"%s\" has invalid number of data values (%d)", objName, len(objData))
				continue
			}

			if objName == "g r" {
				gp.GoalRight = objPos
			} else if objName == "g l" {
				gp.GoalLeft = objPos
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

			gp.Ball = objPos
		}

		// if object is a player
		if strings.HasPrefix(objName, "p ") {
			if len(objData) != 6 && len(objData) != 7 {
				errCh <- fmt.Sprintf("player object \"%s\" has invalid number of data values (%d)", objName, len(objData))
				continue
			}

			// Splits "p \"team-name\" 1" into ["p" "\"team-name\"" "1"]
			player := strings.Split(objName, " ")

			teamName := player[1]
			teamName = strings.TrimPrefix(teamName, "\"")
			teamName = strings.TrimSuffix(teamName, "\"")

			unum, err := strconv.Atoi(player[2])
			if err != nil {
				errCh <- fmt.Sprintf("failed to parse '%s' unum '%s'", objName, player[2])
			}

			if gp.Teams[teamName] == nil {
				gp.Teams[teamName] = make(types.Team)
			}

			gp.Teams[teamName][unum] = objPos

		}

	}

	return &gp
}
