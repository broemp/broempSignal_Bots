package util

import "github.com/spf13/viper"

// Config stores all settings
// Values are read by viper from a .env file or enviroment variables
type Config struct {
	// API Settings
	API_ENDPOINT string `mapstructure:"API_ENDPOINT"`

	// Bot Tokens Settings
	DISCORD_TOKEN  string `mapstructure:"DISCORD_TOKEN"`
	TELEGRAM_TOKEN string `mapstructure:"TELEGRAM_TOKEN"`
}

// LoadConfig reads configuration file/enviroment
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	_ = viper.ReadInConfig()

	err = viper.Unmarshal(&config)
	return
}
