package test

import (
	"testing"

	"github.com/samjtro/go-trade/schwab/data"
)

func TestInstrument(t *testing.T) {
	_, err := data.SearchInstrument("AAPL")

	if err != nil {
		t.Fatalf(err.Error())
	}
}
