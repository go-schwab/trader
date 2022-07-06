package option

import (
	"log"
	"net/http"

	. "github.com/samjtro/go-tda/utils"
)

// ANALYTICAL returns a string; allows you to control additional parameters for theoretical value calculations;
// it takes nine parameters
func Analytical(ticker, contractType, strikeRange, strikeCount, toDate, volatility, underlyingPrice, interestRate, daysToExpiration string) string {
	req, _ := http.NewRequest("GET", endpoint_option, nil)
	q := req.URL.Query()
	q.Add("strategy", "ANALYTICAL")
	q.Add("symbol", ticker)
	q.Add("contractType", contractType)
	q.Add("range", strikeRange)
	q.Add("strikeCount", strikeCount)
	q.Add("toDate", toDate)
	q.Add("volatility", volatility)
	q.Add("underlyingPrice", underlyingPrice)
	q.Add("interestRate", interestRate)
	q.Add("daysToExpiration", underlyingPrice)
	req.URL.RawQuery = q.Encode()
	body, err := Handler(req)

	if err != nil {
		log.Fatal(err)
	}

	return body
}
