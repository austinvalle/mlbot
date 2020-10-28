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
	todaysDate := time.Now().Format("01/02/2006")
	// datePtr := flag.String("date", todaysDate, "the date to get MLB games for (defaults to today)")
	// flag.Parse()
	getCurrentLiveGames(todaysDate, logger)
	// Store info about games and their start times
	// Somehow schedule with live tracker if the bot config is good to go?
}
