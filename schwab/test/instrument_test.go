package test

import (
	"testing"

	"github.com/samjtro/go-trade/schwab"
)

func TestInstrument(t *testing.T) {
	_, err := schwab.Simple("AAPL")

	if err != nil {
		t.Fatalf(err.Error())
	}

	_, err = schwab.Fundamental("AAPL")

	if err != nil {
		t.Fatalf(err.Error())
	}
}
