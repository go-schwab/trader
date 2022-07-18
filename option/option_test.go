package option

import (
	"testing"
)

func TestOption(t *testing.T) {
	_, err := Single("AAPL", "ALL", "ALL", "15", "2022-09-20")

	if err != nil {
		t.Fatalf(err.Error())
	}
}
