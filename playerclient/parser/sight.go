package parser

import (
	"fmt"
	"sort"

	"github.com/rcssggb/ggb-lib/playerclient/lexer"
)

// SightData is the final format of everything that was seen by the player
type SightData struct {
	Time  int
	Flags FlagArray
	Ball  *BallData
}

// Sight parses sight data coming from lexer
func Sight(symbols lexer.SightSymbols) (sightData *SightData, err error) {
	sightData = &SightData{
		Time: symbols.Time,
	}

	for objName, data := range symbols.SightMap {

		// Test if object is ball
		if objName == "b" {
			ball, ballErr := parseBall(data)
			if ballErr != nil {
				err = fmt.Errorf("error parsing ball: %s", err)
				return
			}
			sightData.Ball = ball
		}

		// Test if object is flag
		flagID, isFlag := flagMap[objName]
		if isFlag {
			flagData, flagErr := parseFlag(flagID, data)
			if flagErr != nil {
				err = fmt.Errorf("error parsing flag: %s", err)
				return
			}
			sightData.Flags = append(sightData.Flags, flagData)
			continue
		}
	}

	sort.Sort(sightData.Flags)

	return
}
