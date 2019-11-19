package body

// Quality specifies the quality level of the current view mode
type Quality byte

const (
	High Quality = iota + 0

	Low
)

// Width specifies the width of the current view mode
type Width byte

const (
	Narrow Width = iota + 0

	Normal

	Wide
)
