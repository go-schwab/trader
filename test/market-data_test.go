package test

import (
	"testing"

	schwab "github.com/samjtro/schwab-go/schwab"
)

func TestMarketDataAPI(t *testing.T) {
	agent := schwab.Initiate()

	_, err := agent.GetQuote("AAPL")

	if err != nil {
		t.Fatalf(err.Error())
	}

	_, err = agent.GetPriceHistory("AAPL", "month", "1", "daily", "1", "", "")

	if err != nil {
		t.Fatalf(err.Error())
	}

	_, err = agent.SearchInstrumentSimple("AAPL")

	if err != nil {
		t.Fatalf(err.Error())
	}

	_, err = agent.SearchInstrumentFundamental("AAPL")

	if err != nil {
		t.Fatalf(err.Error())
	}

	_, err = agent.GetMovers("$DJI", "up", "percent")

	if err != nil {
		t.Fatalf(err.Error())
	}
}
