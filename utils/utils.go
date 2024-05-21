package utils

import "github.com/spf13/viper"

type Config struct {
	APIKEY   string `mapstructure:"APIKEY"`
	UTC_DIFF string `mapstructure:"UTC_DIFF"`
}

func LoadConfig() (config Config, err error) {
	viper.SetConfigName("tda-config")
	viper.SetConfigType("env")
	viper.AddConfigPath("$HOME/")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
