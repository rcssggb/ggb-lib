package trainerclient

/*
CheckBall provides information the ball position
*/
func (c *Client) CheckBall() string {
	checkBallString := "(check_ball)"
	c.cmdChannel <- command{
		time:      c.currentTime,
		cmdString: checkBallString,
	}
	return checkBallString
}
