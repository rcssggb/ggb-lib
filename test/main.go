package main

import (
	"flag"
	"log"
	"time"

	playerclient "github.com/rcssggb/ggb-lib/playerclient"
	"github.com/rcssggb/ggb-lib/rcsscommon"
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

	trainer.Log(trainer.EyeOn())
	trainer.TeamNames()
	trainer.MovePlayer("ggb-lib-test", 1, -5, 0, 0, 0, 0)
	trainer.Start()

	for {
		currentTime := player.Time()
		sight := player.See()
		body := player.SenseBody()

		if sight.Ball == nil {
			player.Turn(currentTime, 30)
		} else {
			ballAngle := sight.Ball.Direction + body.HeadAngle
			ballDist := sight.Ball.Distance
			if ballDist < 0.7 {
				player.Kick(currentTime, 20, 0)
			} else {
				player.Dash(currentTime, 70, ballAngle)
				player.TurnNeck(currentTime, sight.Ball.Direction)
			}
		}

		if currentTime%100 == 0 {
			ballPos := rcsscommon.RandomBallPosition()
			trainer.Log(trainer.MoveBall(ballPos.X, ballPos.Y, 0, 0))
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
