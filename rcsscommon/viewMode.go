package rcsscommon

// ViewQuality specifies the quality level of the current view mode
type ViewQuality float64

const (
	ViewQualityHigh ViewQuality = 1

	ViewQualityLow ViewQuality = 0.5
)

// ViewWidth specifies the width of the current view mode
type ViewWidth float64

const (
	ViewWidthNarrow ViewWidth = 0.5

	ViewWidthNormal ViewWidth = 1

	ViewWidthWide ViewWidth = 2
)
