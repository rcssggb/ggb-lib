package playerclient

import (
	"strings"
	"time"
)

// message is the struct used to move and parse RCSS commands
type message struct {
	timestamp time.Time
	data      string
}

// String implements the native Stringer interface
func (m message) String() string {
	return string(m.data)
}

// Type parses and returns the MessageType for the message
func (m message) Type() (mType messageType) {
	switch {
	case strings.HasPrefix(m.data, "(see "):
		mType = sightMsg
	case strings.HasPrefix(m.data, "(sense_body "):
		mType = bodyMsg
	case strings.HasPrefix(m.data, "(hear "):
		mType = hearMsg
	case strings.HasPrefix(m.data, "(warning "):
		fallthrough
	case strings.HasPrefix(m.data, "(error "):
		mType = errorMsg
	case strings.HasPrefix(m.data, "(init "):
		mType = initMsg
	case strings.HasPrefix(m.data, "(player_type "):
		mType = playerTypeMsg
	case strings.HasPrefix(m.data, "(server_param "):
		mType = serverParamMsg
	case strings.HasPrefix(m.data, "(player_param "):
		mType = playerParamMsg
	// Unsupported msgs below
	case strings.HasPrefix(m.data, "(score "):
		mType = unsupportedMsg
	case strings.HasPrefix(m.data, "(ok "):
		mType = unsupportedMsg
	case strings.HasPrefix(m.data, "(change_player_type "):
		mType = unsupportedMsg
	case strings.HasPrefix(m.data, "(fullstate "):
		mType = unsupportedMsg
	case strings.HasPrefix(m.data, "(change "):
		mType = unsupportedMsg
	case strings.HasPrefix(m.data, "(clang "):
		mType = unsupportedMsg
	default:
		mType = unsupportedMsg
	}
	return mType
}
