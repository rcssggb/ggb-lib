package trainerclient

import "fmt"

/*
EarOn prompts the server to send all auditory information to the trainer
*/
func (c *Client) EarOn() string {
	return c.ear("on")

}

/*
EarOff prompts the server to stop sending all auditory information to the trainer
*/
func (c *Client) EarOff() string {
	return c.ear("off")
}

func (c *Client) ear(mode string) string {
	earStr := fmt.Sprintf("(ear %s)", mode)
	c.cmdChannel <- command{
		time:      c.currentTime,
		cmdString: earStr,
	}
	return earStr
}
