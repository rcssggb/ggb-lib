package playerclient

import (
	"fmt"
	"net"

	"github.com/rcssggb/ggb-lib/playerclient/parser"
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
	sightData   parser.SightData
	bodyData    parser.SenseBodyData
	teamName    string
	teamSide    rcsscommon.SideType
	shirtNum    int
	playMode    string
}

// NewPlayerClient is the constructor for the playerclient.Client object
func NewPlayerClient(teamName, serverIP string) (*Client, error) {
	// Instantiate new Player struct
	client := &Client{}
	client.teamName = teamName
	client.cmdChannel = make(chan command, 32)
	client.recvChannel = make(chan message, 32)
	client.errChannel = make(chan string, 32)
	client.currentTime = 0

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
	client.conn = conn

	go client.listen()
	go client.decode()
	go client.execute()

	// Send connect message
	fmt.Printf("Connecting to %s as a player for %s\n", serverHost, teamName)
	cmdMessage := fmt.Sprintf("(init %s (version 16))\x00", teamName)
	_, err = conn.WriteToUDP([]byte(cmdMessage), serverAddr)
	if err != nil {
		return nil, err
	}

	return client, nil
}
