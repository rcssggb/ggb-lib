package rcsscommon

// SideType is an alias for defining the field side where team is playing
type SideType rune

const (
	// LeftSide ...
	LeftSide SideType = 'l'
	// RightSide ...
	RightSide SideType = 'l'
)

func (s SideType) String() string {
	if s == 'r' {
		return "right"
	}
	return "left"
}
