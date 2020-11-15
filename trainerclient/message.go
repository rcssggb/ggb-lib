package trainerclient

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
	case strings.HasPrefix(m.data, "(think"):
		mType = thinkMsg
	case strings.HasPrefix(m.data, "(player_type "):
		mType = playerTypeMsg
	case strings.HasPrefix(m.data, "(server_param "):
		mType = serverParamMsg
	case strings.HasPrefix(m.data, "(player_param "):
		mType = playerParamMsg
	case strings.HasPrefix(m.data, "(init ok)"):
		mType = initMsg
	case strings.HasPrefix(m.data, "(start ok)"):
		mType = startMsg
	case strings.HasPrefix(m.data, "(recover ok)"):
		mType = recoverMsg
	case strings.HasPrefix(m.data, "(move ok)"):
		mType = moveMsg
	case strings.HasPrefix(m.data, "(warning "):
		fallthrough
	case strings.HasPrefix(m.data, "(error "):
		mType = errorMsg
	case strings.HasPrefix(m.data, "(warning "):
		mType = warningMsg
	case strings.HasPrefix(m.data, "(ok look"):
		fallthrough
	case strings.HasPrefix(m.data, "(see_global"):
		mType = lookMsg
	case strings.HasPrefix(m.data, "(ok eye"):
		mType = eyeMsg
	case strings.HasPrefix(m.data, "(ok ear"):
		mType = earMsg
	case strings.HasPrefix(m.data, "(ok team_names"):
		mType = teamNamesMsg
	case strings.HasPrefix(m.data, "(ok check_ball"):
		mType = checkBallMsg
	case strings.HasPrefix(m.data, "(ok change_player_type"):
		mType = changePlayerTypeMsg
	case strings.HasPrefix(m.data, "(ok move"):
		fallthrough
	case strings.HasPrefix(m.data, "(ok start"):
		mType = genericOkMsg
	default:
		mType = unsupportedMsg
	}
	return mType
}
