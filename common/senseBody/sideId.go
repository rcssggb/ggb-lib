package body

// SideID is the type representing the side the player is focusing on
type SideID byte

const (
	Left SideID = iota + 0

	Neutral

	Right
)
