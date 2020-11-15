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

	playersL := map[int]*playerclient.Client{}

	for i := 0; i < 11; i++ {
		p, err := playerclient.NewPlayerClient("ggb-lib-A", hostName)
		playersL[i] = p
		if err != nil {
			log.Fatalln(err)
		}
	}

	playersR := map[int]*playerclient.Client{}

	for i := 0; i < 11; i++ {
		p, err := playerclient.NewPlayerClient("ggb-lib-B", hostName)
		playersR[i] = p
		if err != nil {
			log.Fatalln(err)
		}
	}

	trainer, err := trainerclient.NewTrainerClient(hostName)
	if err != nil {
		log.Fatalln(err)
	}

	time.Sleep(2 * time.Second)

	trainer.Log(trainer.EyeOn())
	trainer.TeamNames()
	trainer.Start()

	for {
		for i := 0; i < 11; i++ {
			player := playersL[i]
			sight := player.See()
			body := player.SenseBody()

			if sight.Ball == nil {
				player.Turn(30)
			} else {
				ballAngle := sight.Ball.Direction + body.HeadAngle
				ballDist := sight.Ball.Distance
				if ballDist < 0.7 {
					player.Kick(20, 0)
				} else {
					player.Dash(50, ballAngle)
					player.TurnNeck(sight.Ball.Direction)
				}
			}

			err = player.Error()
			for err != nil {
				player.Log(err)
				err = player.Error()
			}
		}

		for i := 0; i < 11; i++ {
			player := playersR[i]
			sight := player.See()
			body := player.SenseBody()

			if sight.Ball == nil {
				player.Turn(30)
			} else {
				ballAngle := sight.Ball.Direction + body.HeadAngle
				ballDist := sight.Ball.Distance
				if ballDist < 0.7 {
					player.Kick(20, 0)
				} else {
					player.Dash(50, ballAngle)
					player.TurnNeck(sight.Ball.Direction)
				}
			}

			err = player.Error()
			for err != nil {
				player.Log(err)
				err = player.Error()
			}
		}

		currentTime := trainer.Time()
		if (currentTime+1)%300 == 0 {
			ballPos := rcsscommon.RandomBallPosition()
			trainer.Log(trainer.MoveBall(ballPos.X, ballPos.Y, 0, 0))
		}

		err = trainer.Error()
		for err != nil {
			trainer.Log(err)
			err = trainer.Error()
		}

		trainer.WaitNextStep(currentTime)
	}
}
