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
	fmt.Println(quote)

	ph, err := agent.GetPriceHistory("AAPL", "month", "1", "daily", "1", "", "")
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(ph)

	sis, err := agent.SearchInstrumentSimple("AAPL")
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(sis)

	sif, err := agent.SearchInstrumentFundamental("AAPL")
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(sif)

	movers, err := agent.GetMovers("$DJI", "up", "percent")
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(movers)

	chains, err := agent.GetChains("AAPL")
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(chains)
}
