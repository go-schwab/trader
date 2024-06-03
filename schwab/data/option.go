package data

import (
	"encoding/json"
	"net/http"

	schwabutils "github.com/samjtro/go-trade/schwab/utils"
	utils "github.com/samjtro/go-trade/utils"
)

// Single returns a []CONTRACT; containing a SINGLE option chain of your desired strike, type, etc.,
// it takes four parameters:
// ticker = "AAPL", etc.
// contractType = "CALL", "PUT", "ALL";
// strikeRange = returns option chains for a given range:
// ITM = in da money
// NTM = near da money
// OTM = out da money
// SAK = strikes above market
// SBK = strikes below market
// SNK = strikes near market
// ALL* = default, all strikes;
// strikeCount = The number of strikes to return above and below the at-the-money price;
// toDate = Only return expirations before this date. Valid ISO-8601 formats are: yyyy-MM-dd and yyyy-MM-dd'T'HH:mm:ssz.
// Lets examine a sample call of Single: Single("AAPL","CALL","ALL","5","2022-07-01").
// This returns 5 AAPL CALL contracts both above and below the at the money price, with no preference as to the status of the contract ("ALL"), expiring before 2022-07-01
func Single(ticker, contractType, strikeRange, strikeCount, toDate string) ([]Contract, error) {
	req, err := http.NewRequest("GET", Endpoint_option, nil)
	utils.Check(err)
	q := req.URL.Query()
	q.Add("symbol", ticker)
	q.Add("contractType", contractType)
	q.Add("range", strikeRange)
	q.Add("strikeCount", strikeCount)
	q.Add("toDate", toDate)
	req.URL.RawQuery = q.Encode()
	body, err := schwabutils.Handler(req)
	utils.Check(err)
	var chain []Contract
	// WIP
	err = json.Unmarshal([]byte(body), &chain)
	utils.Check(err)
	return chain, nil
}

// Covered returns a string; containing covered option calls.
// Not functional ATM.
func Covered(ticker, contractType, strikeRange, strikeCount, toDate string) (string, error) {
	req, _ := http.NewRequest("GET", Endpoint_option, nil)
	q := req.URL.Query()
	q.Add("strategy", "COVERED")
	q.Add("symbol", ticker)
	q.Add("contractType", contractType)
	q.Add("range", strikeRange)
	q.Add("strikeCount", strikeCount)
	q.Add("toDate", toDate)
	body, err := schwabutils.Handler(req)
	utils.Check(err)

	return body, nil
}

// Butterfly returns a string; containing Butterfly spread option calls.
// Not functional ATM.
func Butterfly(ticker, contractType, strikeRange, strikeCount, toDate string) (string, error) {
	req, _ := http.NewRequest("GET", Endpoint_option, nil)
	q := req.URL.Query()
	q.Add("strategy", "BUTTERFLY")
	q.Add("symbol", ticker)
	q.Add("contractType", contractType)
	q.Add("range", strikeRange)
	q.Add("strikeCount", strikeCount)
	q.Add("toDate", toDate)
	body, err := schwabutils.Handler(req)
	utils.Check(err)

	return body, nil
}

// ANALYTICAL returns a string; allows you to control additional parameters for theoretical value calculations:
// It takes nine parameters:
// Not functional ATM.
func Analytical(ticker, contractType, strikeRange, strikeCount, toDate, volatility, underlyingPrice, interestRate, daysToExpiration string) (string, error) {
	req, _ := http.NewRequest("GET", Endpoint_option, nil)
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
	body, err := schwabutils.Handler(req)
	utils.Check(err)

	return body, nil
}

// func Vertical() string {}
// func Calendar() string {}
// func Strangle() string {}
// func Straddle() string {}
// func Condor() string {}
// func Diagonal() string {}
// func Collar() string {}
// func Roll() string {}
