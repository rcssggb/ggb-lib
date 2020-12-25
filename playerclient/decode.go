package playerclient

import (
	"fmt"

	"github.com/rcssggb/ggb-lib/playerclient/lexer"
	"github.com/rcssggb/ggb-lib/playerclient/parser"
	"github.com/rcssggb/ggb-lib/rcsscommon"
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

			sightData := parser.Sight(*sightSymbols, c.errChannel)

			if sightData != nil && sightData.Time >= c.currentTime {
				c.sightData = *sightData
				c.currentTime = sightData.Time
			}

			select {
			case c.sightChan <- struct{}{}:
			default:
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

			select {
			case c.bodyChan <- struct{}{}:
			default:
			}
		case hearMsg:
			hearSyms, err := lexer.Hear(m.data)
			if err != nil {
				c.errChannel <- err.Error()
				continue
			}

			if hearSyms.Time >= c.currentTime {
				switch hearSyms.Sender {
				case "referee":
					c.playMode = hearSyms.Message
				default:
					c.errChannel <- fmt.Sprintf("ignoring heard message %s", m.data)
				}
			}
		case serverParamMsg:
			c.serverParams.Parse(m.data, c.errChannel)
		case playerParamMsg:
			c.playerParams.Parse(m.data, c.errChannel)
		case playerTypeMsg:
			pType := rcsscommon.ParsePlayerType(m.data, c.errChannel)
			if pType.ID != -1 {
				c.playerTypes[pType.ID] = pType
			}
		case changePlayerTypeMsg:
			cptData, err := parser.ChangePlayerType(m.data)
			if err != nil {
				c.errChannel <- err.Error()
				continue
			}
			if cptData.PlayerType != -1 {
				c.teammateTypes[cptData.Unum] = cptData.PlayerType
			}
		case thinkMsg:
			select {
			case c.thinkChan <- struct{}{}:
			default:
			}
		case errorMsg:
			c.errChannel <- m.String()
		case unsupportedMsg:
			c.errChannel <- fmt.Sprintf("unsupported message received from server: %s", m.String())
		}
	}
}
