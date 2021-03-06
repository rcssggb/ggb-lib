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

	nPlayers := 1

	hostName := "rcssserver"

	playersL := map[int]*playerclient.Client{}

	for i := 0; i < nPlayers; i++ {
		p, err := playerclient.NewPlayerClient("ggb-lib-A", hostName)
		playersL[i] = p
		if err != nil {
			log.Fatalln(err)
		}
	}

	playersR := map[int]*playerclient.Client{}

	for i := 0; i < nPlayers; i++ {
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
	trainer.EarOn()
	trainer.EyeOn()

	time.Sleep(1 * time.Second)

	for i := 0; i < nPlayers; i++ {
		p := playersL[i]
		go player(p)
		p = playersR[i]
		go player(p)
	}

	trainer.TeamNames()
	serverParams := trainer.ServerParams()
	for {
		currentTime := trainer.Time()
		if (currentTime+1)%300 == 0 {
			ballPos := rcsscommon.RandomBallPosition()
			trainer.MoveBall(ballPos.X, ballPos.Y, 0, 0)
		}

		if trainer.PlayMode().String() == "before_kick_off" {
			trainer.Start()
		}

		if trainer.PlayMode().String() == "time_over" {
			break
		}

		err = trainer.Error()
		for err != nil {
			trainer.Log(err)
			err = trainer.Error()
		}

		if serverParams.SynchMode {
			trainer.DoneSynch()
			trainer.WaitSynch()
		} else {
			trainer.WaitNextStep(currentTime)
		}
	}

	time.Sleep(2 * time.Second)
}

func player(c *playerclient.Client) {
	c.Log(c.PlayMode())
	serverParams := c.ServerParams()
	c.SynchSee()
	c.ChangeView(rcsscommon.ViewWidthNarrow, rcsscommon.ViewQualityHigh)
	for {
		sight := c.See()
		body := c.SenseBody()
		currentTime := c.Time()

		if c.PlayMode().String() == "time_over" {
			c.Bye()
			break
		}

		if sight.Ball == nil {
			c.Turn(30)
		} else {
			ballAngle := sight.Ball.Direction + body.HeadAngle
			ballDist := sight.Ball.Distance
			if ballDist < 0.7 {
				c.Kick(20, 0)
			} else {
				c.Dash(50, ballAngle)
				c.TurnNeck(sight.Ball.Direction)
			}
		}

		err := c.Error()
		for err != nil {
			c.Log(err)
			err = c.Error()
		}

		if serverParams.SynchMode {
			c.DoneSynch()
			c.WaitSynch()
		} else {
			c.WaitNextStep(currentTime)
		}
	}
}
