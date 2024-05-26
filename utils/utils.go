package utils

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	APPKEY   string `mapstructure:"APPKEY"`
	SECRET   string `mapstructure:"SECRET"`
	UTC_DIFF string `mapstructure:"UTC_DIFF"`
	DBPATH   string `mapstructure:"DBPATH"`
	CBURL    string `mapstructure:"CBURL"`
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

// Trim the FIRST character in the string
func TrimF(s string) string {
	return s[1:]
}

// Trim the LAST character in the string
func TrimL(s string) string {
	return s[:len(s)-1]
}

// Trim the FIRST & LAST character in the string
func TrimFL(s string) string {
	str := s[:len(s)-1]
	return str[1:]
}

// Now returns a string; containing the current time in ISO 8601 format:
// This is the standard datetime format for quotes and dataframes across this library.
// This function uses your local `$HOME/config.env` for generation of your local time. This requires setting the UTC_DIFF variable in the following format: `-06:00`
func Now(t time.Time) string {
	config, err := LoadConfig()

	if err != nil {
		log.Fatalf(err.Error())
	}

	str := fmt.Sprintf("%d-%d-%dT%d:%d:%d%s",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		config.UTC_DIFF)

	return str
}

func UnixToLocal(timestamp string) string {
	i, err := strconv.ParseInt(timestamp, 10, 64)

	if err != nil {
		log.Fatalf(err.Error())
	}

	return Now(time.Unix(i, 0))
}
