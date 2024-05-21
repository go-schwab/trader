package test

import (
	"testing"

	"github.com/samjtro/go-trade/schwab"
)

func TestMovers(t *testing.T) {
	_, err := schwab.Get("$DJI", "up", "percent")

	if err != nil {
		t.Fatalf(err.Error())
	}
}
