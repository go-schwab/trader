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

	ac, err := agent.GetAccounts()
	if err != nil {
		t.Fatalf(err.Error())
	}
	fmt.Println(ac)
}
