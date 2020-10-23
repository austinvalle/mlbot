package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// LoadConfig will attempt to read the .toml config file and return a struct representing the config
func LoadConfig(configPath string) (Config, error) {
	var botConfig Config
	viper.SetConfigName("config")
	viper.AddConfigPath(configPath)
	viper.BindEnv("MLBOT_DISCORD_TOKEN")

	if readError := viper.ReadInConfig(); readError != nil {
		return Config{}, fmt.Errorf("unable to read config: %v", readError)
	}

	if unmarshalError := viper.Unmarshal(&botConfig); unmarshalError != nil {
		return Config{}, fmt.Errorf("unable to unmarshal config: %v", unmarshalError)
	}

	return botConfig, nil
}
