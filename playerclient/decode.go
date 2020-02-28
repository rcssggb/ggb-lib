package playerclient

import (
	"fmt"

	"github.com/rcssggb/ggb-lib/playerclient/lexer"
	"github.com/rcssggb/ggb-lib/playerclient/parser"
)

// decode continuously receives messages from recvChannel and calls lexer and parser to structure the message data
func (c *Client) decode() {
	var m message

	for {
		m = <-c.recvChannel
		switch m.Type() {
		case initMsg:
			initData, err := parser.Init(m.data)
			if err != nil {
				c.errChannel <- err.Error()
				continue
			}
			c.shirtNum = initData.Unum
			c.teamSide = initData.Side
			c.playMode = initData.PlayMode
		case sightMsg:
			sightSymbols, err := lexer.Sight(m.data)
			if err != nil {
				c.errChannel <- err.Error()
				continue
			}

			sightData, err := parser.Sight(*sightSymbols)
			if err != nil {
				c.errChannel <- err.Error()
				continue
			}

			if sightData.Time >= c.currentTime {
				c.sightData = *sightData
				c.currentTime = sightData.Time
			}
		case bodyMsg:
			bodySymbols, err := lexer.SenseBody(m.data)
			if err != nil {
				c.errChannel <- err.Error()
				continue
			}

			bodyData, err := parser.SenseBody(*bodySymbols)
			if err != nil {
				c.errChannel <- err.Error()
				continue
			}

			if bodyData.Time >= c.currentTime {
				c.bodyData = *bodyData
				c.currentTime = bodyData.Time
			}
		case hearMsg:
			hearSyms, err := lexer.Hear(m.data)
			if err != nil {
				c.errChannel <- err.Error()
			}

			if hearSyms.Time >= c.currentTime {
				if hearSyms.Sender == "referee" {
					c.playMode = hearSyms.Message
				}
			}

		// case serverParamMsg:
		// _, err := lexer.ServerParam(m.data)
		case errorMsg:
			c.errChannel <- m.String()
		case unsupportedMsg:
			c.errChannel <- fmt.Sprintf("unsupported message received from server: %s", m.String())
		}
	}
}
