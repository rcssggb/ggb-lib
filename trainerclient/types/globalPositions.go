package types

type GlobalPositionsSymbols struct {
	Time    int
	Objects map[string][]string
}

type GlobalPositions struct {
	Time      int
	Ball      AbsPosition
	GoalLeft  AbsPosition
	GoalRight AbsPosition
	Teams     map[string]Team
}

type Team map[int]AbsPosition

// AbsPosition defines the generic absolute position coordinates definition
type AbsPosition struct {
	X           float64
	Y           float64
	DeltaX      float64
	DeltaY      float64
	BodyAngle   float64
	NeckAngle   float64
	IsPointing  bool
	PointingDir float64
	Action      string
}
