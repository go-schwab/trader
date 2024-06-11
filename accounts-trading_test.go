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
}
