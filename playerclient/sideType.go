package playerclient

// sideType is an alias for defining the field side where team is playing
type sideType bool

const (
	// LeftSide ...
	LeftSide sideType = false
	// RightSide ...
	RightSide sideType = true
)

func (s sideType) String() string {
	if s {
		return "right"
	}
	return "left"
}
