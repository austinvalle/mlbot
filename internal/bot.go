package internal

import (
	"context"

	"github.com/andersfylling/disgord"
	"github.com/andersfylling/disgord/std"
	"github.com/moosebot/mlbot/config"
	"github.com/sirupsen/logrus"
)

// StartBot will initialize and run the discord bot
func StartBot(botConfig config.Config, logger *logrus.Logger) {

	prefix := botConfig.CommandPrefix

	client := disgord.New(disgord.Config{
		ProjectName: "mlbot",
		BotToken:    botConfig.DiscordBotToken,
		Logger:      logger,
	})
	defer client.StayConnectedUntilInterrupted(context.Background())

	// Default pong message
	log, _ := std.NewLogFilter(client)
	filter, _ := std.NewMsgFilter(context.Background(), client)
	filter.SetPrefix(prefix)
	client.On(disgord.EvtMessageCreate,
		// middleware
		filter.NotByBot,    // ignore bot messages
		filter.HasPrefix,   // read original
		log.LogMsg,         // log command message
		std.CopyMsgEvt,     // read & copy original
		filter.StripPrefix, // write copy
		// handler
		replyPongToPing) // handles copy

	client.On(disgord.EvtGuildCreate, func() {
		logger.Infoln("Guild created!")
	})
}

func replyPongToPing(s disgord.Session, data *disgord.MessageCreate) {
	msg := data.Message
	if msg.Content == "ping" {
		_, _ = msg.Reply(context.Background(), s, "pong")
	}
}
