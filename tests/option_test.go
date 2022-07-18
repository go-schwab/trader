package tests

import (
	"testing"

	"github.com/samjtro/go-tda/option"
)

func TestOption(t *testing.T) {
	_, err := option.Single("AAPL", "ALL", "ALL", "15", "2022-09-20")

	if err != nil {
		t.Fatalf(err.Error())
	}
}
