package rcsscommon

// SideType is an alias for defining the field side where team is playing
type SideType rune

const (
	// LeftSide ...
	LeftSide SideType = 'l'
	// RightSide ...
	RightSide SideType = 'r'
)

func (s SideType) String() string {
	if s == LeftSide {
		return "left"
	} else if s == RightSide {
		return "right"
	}
	return ""
}
