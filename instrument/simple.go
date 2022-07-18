package instrument

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/samjtro/go-tda/utils"
)

// Simple returns a SIMPLE; with simple fundamental information regarding the desired ticker.
// It takes one parameter:
// cusip = "037833100", etc.
func Simple(ticker string) (SIMPLE, error) {
	req2, _ := http.NewRequest("GET", endpoint_searchinstrument, nil)
	q2 := req2.URL.Query()
	q2.Add("symbol", ticker)
	q2.Add("projection", "fundamental")
	req2.URL.RawQuery = q2.Encode()
	body2, err := utils.Handler(req2)

	if err != nil {
		return SIMPLE{}, err
	}

	var cusip string
	split2 := strings.Split(body2, "\"")

	for i, x := range split2 {
		if x == "cusip" {
			cusip = split2[i+2]
		}
	}

	url := fmt.Sprintf(endpoint_getinstrument, cusip)
	req, _ := http.NewRequest("GET", url, nil)
	body, err := utils.Handler(req)

	if err != nil {
		return SIMPLE{}, err
	}

	var desc, exchange, Type string
	split := strings.Split(body, "\"")

	for i, x := range split {
		if x == "description" {
			desc = split[i+2]
		} else if x == "exchange" {
			exchange = split[i+2]
		} else if x == "assetType" {
			Type = split[i+2]
		}
	}

	return SIMPLE{
		CUSIP:       cusip,
		TICKER:      ticker,
		DESCRIPTION: desc,
		EXCHANGE:    exchange,
		TYPE:        Type,
	}, nil
}
