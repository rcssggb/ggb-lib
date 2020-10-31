package playerclient

import (
	"fmt"
	"net"

	"github.com/rcssggb/ggb-lib/playerclient/parser"
	"github.com/rcssggb/ggb-lib/rcsscommon"
)

// Client ...
// TODO: unify playerclient and trainerclient common functions and variables
// in a clientcommon package
type Client struct {
	conn         *net.UDPConn
	serverAddr   *net.UDPAddr
	cmdChannel   chan string
	recvChannel  chan message
	errChannel   chan string
	currentTime  int
	serverParams rcsscommon.ServerParams
	playerParams rcsscommon.PlayerParams
	playerTypes  map[int64]rcsscommon.PlayerType
	sightData    parser.SightData
	bodyData     parser.SenseBodyData
	teamName     string
	teamSide     rcsscommon.SideType
	shirtNum     int
	playMode     string
}

// NewPlayerClient is the constructor for the playerclient.Client object
func NewPlayerClient(teamName, serverIP string) (*Client, error) {
	// Instantiate new Player struct
	client := &Client{}
	client.teamName = teamName
	client.cmdChannel = make(chan string, 32)
	client.recvChannel = make(chan message, 32)
	client.errChannel = make(chan string, 32)
	client.playerTypes = make(map[int64]rcsscommon.PlayerType)
	client.currentTime = 0
	client.serverParams = rcsscommon.DefaultServerParams()
	client.playerParams = rcsscommon.DefaultPlayerParams()

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
