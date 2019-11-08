package parser

import "../lexer"

// SightData is the final format of everything that was seen by the player
type SightData struct {
	Time      int
	FlagArray []FlagData
}

// FlagData contains data about the field markers
type FlagData struct {
	ID        FlagID
	Distance  float64
	Direction float64
}

// Sight parses sight data coming from lexer
func Sight(symbols lexer.SightSymbols) (*SightData, error) {

	return nil, nil
}
