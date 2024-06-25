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

	instrument := InstrumentRef{
		Symbol: "BTTX",
		Type:   "EQUITY",
	}
	newMarketOrder := CreateMarketOrder(Session("NORMAL"), Duration("DAY"), Strategy("SINGLE"), Leg(OrderLeg{
		Instruction: "BUY",
		Quantity:    1,
		Instrument:  instrument,
	}))
	err = agent.Submit(an[0].HashValue, newMarketOrder)

	if err != nil {
		t.Fatalf(err.Error())
	}
}
