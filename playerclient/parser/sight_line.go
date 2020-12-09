package parser

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/rcssggb/ggb-lib/rcsscommon"
)

// LineArray is an array of LineData
type LineArray []LineData

// LineData contains data about the field lines
type LineData struct {
	ID        rcsscommon.LineID
	Distance  float64
	Direction float64
}

func parseLine(lineID rcsscommon.LineID, data []string) (line LineData, err error) {
	if len(data) < 2 {
		err = fmt.Errorf("missing line data, got: %s", strings.Join(data, " "))
		return
	}
	distStr := data[0]
	dirStr := data[1]
	dist, parseErr := strconv.ParseFloat(distStr, 64)
	if parseErr != nil {
		err = fmt.Errorf("unable to parse float \"%s\": %s", distStr, parseErr)
		return
	}
	dir, parseErr := strconv.ParseFloat(dirStr, 64)
	if parseErr != nil {
		err = fmt.Errorf("unable to parse float \"%s\": %s", dirStr, parseErr)
		return
	}
	line = LineData{
		ID:        lineID,
		Distance:  dist,
		Direction: dir,
	}
	return
}

// Len is the number of elements in the LineArray.
func (l LineArray) Len() int {
	return len(l)
}

// Swap swaps the elements with indexes i and j.
func (l LineArray) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

// Less reports whether the element with index i should sort before the element with index j.
func (l LineArray) Less(i, j int) bool {
	if l[i].Distance == l[j].Distance {
		if math.Abs(l[i].Direction) == math.Abs(l[j].Direction) {
			return l[i].Direction < l[j].Direction
		}
		return math.Abs(l[i].Direction) < math.Abs(l[j].Direction)
	}
	return l[i].Distance < l[j].Distance
}
