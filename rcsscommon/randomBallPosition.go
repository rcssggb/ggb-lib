package rcsscommon

import (
	"math"
	"math/rand"
	"time"
)

// RandomBallPosition randomizes a valid ball position
func RandomBallPosition() BallPosition {
	rand.Seed(time.Now().UnixNano())

	x := rand.Float64() * FieldMaxX
	if rand.Intn(2) != 0 {
		x = -x
	}
	y := rand.Float64() * FieldMaxY
	if rand.Intn(2) != 0 {
		y = -y
	}
	speed := rand.Float64() * 3 // default ball max speed
	dir := rand.Float64() * 360

	return BallPosition{
		X:      x,
		Y:      y,
		DeltaX: speed * math.Sin(dir),
		DeltaY: speed * math.Cos(dir),
	}
}
