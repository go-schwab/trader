package test

import (
	"fmt"
	"testing"

	data "github.com/samjtro/schwab-go/market-data"
)

func TestQuote(t *testing.T) {
	quote, err := data.GetQuote("AAPL")

	if err != nil {
		t.Fatalf(err.Error())
	}

	fmt.Println(quote)
}

func TestPriceHistory(t *testing.T) {
	priceHistory, err := data.GetPriceHistory("AAPL", "month", "1", "daily", "1", "", "")

	if err != nil {
		t.Fatalf(err.Error())
	}

	fmt.Println(priceHistory)
}
