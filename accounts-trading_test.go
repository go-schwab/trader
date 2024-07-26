package schwab

import (
	"fmt"
	"testing"
)

func TestAccountsTradingAPI(t *testing.T) {
	agent := Initiate()

	an, err := agent.GetAccountNumbers()
	if err != nil {
		t.Fatalf(err.Error())
	}
	fmt.Println(an)

	aca, err := agent.GetAccounts()
	if err != nil {
		t.Fatalf(err.Error())
	}
	fmt.Println(aca)

	ac, err := agent.GetAccount(an[0].HashValue)
	if err != nil {
		t.Fatalf(err.Error())
	}
	fmt.Println(ac)

	orders, err := agent.GetAllOrders("2023-06-12T00:00:00.000Z", "2024-06-12T00:00:00.000Z")
	if err != nil {
		t.Fatalf(err.Error())
	}

	fmt.Println(orders)

	instrument := MarketOrderInstrument{
		Symbol:    "BTTX",
		AssetType: "EQUITY",
	}
	newMarketOrder := CreateMarketOrder(MarketSession("NORMAL"), MarketDuration("DAY"), MarketStrategy("SINGLE"), MarketLeg(MarketOrderLeg{
		Instruction: "BUY",
		Quantity:    1,
		Instrument:  instrument,
	}))
	err = agent.SubmitMarketOrder(an[0].HashValue, newMarketOrder)

	if err != nil {
		t.Fatalf(err.Error())
	}
}
