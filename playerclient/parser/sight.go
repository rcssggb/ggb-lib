package parser

import (
	"fmt"
	"sort"

	"github.com/rcssggb/ggb-lib/playerclient/lexer"
)

// SightData is the final format of everything that was seen by the player
type SightData struct {
	Time    int
	Flags   FlagArray
	Players PlayerArray
	Ball    *BallData
}

// Sight parses sight data coming from lexer
func Sight(symbols lexer.SightSymbols, errCh chan string) (sightData *SightData) {
	sightData = &SightData{
		Time: symbols.Time,
	}

	// Parse flag markers, ball and goal object
	for objName, data := range symbols.ObjMap {

		// Make sure objName is not empty
		if len(objName) == 0 {
			errCh <- fmt.Sprintf("empty object name in sight symbols map")
			return
		}

		objType := objName[0]

		// Test if object is ball
		if objType == 'b' {
			ball, err := parseBall(data)
			if err != nil {
				errCh <- fmt.Sprintf("error parsing ball sight data: %s", err)
				continue
			}
			sightData.Ball = ball
		}

		// Test if object is flag
		if objType == 'f' || objType == 'g' {
			flagID, flagDefined := flagMap[objName]
			if !flagDefined {
				errCh <- fmt.Sprintf("seen unknown flag marker (%s)", objName)
				continue
			}
			flagData, err := parseFlag(flagID, data)
			if err != nil {
				errCh <- fmt.Sprintf("unable to parse flag sight data: %s", err)
				continue
			}
			sightData.Flags = append(sightData.Flags, flagData)
		}
	}

	sort.Sort(sightData.Flags)

	// Parse players
	sightData.Players = parsePlayers(symbols.Players, errCh)
	sort.Sort(sightData.Players)

	return
}
