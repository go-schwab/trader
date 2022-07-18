package data

import (
	"testing"
)

func TestData(t *testing.T) {
	_, err := RealTime("AAPL")

	if err != nil {
		t.Fatalf(err.Error())
	}
	_, err = PriceHistory("AAPL", "month", "1", "daily", "1")

	if err != nil {
		t.Fatalf(err.Error())
	}
}
