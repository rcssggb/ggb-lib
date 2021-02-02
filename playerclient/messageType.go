package playerclient

// messageType is the type for message type constants
type messageType byte

const (
	unsupportedMsg messageType = iota + 0
	errorMsg
	initMsg
	sightMsg
	bodyMsg
	hearMsg
	playerTypeMsg
	playerParamMsg
	serverParamMsg
	changePlayerTypeMsg
	thinkMsg
	genericOkMsg
)
