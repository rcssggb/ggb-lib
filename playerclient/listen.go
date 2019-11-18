package playerclient

import (
	"log"
	"time"
)

// listen continuously receives and posts server messages to recvChannel
func (c *Client) listen() {
	response := make([]byte, 8192)

	for {
		c.conn.SetReadDeadline(time.Now().Add(1 * time.Second))
		n, serverAddr, err := c.conn.ReadFromUDP(response)
		if c.serverAddr == nil {
			c.serverAddr = serverAddr
		}
		now := time.Now()
		if err != nil {
			log.Println(err)
			continue
		}
		c.recvChannel <- message{
			timestamp: now,
			data:      string(response[:n]),
		}
	}
}
