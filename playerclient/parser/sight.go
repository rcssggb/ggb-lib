package parser

import (
	"fmt"
	"log"
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
		Time:  symbols.Time,
		Flags: FlagArray{},
	}

	for objName, data := range symbols.ObjMap {

		// Make sure objName is not empty
		if len(objName) == 0 {
			err = fmt.Errorf("empty object name")
			return
		}

		objType := objName[0]

		// Test if object is ball
		if objType == 'b' {
			ball, ballErr := parseBall(data)
			if ballErr != nil {
				err = fmt.Errorf("error parsing ball: %s", err)
				return
			}
			sightData.Ball = ball
		}

		// Test if object is flag
		if objType == 'f' || objType == 'g' {
			flagID, flagDefined := flagMap[objName]
			if !flagDefined {
				log.Printf("seen unknown flag marker (%s)\n", objName)
				continue
			}
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
