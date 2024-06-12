// fromEnteredTime, toEnteredTime format:
// yyyy-MM-ddTHH:mm:ss.SSSZ
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
}
