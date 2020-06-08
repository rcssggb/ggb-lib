package main

import (
	"flag"
	"log"
	"time"

	playerclient "github.com/rcssggb/ggb-lib/playerclient"
	trainerclient "github.com/rcssggb/ggb-lib/trainerclient"
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

	trainer, err := trainerclient.NewTrainerClient(hostName)
	if err != nil {
		log.Fatalln(err)
	}

	time.Sleep(1 * time.Second)

	for {
		currentTime := player.Time()
		sight := player.See()
		body := player.SenseBody()
		playMode := player.PlayMode()

		trainer.Look()

		if currentTime == 0 {
			player.Move(currentTime, -5, 0)

		} else if sight.Ball == nil {
			player.Turn(currentTime, 20)
		} else if playMode == "kick_off_l" {
			ballAngle := sight.Ball.Direction + body.HeadAngle
			ballDist := sight.Ball.Distance
			if ballDist < 0.7 {
				player.Kick(currentTime, 20, 0)
			} else {
				player.Dash(currentTime, 50, ballAngle)
				player.TurnNeck(currentTime, sight.Ball.Direction)
			}
		} else {
			ballAngle := sight.Ball.Direction + body.HeadAngle
			player.Dash(currentTime, 50, ballAngle-85)
			player.TurnNeck(currentTime, sight.Ball.Direction)
		}

		err = player.Error()
		for err != nil {
			player.Log(err)
			err = player.Error()
		}

		err = trainer.Error()
		for err != nil {
			trainer.Log(err)
			err = trainer.Error()
		}

		player.WaitNextStep(currentTime)
	}
}
