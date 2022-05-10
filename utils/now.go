package utils

import (
	"fmt"
	"time"
)

// Now returns a string; containing the current time in the format:
// YMD-HMS
// this is the standard datetime format for quotes and dataframes across this library
func Now(t time.Time) string {
	str := fmt.Sprintf("%d%d%d-%d%d%d",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second())

	return str
}
