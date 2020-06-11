package trainerclient

// messageType is the type for message type constants
type messageType byte

const (
	unsupportedMsg messageType = iota + 0
	initMsg
	serverParamMsg
	playerTypeMsg
	lookMsg
	eyeMsg
	errorMsg
)
