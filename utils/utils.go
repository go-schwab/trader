package utils

import (
	"fmt"
	"log"
	"os/user"
	"strconv"
	"time"

	"github.com/spf13/viper"
)

func Check(err error) {
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func HomeDir() string {
	currentUser, err := user.Current()
	Check(err)

	return fmt.Sprintf("/home/%s", currentUser.Username)
}

func LoadConfig() (err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath("$HOME/.foo/trade/")
	viper.AutomaticEnv()

	return viper.ReadInConfig()
}

// Trim one FIRST character in the string
func TrimOneFirst(s string) string {
	return s[1:]
}

// Trim one LAST character in the string
func TrimOneLast(s string) string {
	return s[:len(s)-1]
}

// Trim one FIRST & one LAST character in the string
func TrimOneFirstOneLast(s string) string {
	return s[1 : len(s)-1]
}

// Trim two FIRST & one LAST character in the string
func TrimTwoFirstOneLast(s string) string {
	return s[2 : len(s)-1]
}

// Trim one FIRST & two LAST character in the string
func TrimOneFirstTwoLast(s string) string {
	return s[1 : len(s)-2]
}

// Trim one FIRST & two LAST character in the string
func TrimOneFirstThreeLast(s string) string {
	return s[1 : len(s)-3]
}

// Now returns a string; containing the current time in ISO 8601 format:
// This is the standard datetime format for quotes and dataframes across this library.
// This function uses your local `$HOME/config.env` for generation of your local time. This requires setting the UTC_DIFF variable in the following format: `-06:00`
func Now(t time.Time) string {
	str := fmt.Sprintf("%d-%d-%dT%d:%d:%d%s",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		viper.Get("UTC_DIFF"))

	return str
}

func UnixToLocal(timestamp string) string {
	i, err := strconv.ParseInt(timestamp, 10, 64)

	if err != nil {
		log.Fatalf(err.Error())
	}

	return Now(time.Unix(i, 0))
}
