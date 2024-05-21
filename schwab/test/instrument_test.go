package test

import (
	"testing"

	"github.com/samjtro/go-trade/instrument"
)

func TestInstrument(t *testing.T) {
	_, err := instrument.Simple("AAPL")

	if err != nil {
		t.Fatalf(err.Error())
	}

	_, err = instrument.Fundamental("AAPL")

	if err != nil {
		t.Fatalf(err.Error())
	}
}
