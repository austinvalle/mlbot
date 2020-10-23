package config

// TeamConfig contains the configuration for each channel subscription
type TeamConfig struct {
	ChannelName string `mapstructure:"channel_name"`
	TeamCode    string `mapstructure:"team_code"`
	UpdateType  string `mapstructure:"update_type"`
}
