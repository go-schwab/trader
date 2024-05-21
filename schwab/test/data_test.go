package test

import (
	"testing"

	"github.com/samjtro/go-tda/data"
)

func TestData(t *testing.T) {
	_, err := data.RealTime("AAPL")

	if err != nil {
		t.Fatalf(err.Error())
	}

	_, err = data.PriceHistory("AAPL", "month", "1", "daily", "1")

	if err != nil {
		t.Fatalf(err.Error())
	}
}
