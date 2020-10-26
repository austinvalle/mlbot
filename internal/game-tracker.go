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
	// Call MLB API to find games using current date/time
	// Store info about games and their start times
	// Somehow schedule with live tracker if the bot config is good to go?
	fmt.Println(time.Now())
}
