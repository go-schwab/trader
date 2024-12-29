package trader

import (
	"fmt"
	"testing"
)

func TestMarketDataAPI(t *testing.T) {
	agent := Initiate()

	quote, err := agent.GetQuote("AAPL")

	if err != nil {
		t.Fatal(err.Error())
	}

	ph, err := agent.GetPriceHistory("AAPL", "month", "1", "daily", "1", "", "")

	if err != nil {
		t.Fatal(err.Error())
	}

	sis, err := agent.SearchInstrumentSimple("AAPL")

	if err != nil {
		t.Fatal(err.Error())
	}

	sif, err := agent.SearchInstrumentFundamental("AAPL")

	if err != nil {
		t.Fatal(err.Error())
	}

	movers, err := agent.GetMovers("$DJI", "up", "percent")

	if err != nil {
		t.Fatal(err.Error())
	}

	chains, err := agent.GetChains("AAPL")

	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println(quote)
	fmt.Println(ph)
	fmt.Println(sis)
	fmt.Println(sif)
	fmt.Println(movers)
	fmt.Println(chains)
}
