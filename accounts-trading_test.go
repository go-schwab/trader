package trader_test

import (
	"testing"
)

func TestAccountsTradingAPI(t *testing.T) {
	an, err := agent.GetAccountNumbers()
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = agent.GetAccounts()
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = agent.GetAccount(an[0].HashValue)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = agent.GetAllOrders("2023-06-12T00:00:00.000Z", "2024-06-12T00:00:00.000Z")
	if err != nil {
		t.Fatal(err.Error())
	}
}

// Benchmarking functions for accounts-trading_test.go
func BenchmarkGetAccountNumbers(b *testing.B) {
	for a := 0; a < b.N; a++ {
		_, _ = agent.GetAccountNumbers()
	}
}

func BenchmarkGetAccounts(b *testing.B) {
	for a := 0; a < b.N; a++ {
		_, _ = agent.GetAccounts()
	}
}

func BenchmarkGetAccount(b *testing.B) {
	for a := 0; a < b.N; a++ {
		_, _ = agent.GetAccount(an[0].HashValue)
	}
}

/*WIP:
func BenchmarkGetOrders(b *testing.B) {
	for a := 0; a < b.N; a++ {
		_, _ = agent.GetAllOrders("2023-06-12T00:00:00.000Z", "2024-06-12T00:00:00.000Z")
	}
}*/
