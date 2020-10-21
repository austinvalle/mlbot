package main

import (
	"context"
	"fmt"
	"os"

	"github.com/andersfylling/disgord"
	"github.com/andersfylling/disgord/std"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var log = &logrus.Logger{
	Out:       os.Stderr,
	Formatter: new(logrus.TextFormatter),
	Hooks:     make(logrus.LevelHooks),
	Level:     logrus.ErrorLevel,
}

type config struct {
	CommandPrefix string       `mapstructure:"command_prefix"`
	Teams         []teamConfig `mapstructure:"team"`
}

type teamConfig struct {
	ChannelName string `mapstructure:"channel_name"`
}

var botConfig config

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("unable to read config: %v", err)
	}

	err := viper.Unmarshal(&botConfig)
	if err != nil {
		log.Fatalf("unable to decode into struct: %v", err)
	}

	fmt.Println(botConfig)

	prefix := botConfig.CommandPrefix

	client := disgord.New(disgord.Config{
		ProjectName: "mlbot",
		BotToken:    os.Getenv("MLBOT_DISCORD_TOKEN"),
		Logger:      log,
		Presence: &disgord.UpdateStatusPayload{
			Game: &disgord.Activity{
				Name: "write " + prefix + "ping",
			},
		},
	})
	defer client.StayConnectedUntilInterrupted(context.Background())

	log, _ := std.NewLogFilter(client)
	filter, _ := std.NewMsgFilter(context.Background(), client)
	filter.SetPrefix(prefix)

	// create a handler and bind it to new message events
	// tip: read the documentation for std.CopyMsgEvt and understand why it is used here.
	client.On(disgord.EvtMessageCreate,
		// middleware
		filter.NotByBot,    // ignore bot messages
		filter.HasPrefix,   // read original
		log.LogMsg,         // log command message
		std.CopyMsgEvt,     // read & copy original
		filter.StripPrefix, // write copy
		// handler
		replyPongToPing) // handles copy
}

// replyPongToPing is a handler that replies pong to ping messages
func replyPongToPing(s disgord.Session, data *disgord.MessageCreate) {
	msg := data.Message

	// whenever the message written is "ping", the bot replies "pong"
	if msg.Content == "ping" {
		_, _ = msg.Reply(context.Background(), s, "pong")
	}
}
