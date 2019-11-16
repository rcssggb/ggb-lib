package parser

import (
	"fmt"
	"strconv"
	"strings"
)

// BallData contains visual information about the ball position and velocity
type BallData struct {
	Distance   float64
	Direction  float64
	DistChange float64
	DirChange  float64
}

func parseBall(data []string) (ball BallData, err error) {
	var distStr, dirStr, distChngStr, dirChngStr string
	var dist, dir, distChng, dirChng float64
	if len(data) < 2 {
		err = fmt.Errorf("missing ball data, got: %s", strings.Join(data, " "))
		return
	}

	distStr = data[0]
	dist, parseErr := strconv.ParseFloat(distStr, 64)
	if parseErr != nil {
		err = fmt.Errorf("unable to parse float \"%s\": %s", distStr, parseErr)
		return
	}

	dirStr = data[1]
	dir, parseErr = strconv.ParseFloat(dirStr, 64)
	if parseErr != nil {
		err = fmt.Errorf("unable to parse float \"%s\": %s", dirStr, parseErr)
		return
	}

	if len(data) >= 4 {
		distChngStr = data[2]
		distChng, parseErr = strconv.ParseFloat(distChngStr, 64)
		if parseErr != nil {
			err = fmt.Errorf("unable to parse float \"%s\": %s", distStr, parseErr)
			return
		}

		dirChngStr = data[3]
		dirChng, parseErr = strconv.ParseFloat(dirChngStr, 64)
		if parseErr != nil {
			err = fmt.Errorf("unable to parse float \"%s\": %s", dirStr, parseErr)
			return
		}
	}

	ball = BallData{
		Distance:   dist,
		Direction:  dir,
		DistChange: distChng,
		DirChange:  dirChng,
	}

	return
}
