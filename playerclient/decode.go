package playerclient

import (
	"log"

	"./lexer"
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
			_, err = lexer.Sight(m.data)
		case serverParamMsg:
			_, err = lexer.ServerParam(m.data)
		case unsupportedMsg:
			continue
		}
		if err != nil {
			log.Println(err)
		}
	}
}
