package test

import (
	"testing"

	"github.com/samjtro/go-trade/schwab"
)

func TestQuote(t *testing.T) {
	_, err := schwab.RealTime("AAPL")

	if err != nil {
		t.Fatalf(err.Error())
	}	
}

func TestPriceHistory(t *testing.T) {
	_, err = schwab.PriceHistory("AAPL", "month", "1", "daily", "1")

	if err != nil {
		t.Fatalf(err.Error())
	}
}
