package test

import (
	"testing"

	data "github.com/samjtro/go-schwab-traderapi/market-data"
)

func TestMovers(t *testing.T) {
	_, err := data.GetMovers("$DJI", "up", "percent")

	if err != nil {
		t.Fatalf(err.Error())
	}
}
