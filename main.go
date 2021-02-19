package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/PeterConway314/magic_8_ball/ball"
	"github.com/PeterConway314/magic_8_ball/utils"
)

func main() {
	log.SetLevel(log.DebugLevel)
	utils.SeedRNG()

	ballType := "standard"
	responses, err := utils.Load(fmt.Sprintf("responses/%s.json", ballType))
	if err != nil {
		log.Errorf("Error loading responses from file: %s", err)
		return
	}

	magicBall := ball.Ball{responses}
	fmt.Printf("%s ball says: %s", ballType, magicBall.GetResponse())
}
