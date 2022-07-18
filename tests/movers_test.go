package tests

import (
	"testing"

	"github.com/samjtro/go-tda/movers"
)

func TestMovers(t *testing.T) {
	_, err := movers.Get("$DJI", "up", "percent")

	if err != nil {
		t.Fatalf(err.Error())
	}
}
