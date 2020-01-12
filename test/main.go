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
		sight := player.See()
		body := player.SenseBody()

		if currentTime == 0 {
			player.Move(currentTime, -5, 0)
		} else if sight.Ball == nil {
			player.Turn(currentTime, 20)
		} else {
			ballAngle := sight.Ball.Direction + body.HeadAngle
			player.Dash(currentTime, 70, ballAngle-85)
			player.TurnNeck(currentTime, sight.Ball.Direction)
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
