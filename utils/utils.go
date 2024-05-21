package utils

import "github.com/spf13/viper"

type Config struct {
	APPKEY   string `mapstructure:"APPKEY"`
	SECRET	 string `mapstructure:"SECRET"`
	UTC_DIFF string `mapstructure:"UTC_DIFF"`
}

func LoadConfig() (config Config, err error) {
	viper.SetConfigName("config")
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
