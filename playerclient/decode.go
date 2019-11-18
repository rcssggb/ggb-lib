package playerclient

import (
	"log"

	"github.com/rcssggb/ggb-lib/playerclient/lexer"
	"github.com/rcssggb/ggb-lib/playerclient/parser"
)

// decode continuously receives messages from recvChannel and calls lexer and parser to structure the message data
func (c *Client) decode() {
	var m message
	var err error
	for {
		m = <-c.recvChannel
		switch m.Type() {
		case initMsg:
			_, err = lexer.Init(m.data)
			if err != nil {
				log.Println(err)
			}
		case sightMsg:
			sightSymbols, err := lexer.Sight(m.data)
			if err != nil {
				log.Println(err)
			}
			sightData, err := parser.Sight(*sightSymbols)
			if err != nil {
				log.Println(err)
			}
			if sightData.Time >= c.currentTime {
				c.sightData = *sightData
				c.currentTime = sightData.Time
			}
		case serverParamMsg:
			_, err = lexer.ServerParam(m.data)
		case errorMsg:
			c.errChannel <- m.String()
		case unsupportedMsg:
			continue
		}
		if err != nil {
			log.Println(err)
		}
	}
}
