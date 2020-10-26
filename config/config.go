package config

// Config is the main config struct for the bot
type Config struct {
	GameTrackerInterval int          `mapstructure:"game_tracker_interval_minutes"`
	LiveTrackerInterval int          `mapstructure:"live_tracker_interval_seconds"`
	DiscordBotToken     string       `mapstructure:"MLBOT_DISCORD_TOKEN"`
	LogLevel            string       `mapstructure:"log_level"`
	CommandPrefix       string       `mapstructure:"command_prefix"`
	Teams               []TeamConfig `mapstructure:"team"`
}
