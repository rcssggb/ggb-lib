package trainerclient

import "fmt"

/*
ChangePlayerType changes the player type
*/
func (c *Client) ChangePlayerType(unum, playerType int64, teamName string) string {
	changePlayerTypeString := fmt.Sprintf("(change_player_type %s %d %d)", teamName, unum, playerType)
	c.cmdChannel <- changePlayerTypeString
	return changePlayerTypeString
}
