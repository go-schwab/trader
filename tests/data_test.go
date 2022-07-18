package tests

import (
	"testing"

	"github.com/samjtro/go-tda/data"
)

func TestData(t *testing.T) {
	_, err := data.RealTime("AAPL")

	if err != nil {
		t.Fatalf("Encountered Error: %s", err)
	}
}
