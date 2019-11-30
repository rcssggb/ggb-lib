package playerclient

import (
	"log"

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
				log.Println(err)
			}
			c.shirtNum = initData.Unum
			c.teamSide = initData.Side
			c.playMode = initData.PlayMode
		case sightMsg:
			sightSymbols, err := lexer.Sight(m.data)
			if err != nil {
				log.Println(err)
				continue
			}
			sightData, err := parser.Sight(*sightSymbols)
			if err != nil {
				log.Println(err)
				continue
			}
			if sightData.Time >= c.currentTime {
				c.sightData = *sightData
				c.currentTime = sightData.Time
			}
		case serverParamMsg:
			// _, err := lexer.ServerParam(m.data)
		case errorMsg:
			c.errChannel <- m.String()
		case unsupportedMsg:
			continue
		}
	}
}
