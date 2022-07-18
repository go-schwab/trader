package option

import (
	"net/http"

	"github.com/samjtro/go-tda/utils"
)

// Butterfly returns a string; containing Butterfly spread option calls.
// Not functional ATM.
func Butterfly(ticker, contractType, strikeRange, strikeCount, toDate string) (string, error) {
	req, _ := http.NewRequest("GET", endpoint_option, nil)
	q := req.URL.Query()
	q.Add("strategy", "BUTTERFLY")
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
