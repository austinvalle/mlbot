package internal

import (
	"context"
	"time"

	"github.com/andersfylling/disgord"
	"github.com/moosebot/mlbot/config"
	"github.com/sirupsen/logrus"
)

// StartBot will initialize and run the discord bot
func StartBot(botConfig config.Config, logger *logrus.Logger) {
	scheduleGameTracker(botConfig, logger)

	client := disgord.New(disgord.Config{
		ProjectName: "mlbot",
		BotToken:    botConfig.DiscordBotToken,
		Logger:      logger,
	})
	defer client.StayConnectedUntilInterrupted(context.Background())
	client.On(disgord.EvtGuildCreate, RegisterLiveUpdater(botConfig, logger))
}

func scheduleGameTracker(botConfig config.Config, logger *logrus.Logger) {
	interval := time.Duration(botConfig.GameTrackerInterval) * time.Minute

	go startTrackingGames(botConfig, logger, time.NewTicker(interval))
}
