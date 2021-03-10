package playerclient

import (
	"fmt"
	"strings"
	"time"
)

// execute continuously receives messages from cmdChannel and sends them to server
func (c *Client) execute() {
	var cmd string
	var err error
	for {
		select {
		case cmd = <-c.cmdChannel:
		case <-time.After(10 * time.Minute):
			c.errChannel <- "execute loop timed out"
			return
		}

		// Wait until client receives player port
		for c.serverAddr == nil {
			time.Sleep(1 * time.Millisecond)
		}

		// Append "\x00" to command
		if !strings.HasSuffix(cmd, "\x00") {
			cmd = cmd + "\x00"
		}

		_, err = c.conn.WriteToUDP([]byte(cmd), c.serverAddr)
		if err != nil {
			c.errChannel <- fmt.Sprintf("error sending command to server: %s", err)
		}
	}
}
