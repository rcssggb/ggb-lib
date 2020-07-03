package trainerclient

import (
	"fmt"
	"time"
)

// execute continuously receives messages from cmdChannel and sends them to server
func (c *Client) execute() {
	var cmd command
	var err error
	for {
		cmd = <-c.cmdChannel

		// TODO: maybe this is not really necessary?
		// Check if command is supposed to be sent in current time
		if cmd.time < c.currentTime {
			c.errChannel <- fmt.Sprintf("warning: %s for time %d was too late, discarding", cmd, cmd.time)
			continue
		}
		if cmd.time > c.currentTime {
			c.cmdChannel <- cmd
		}

		// Wait until client receives player port
		for c.serverAddr == nil {
			time.Sleep(1 * time.Millisecond)
		}

		_, err = c.conn.WriteToUDP([]byte(cmd.cmdString+"\x00"), c.serverAddr)
		if err != nil {
			c.errChannel <- fmt.Sprintf("error sending command to server: %s", err)
		}
	}
}
