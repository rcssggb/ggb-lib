package trainerclient

import (
	"fmt"
)

// decode continuously receives messages from recvChannel and calls lexer and parser to structure the message data
func (c *Client) decode() {
	var m message

	for {
		m = <-c.recvChannel
		switch m.Type() {
		case serverParamMsg:
			c.serverParams.Parse(m.data, c.errChannel)
		case initMsg:
			continue
		case errorMsg:
			c.errChannel <- m.String()
		case unsupportedMsg:
			c.errChannel <- fmt.Sprintf("unsupported message received from server: %s", m.String())
			continue
		}
	}
}
