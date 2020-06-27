package trainerclient

/*
CheckBall provides information the ball position
*/
func (c *Client) CheckBall() string {
	checkBallString := "(check_ball)"
	c.cmdChannel <- command{
		cmdString: checkBallString,
	}
	return checkBallString
}
