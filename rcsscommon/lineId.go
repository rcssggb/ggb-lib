package rcsscommon

// LineID is the type representing each line in the field
type LineID byte

const (
	// LineTop y = -34.00
	LineTop LineID = iota + 0

	// LineRight x = 52.50
	LineRight

	// LineBottom y = 34.00
	LineBottom

	// LineLeft x = -52.50
	LineLeft
)
