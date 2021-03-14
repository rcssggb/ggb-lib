package trainerclient

import (
	"fmt"
	"net"
	"sync"

	"github.com/rcssggb/ggb-lib/rcsscommon"
	"github.com/rcssggb/ggb-lib/trainerclient/types"
)

// Client ...
type Client struct {
	mutex           sync.RWMutex
	conn            *net.UDPConn
	serverAddr      *net.UDPAddr
	cmdChannel      chan string
	recvChannel     chan message
	errChannel      chan string
	thinkChan       chan struct{}
	currentTime     int
	serverParams    rcsscommon.ServerParams
	playerParams    rcsscommon.PlayerParams
	globalPositions types.GlobalPositions
	eyeMode         bool
	earMode         bool
	playerTypes     map[int64]rcsscommon.PlayerType
	lTeamName       string
	rTeamName       string
	playMode        rcsscommon.PlayModeID
	ballInfo        string
}

// NewTrainerClient is the constructor for the trainerclient.Client object
func NewTrainerClient(serverIP string) (*Client, error) {
	// Instantiate new Player struct
	client := &Client{}
	client.cmdChannel = make(chan string, 32)
	client.recvChannel = make(chan message, 32)
	client.errChannel = make(chan string, 32)
	client.thinkChan = make(chan struct{}, 32)
	client.playerTypes = make(map[int64]rcsscommon.PlayerType)
	client.currentTime = 0
	client.serverParams = rcsscommon.DefaultServerParams()
	client.playerParams = rcsscommon.DefaultPlayerParams()
	client.playMode = rcsscommon.PlayModeBeforeKickOff

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
