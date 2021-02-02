package playerclient

/*
SynchSee enables visual sensor synchronous mode.
This is irreversible. Once a client enables synchronous mode
it can't go back to asynchronous mode.
*/
func (c *Client) SynchSee() string {
	synchSeeStr := "(synch_see)"
	c.cmdChannel <- synchSeeStr
	return synchSeeStr
}
