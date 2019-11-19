package playerclient

import (
	"fmt"
	"log"
	"time"
)

// execute continuously receives messages from cmdChannel and sends them to server
func (c *Client) execute() {
	var cmd command
	var err error
	for {
		cmd = <-c.cmdChannel

		// Check if command is supposed to be sent in current time
		if cmd.time < c.currentTime {
			err = fmt.Errorf("command too late, discarding")
			log.Println(err)
			continue
		}
		if cmd.time > c.currentTime {
			err = fmt.Errorf("command too early, putting it back to channel")
			log.Println(err)
			c.cmdChannel <- cmd
		}

		// Wait until client receives player port
		for c.serverAddr == nil {
			time.Sleep(1 * time.Millisecond)
		}

		_, err = c.conn.WriteToUDP([]byte(cmd.cmdString), c.serverAddr)
		if err != nil {
			err = fmt.Errorf("error sending command to server: %s", err)
			log.Println(err)
		}
	}
}
