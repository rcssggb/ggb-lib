package playerclient

import "errors"

// Error returns next element in errChannel or nil if channel is empty
func (c *Client) Error() error {
	select {
	case err := <-c.errChannel:
		return errors.New(err)
	default:
		return nil
	}
}
