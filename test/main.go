package main

import (
	"flag"
	"log"
	"time"

	playerclient "github.com/rcssggb/ggb-lib/playerclient"
)

const logPath = "/logs/ggb-team.log"

var verbose = flag.Bool("verbose", true, "print info level logs to stdout")

func main() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)

	hostName := "rcssserver"

	player, err := playerclient.NewPlayerClient("ggb-lib-test", hostName)
	if err != nil {
		log.Fatalln(err)
	}

	time.Sleep(1 * time.Second)

	for {
		currentTime := player.Time()

		if currentTime == 0 {
			player.Move(currentTime, -25, -15)
		} else {
			player.Dash(currentTime, 50, float64(currentTime))
		}
		err = player.Error()
		if err != nil {
			log.Println(err)
		}

		for player.Time() <= currentTime {
			time.Sleep(10 * time.Millisecond)
		}
	}
}
