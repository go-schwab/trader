package options

import (
	"fmt"
	"net/http"
)

// simpleOption takes four parameters:
// ticker = "AAPL", etc.
// contractType = "CALL", "PUT", "ALL"
// strikeCount = number of strikes to return above and below the at-the-money price
// includeQuotes = "TRUE", "FALSE"
func simpleOption(ticker,contractType,strikeCount,includeQuotes string) string {
	req,_ := http.NewRequest("GET",endpoint_option,nil)
	q := req.URL.Query()
	q.Add("symbol",ticker)
	q.Add("contractType",contractType)
	q.Add("strikeCount",strikeCount)
	q.Add("includeQuotes",includeQuotes)
	req.URL.RawQuery = q.Encode()
	body := handler(req)

	return body
}

