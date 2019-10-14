package playerclient

import (
	"fmt"
	"log"
	"net"
	"time"

	"./parse"
)

// Client ...
type Client struct {
	conn        *net.UDPConn
	teamName    string
	teamSide    sideType
	shirtNum    int
	playMode    string
	recvChannel chan message
}

// NewPlayerClient is the constructor for the playerclient.Client object
func NewPlayerClient(teamName, serverIP string) (*Client, error) {
	// Open listener socker to request player connection
	serverHost := serverIP + ":6000"
	serverAddr, err := net.ResolveUDPAddr("udp", serverHost)
	if err != nil {
		return nil, err
	}
	conn, err := net.ListenUDP("udp", nil)
	if err != nil {
		return nil, err
	}

	// Instantiate new Player struct
	client := &Client{}
	client.conn = conn
	client.teamName = teamName
	client.recvChannel = make(chan message, 32)

	go client.listen()
	go client.parse()

	// Send connect message
	log.Printf("Connecting to %s as a player for %s\n", serverHost, teamName)
	cmdMessage := fmt.Sprintf("(init %s (version 15))", teamName)
	_, err = conn.WriteToUDP([]byte(cmdMessage), serverAddr)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// pisten continuously receives and posts server messages to recvChannel
func (c *Client) listen() {
	response := make([]byte, 8192)

	for {
		c.conn.SetReadDeadline(time.Now().Add(1 * time.Second))
		n, _, err := c.conn.ReadFromUDP(response)
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

// parse continuously receives messages from recvChannel and calls parser to structure the message data
func (c *Client) parse() {
	var m message
	var err error
	for {
		m = <-c.recvChannel
		switch m.Type() {
		case initMsg:
			_, err = parse.Init(m.data)
		case sightMsg:
			_, err = parse.Sight(m.data)
		case serverParamMsg:
			_, err = parse.ServerParam(m.data)
		case unsupportedMsg:
			continue
		}
		if err != nil {
			log.Println(err)
		}
	}
}
