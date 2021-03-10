package playerclient

import (
	"fmt"
	"time"
)

// listen continuously receives and posts server messages to recvChannel
func (c *Client) listen() {
	errCount := 0
	response := make([]byte, 8192)
	defer c.conn.Close()
	for {
		c.conn.SetReadDeadline(time.Now().Add(2 * time.Second))
		n, serverAddr, err := c.conn.ReadFromUDP(response)
		if c.serverAddr == nil {
			c.serverAddr = serverAddr
		}
		now := time.Now()
		if err != nil {
			errCount++
			c.errChannel <- fmt.Sprintf("error reading from UDP port: %s", err)
			if errCount > 10 {
				break
			}
			continue
		}
		c.recvChannel <- message{
			timestamp: now,
			data:      string(response[:n]),
		}
	}
}
