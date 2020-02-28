package trainerclient

import (
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
	default:
		mType = unsupportedMsg
	}
	return mType
}
