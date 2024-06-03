package test

import (
	"fmt"
	"testing"

	"github.com/samjtro/go-trade/schwab/data"
)

func TestQuote(t *testing.T) {
	quote, err := data.GetQuotes("AAPL")

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
