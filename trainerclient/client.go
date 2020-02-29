package trainerclient

import (
	"fmt"
	"net"

	"github.com/rcssggb/ggb-lib/rcsscommon"
)

// Client ...
type Client struct {
	conn        *net.UDPConn
	serverAddr  *net.UDPAddr
	cmdChannel  chan command
	recvChannel chan message
	errChannel  chan string
	currentTime int
	teamSide    rcsscommon.SideType
	playMode    string
}

// NewTrainerClient is the constructor for the trainerclient.Client object
func NewTrainerClient(serverIP string) (*Client, error) {
	// Instantiate new Player struct
	client := &Client{}
	client.cmdChannel = make(chan command, 32)
	client.recvChannel = make(chan message, 32)
	client.errChannel = make(chan string, 32)
	client.currentTime = 0

	// Open listener socker to request player connection
	serverHost := serverIP + ":6001"
	serverAddr, err := net.ResolveUDPAddr("udp", serverHost)
	if err != nil {
		return nil, err
	}

	conn, err := net.ListenUDP("udp", nil)
	if err != nil {
		return nil, err
	}
	client.conn = conn

	go client.listen()
	go client.decode()
	go client.execute()

	// Send connect message
	fmt.Printf("Connecting to %s as a trainer\n", serverHost)
	cmdMessage := "(init (version 16))\x00"
	_, err = conn.WriteToUDP([]byte(cmdMessage), serverAddr)
	if err != nil {
		return nil, err
	}

	return client, nil
}