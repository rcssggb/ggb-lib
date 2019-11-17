package body

// CardID is the card the player received
type CardID byte

const (
	NoCard CardID = iota + 0

	Yellow

	Red
)
