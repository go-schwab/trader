package trader_test

import (
	"testing"
)

func TestMarketDataAPI(t *testing.T) {
	_, err := agent.GetQuote("AAPL")
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = agent.GetPriceHistory("AAPL", "month", "1", "daily", "1", "", "")
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = agent.SearchInstrumentSimple("AAPL")
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = agent.SearchInstrumentFundamental("AAPL")
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = agent.GetMovers("$DJI", "up", "percent")
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = agent.GetChains("AAPL")
	if err != nil {
		t.Fatal(err.Error())
	}
}

// Benchmarking functions for market-data_test.go
func BenchmarkGetQuote(b *testing.B) {
	for a := 0; a < b.N; a++ {
		_, _ = agent.GetQuote("AAPL")
	}
}

func BenchmarkPriceHistory(b *testing.B) {
	for a := 0; a < b.N; a++ {
		_, _ = agent.GetPriceHistory("AAPL", "month", "1", "daily", "1", "", "")
	}
}

func BenchmarkSearchInstrumentSimple(b *testing.B) {
	for a := 0; a < b.N; a++ {
		_, _ = agent.SearchInstrumentSimple("AAPL")
	}
}

func BenchmarkSearchInstrumentFundamental(b *testing.B) {
	for a := 0; a < b.N; a++ {
		_, _ = agent.SearchInstrumentFundamental("AAPL")
	}
}

func BenchmarkGetMovers(b *testing.B) {
	for a := 0; a < b.N; a++ {
		_, _ = agent.GetMovers("$DJI", "up", "percent")
	}
}

func BenchmarkGetChains(b *testing.B) {
	for a := 0; a < b.N; a++ {
		_, _ = agent.GetChains("AAPL")
	}
}
