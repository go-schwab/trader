package test

import (
	"testing"

	"github.com/samjtro/go-trade/schwab/data"
)

func TestOption(t *testing.T) {
	_, err := data.Single("AAPL", "ALL", "ALL", "15", "2022-09-20")

	if err != nil {
		t.Fatalf(err.Error())
	}
}
