package playerclient

import "fmt"

// Log prints v to stdout with information about the client
func (c *Client) Log(v interface{}) {
	fmt.Printf("%4d %s %2d: %s\n", c.currentTime, c.teamName, c.shirtNum, v)
}
