package trainerclient

import "fmt"

/*
EyeOn prompts the server to start sending (see global ... ) messages every 100ms
(the interval is specified by the send_vi_step parameter)
*/
func (c *Client) EyeOn() string {
	return c.eye("on")

}

/*
EyeOff prompts the server to stop to send visual information automatically. In
this case the trainer/coach has to ask actively with (look), if it needs
visual information.
*/
func (c *Client) EyeOff() string {
	return c.eye("off")
}

func (c *Client) eye(mode string) string {
	eyeStr := fmt.Sprintf("(eye %s)", mode)
	c.cmdChannel <- command{
		time:      c.currentTime,
		cmdString: eyeStr,
	}
	return eyeStr
}
