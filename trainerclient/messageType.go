package trainerclient

// messageType is the type for message type constants
type messageType byte

const (
	unsupportedMsg messageType = iota + 0
	initMsg
	serverParamMsg
	playerParamMsg
	playerTypeMsg
	lookMsg
	eyeMsg
	earMsg
	teamNamesMsg
	checkBallMsg
	changePlayerTypeMsg
	moveMsg
	startMsg
	genericOkMsg
	recoverMsg
	errorMsg
	warningMsg
	thinkMsg
)
