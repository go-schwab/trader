package test

import (
	"testing"

	"github.com/samjtro/go-trade/schwab/data"
)

func TestMovers(t *testing.T) {
	_, err := data.GetMovers("$DJI", "up", "percent")

	if err != nil {
		t.Fatalf(err.Error())
	}
}
