package option

import (
	"net/http"

	"github.com/samjtro/go-tda/utils"
)

// Covered returns a string; containing covered option calls.
// Not functional ATM.
func Covered(ticker, contractType, strikeRange, strikeCount, toDate string) (string, error) {
	req, _ := http.NewRequest("GET", endpoint_option, nil)
	q := req.URL.Query()
	q.Add("strategy", "COVERED")
	q.Add("symbol", ticker)
	q.Add("contractType", contractType)
	q.Add("range", strikeRange)
	q.Add("strikeCount", strikeCount)
	q.Add("toDate", toDate)
	body, err := utils.Handler(req)

	if err != nil {
		return "", err
	}

	return body, nil
}
