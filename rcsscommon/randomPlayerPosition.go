package rcsscommon

import (
	"math/rand"
	"time"
)

// RandomPosition randomizes a valid ball position
func RandomPosition() (x, y float64) {
	rand.Seed(time.Now().UnixNano())

	x = rand.Float64() * FieldMaxX
	if rand.Intn(2) != 0 {
		x = -x
	}
	y = rand.Float64() * FieldMaxY
	if rand.Intn(2) != 0 {
		y = -y
	}

	return
}
