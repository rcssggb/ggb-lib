package playerclient

import (
	"fmt"
	"log"
	"net"

	"github.com/rcssggb/ggb-lib/playerclient/parser"
)

// Client ...
type Client struct {
	conn        *net.UDPConn
	recvChannel chan message
	sightData   parser.SightData
	TeamName    string
	TeamSide    sideType
	ShirtNum    int
	PlayMode    string
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
	client.TeamName = teamName
	client.recvChannel = make(chan message, 32)

	go client.listen()
	go client.decode()

	// Send connect message
	log.Printf("Connecting to %s as a player for %s\n", serverHost, teamName)
	cmdMessage := fmt.Sprintf("(init %s (version 15))", teamName)
	_, err = conn.WriteToUDP([]byte(cmdMessage), serverAddr)
	if err != nil {
		return nil, err
	}

	return client, nil
}
