package internal

import (
	"context"

	"github.com/andersfylling/disgord"
	"github.com/moosebot/mlbot/config"
	"github.com/sirupsen/logrus"
)

// StartBot will initialize and run the discord bot
func StartBot(botConfig config.Config, logger *logrus.Logger) {
	client := disgord.New(disgord.Config{
		ProjectName: "mlbot",
		BotToken:    botConfig.DiscordBotToken,
		Logger:      logger,
	})
	defer client.StayConnectedUntilInterrupted(context.Background())

	client.On(disgord.EvtGuildCreate, RegisterLiveUpdater(botConfig, logger))
}

func replyPongToPing(s disgord.Session, data *disgord.MessageCreate) {
	msg := data.Message
	if msg.Content == "ping" {
		_, _ = msg.Reply(context.Background(), s, "pong")
	}
}
