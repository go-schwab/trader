package test

import (
	"testing"

	"github.com/samjtro/go-trade/movers"
)

func TestMovers(t *testing.T) {
	_, err := movers.Get("$DJI", "up", "percent")

	if err != nil {
		t.Fatalf(err.Error())
	}
}
