package test

import (
	"testing"

	"github.com/samjtro/go-trade/schwab/data"
)

func TestQuote(t *testing.T) {
	_, err := data.GetQuote("AAPL")

	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestPriceHistory(t *testing.T) {
	_, err := data.GetPriceHistory("AAPL", "month", "1", "daily", "1")

	if err != nil {
		t.Fatalf(err.Error())
	}
}
