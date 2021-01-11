package trainerclient

import (
	"fmt"

	"github.com/rcssggb/ggb-lib/rcsscommon"
	"github.com/rcssggb/ggb-lib/trainerclient/lexer"
	"github.com/rcssggb/ggb-lib/trainerclient/parser"
)

// decode continuously receives messages from recvChannel and calls lexer and parser to structure the message data
func (c *Client) decode() {
	var m message

	for {
		m = <-c.recvChannel
		switch m.Type() {
		case thinkMsg:
			select {
			case c.thinkChan <- struct{}{}:
			default:
			}
		case initMsg:
			continue
		case startMsg:
			continue
		case recoverMsg:
			continue
		case moveMsg:
			continue
		case lookMsg:
			gPosSym, err := lexer.Look(m.data)
			if err != nil {
				c.errChannel <- fmt.Sprintf("failed to parse global positions: %s", err)
				continue
			}
			gPos := parser.Look(*gPosSym, c.errChannel)
			if c.currentTime <= gPos.Time {
				c.globalPositions = *gPos
				c.currentTime = gPos.Time
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
		case hearMsg:
			hearSymbols, err := lexer.Hear(m.data)
			if err != nil {
				c.errChannel <- err.Error()
				continue
			}
			if hearSymbols.Time >= c.currentTime {
				c.currentTime = hearSymbols.Time
				switch hearSymbols.Sender {
				case "referee":
					c.playMode = hearSymbols.Message
				default:
					c.errChannel <- fmt.Sprintf("ignoring heard message %s", m.data)
				}
			}
		case errorMsg:
			c.errChannel <- m.String()
		case warningMsg:
			c.errChannel <- m.String()
		case eyeMsg:
			if m.data == "(ok eye on)\x00" {
				c.eyeMode = true
			} else if m.data == "(ok eye off)\x00" {
				c.eyeMode = false
			} else {
				c.errChannel <- fmt.Sprintf("unsupported eye response: %s", m)
			}
		case earMsg:
			if m.data == "(ok ear on)\x00" {
				c.earMode = true
			} else if m.data == "(ok ear off)\x00" {
				c.earMode = false
			} else {
				c.errChannel <- fmt.Sprintf("unsupported ear response: %s", m)
			}
		case teamNamesMsg:
			lTeam, rTeam, err := lexer.TeamNames(m.data)
			if err != nil {
				c.errChannel <- err.Error()
				continue
			}
			c.lTeamName = lTeam
			c.rTeamName = rTeam
		case checkBallMsg:
			time, ballInfo, err := lexer.CheckBall(m.data)
			if err != nil {
				c.errChannel <- err.Error()
				continue
			}
			if c.currentTime <= time {
				c.ballInfo = ballInfo
				c.currentTime = time
			}
		case changePlayerTypeMsg:
			_, _, _, err := lexer.ChangePlayerType(m.data)
			if err != nil {
				c.errChannel <- err.Error()
				continue
			}
		case genericOkMsg:
			continue
		case unsupportedMsg:
			c.errChannel <- fmt.Sprintf("unsupported message received from server: %s", m)
			continue
		}
	}
}
