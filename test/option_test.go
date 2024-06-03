package test

import (
	"testing"

	data "github.com/samjtro/go-schwab-traderapi/market-data"
)

func TestOption(t *testing.T) {
	_, err := data.Single("AAPL", "ALL", "ALL", "15", "2022-09-20")

	if err != nil {
		t.Fatalf(err.Error())
	}
}
