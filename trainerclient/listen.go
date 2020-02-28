package trainerclient

import (
	"fmt"
	"time"
)

// listen continuously receives and posts server messages to recvChannel
func (c *Client) listen() {
	response := make([]byte, 8192)

	for {
		n, serverAddr, err := c.conn.ReadFromUDP(response)
		if c.serverAddr == nil {
			c.serverAddr = serverAddr
		}
		now := time.Now()
		if err != nil {
			c.errChannel <- fmt.Sprintf("error reading from UDP port: %s", err)
			continue
		}
		c.recvChannel <- message{
			timestamp: now,
			data:      string(response[:n]),
		}
	}
}
