package trainerclient

import (
	"fmt"

	"github.com/rcssggb/ggb-lib/rcsscommon"
)

// decode continuously receives messages from recvChannel and calls lexer and parser to structure the message data
func (c *Client) decode() {
	var m message

	for {
		m = <-c.recvChannel
		switch m.Type() {
		case lookMsg:
			// TODO
			c.Log(m.data)
		case serverParamMsg:
			c.serverParams.Parse(m.data, c.errChannel)
		case playerTypeMsg:
			pType := rcsscommon.ParsePlayerType(m.data, c.errChannel)
			if pType.ID != -1 {
				c.playerTypes[pType.ID] = pType
			}
		case initMsg:
			continue
		case errorMsg:
			c.errChannel <- m.String()
		case eyeMsg:
			if m.data == "(ok eye on)\x00" {
				c.eyeMode = true
			} else if m.data == "(ok eye off)\x00" {
				c.eyeMode = false
			} else {
				c.errChannel <- fmt.Sprintf("unsupported eye response: %s", m)
			}
		case unsupportedMsg:
			c.errChannel <- fmt.Sprintf("unsupported message received from server: %s", m)
			continue
		}
	}
}
