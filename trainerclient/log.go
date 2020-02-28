package trainerclient

import "fmt"

// Log prints v to stdout with information about the client
func (c *Client) Log(v interface{}) {
	fmt.Printf("%4d trainer: %s\n", c.currentTime, v)
}
