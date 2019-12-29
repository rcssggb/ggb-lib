package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rcssggb/ggb-lib/rcsscommon"
)

// FlagArray is an array of FlagData
type FlagArray []FlagData

// FlagData contains data about the field markers
type FlagData struct {
	ID        rcsscommon.FlagID
	Distance  float64
	Direction float64
}

func parseFlag(flagID rcsscommon.FlagID, data []string) (flag FlagData, err error) {
	if len(data) < 2 {
		err = fmt.Errorf("missing flag data, got: %s", strings.Join(data, " "))
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
	flag = FlagData{
		ID:        flagID,
		Distance:  dist,
		Direction: dir,
	}
	return
}

// Len is the number of elements in the FlagArray.
func (f FlagArray) Len() int {
	return len(f)
}

// Swap swaps the elements with indexes i and j.
func (f FlagArray) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

// Less reports whether the element with index i should sort before the element with index j.
func (f FlagArray) Less(i, j int) bool {
	return f[i].Distance < f[j].Distance
}
