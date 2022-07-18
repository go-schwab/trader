package instrument

import (
	"testing"
)

func TestInstrument(t *testing.T) {
	_, err := Simple("AAPL")

	if err != nil {
		t.Fatalf(err.Error())
	}

	_, err = Fundamental("AAPL")

	if err != nil {
		t.Fatalf(err.Error())
	}
}
