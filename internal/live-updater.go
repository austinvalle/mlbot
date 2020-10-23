package internal

import (
	"github.com/andersfylling/disgord"
	"github.com/moosebot/mlbot/config"
	"github.com/sirupsen/logrus"
)

type channelSubscription struct {
	id        disgord.Snowflake
	name      string
	frequency string
	teamCode  string
}

// RegisterLiveUpdater returns the configured handler for starting the channel subscriptions to live updates
func RegisterLiveUpdater(botConfig config.Config, logger *logrus.Logger) func(s disgord.Session, evt *disgord.GuildCreate) {
	return func(s disgord.Session, evt *disgord.GuildCreate) {
		for _, channel := range evt.Guild.Channels {
			for _, teamConfig := range botConfig.Teams {
				if channel.Name == teamConfig.ChannelName {
					startChannelSubscription(logger, channelSubscription{
						id:        channel.ID,
						name:      channel.Name,
						frequency: teamConfig.UpdateType,
						teamCode:  teamConfig.TeamCode,
					})
					break
				}
			}
		}
	}
}

func startChannelSubscription(logger *logrus.Logger, sub channelSubscription) {
	logger.Infof(
		"Starting '%v' team subscription for '%v (%v)' with the frequency '%v'",
		sub.teamCode,
		sub.name,
		sub.id,
		sub.frequency)
}
