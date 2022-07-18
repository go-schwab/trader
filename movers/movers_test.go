package movers

import (
	"testing"
)

func TestMovers(t *testing.T) {
	_, err := Get("$DJI", "up", "percent")

	if err != nil {
		t.Fatalf(err.Error())
	}
}
