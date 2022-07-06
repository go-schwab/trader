package option

import (
	"log"
	"net/http"

	. "github.com/samjtro/go-tda/utils"
)

// Covered returns a string; containing covered option calls
func Covered(ticker, contractType, strikeRange, strikeCount, toDate string) string {
	req, _ := http.NewRequest("GET", endpoint_option, nil)
	q := req.URL.Query()
	q.Add("strategy", "COVERED")
	q.Add("symbol", ticker)
	q.Add("contractType", contractType)
	q.Add("range", strikeRange)
	q.Add("strikeCount", strikeCount)
	q.Add("toDate", toDate)
	body, err := Handler(req)

	if err != nil {
		log.Fatal(err)
	}

	return body
}
