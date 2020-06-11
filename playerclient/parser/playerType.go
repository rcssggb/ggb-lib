package parser

// PlayerTypeData is the individual player information
type PlayerTypeData struct {
	ID             int
	PlayerSpeedMax float64
	StaminaIncMax  float64
	PlayerDecay    float64
	InertiaMoment  float64
	DashPowerRate  float64
	PlayerSize     float64
	KickableMargin float64
	KickRand       float64
	ExtraStamina   float64
	EffortMax      float64
	EffortMin      float64
	KickPowerRate  float64
	FoulDetectProb float64
	CatchableArea  float64
}

// PlayerType parses player type data coming from lexer
func PlayerType(symbols map[string]string) (*PlayerTypeData, error) {
	// TODO
	return nil, nil
}
