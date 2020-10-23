package main

import (
	"github.com/moosebot/mlbot/config"
	"github.com/moosebot/mlbot/internal"
)

func main() {

	botConfig, err := config.LoadConfig(".")
	logger := internal.GetLogger(botConfig.LogLevel)
	if err != nil {
		logger.Fatalf("config error: %v", err)
	}

	internal.StartBot(botConfig, logger)
}
