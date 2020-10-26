package internal

import (
	"fmt"
	"time"

	"github.com/moosebot/mlbot/config"
	"github.com/sirupsen/logrus"
)

func startTrackingGames(botConfig config.Config, logger *logrus.Logger, ticker *time.Ticker) {
	runTracking(botConfig, logger)
	for {
		select {
		case t := <-ticker.C:
			fmt.Println("Tick at ", t)
			runTracking(botConfig, logger)
		}
	}
}

func runTracking(botConfig config.Config, logger *logrus.Logger) {
	fmt.Println(time.Now())
}
