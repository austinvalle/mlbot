package main

import (
	"github.com/austinvalle/mlbot/config"
	"github.com/austinvalle/mlbot/internal"
)

func main() {

	botConfig, err := config.LoadConfig(".")
	logger := internal.GetLogger(botConfig.LogLevel)
	if err != nil {
		logger.Fatalf("config error: %v", err)
	}

	internal.StartBot(botConfig, logger)
}
