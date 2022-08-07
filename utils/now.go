package utils

import (
	"fmt"
	"log"
	"time"
)

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
