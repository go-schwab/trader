/*
Copyright (C) 2025 github.com/go-schwab

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, see
<https://www.gnu.org/licenses/>.
*/

package trader

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/bytedance/sonic"
)

var (
	endpoint                 string = "https://api.schwabapi.com/marketdata/v1"
	endpointQuote            string = endpoint + "/%s/quotes" // Symbol
	endpointQuotes           string = endpoint + "/quotes"
	endpointPriceHistory     string = endpoint + "/pricehistory"
	endpointSearchInstrument string = endpoint + "/instruments"
	endpointMovers           string = endpoint + "/movers/%s" // Index ID
	endpointOptions          string = endpoint + "/chains"
)

// WIP: func GetQuotes(symbols string) Quote, error) {}

// Quote returns a Quote; containing a real time quote of the desired stock's performance with a number of different indicators (including volatility, volume, price, fundamentals & more).
// It takes one parameter:
// ticker = "AAPL", etc.
func (agent *Agent) GetQuote(symbol string) (Quote, error) {
	req, err := http.NewRequest("GET", endpointQuotes, nil)
	if err != nil {
		return Quote{}, err
	}
	q := req.URL.Query()
	q.Add("symbols", symbol)
	q.Add("fields", "quote")
	req.URL.RawQuery = q.Encode()
	resp, err := agent.Handler(req)
	if err != nil {
		return Quote{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Quote{}, err
	}
	var quote Quote
	err = sonic.Unmarshal([]byte(strings.Join(strings.Split(strings.Split(string(body), fmt.Sprintf("\"%s\":", symbol))[1], "\"quote\":{"), "")[:len(strings.Join(strings.Split(strings.Split(string(body), fmt.Sprintf("\"%s\":", symbol))[1], "\"quote\":{"), ""))-2]), &quote)
	if err != nil {
		return Quote{}, err
	}
	return quote, err
}

// SearchInstrumentSimple returns instrument's simples.
// It takes one param:
func (agent *Agent) SearchInstrumentSimple(symbols string) (SimpleInstrument, error) {
	req, err := http.NewRequest("GET", endpointSearchInstrument, nil)
	if err != nil {
		return SimpleInstrument{}, err
	}
	q := req.URL.Query()
	q.Add("symbol", symbols)
	q.Add("projection", "symbol-search")
	req.URL.RawQuery = q.Encode()
	resp, err := agent.Handler(req)
	if err != nil {
		return SimpleInstrument{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return SimpleInstrument{}, err
	}
	var instrument SimpleInstrument
	err = sonic.Unmarshal([]byte(strings.Split(string(body), "[")[1][:len(strings.Split(string(body), "[")[1])-2]), &instrument)
	if err != nil {
		return SimpleInstrument{}, err
	}
	return instrument, nil
}

// SearchInstrumentFundamental returns instrument's fundamentals.
// It takes one param:
func (agent *Agent) SearchInstrumentFundamental(symbol string) (FundamentalInstrument, error) {
	req, err := http.NewRequest("GET", endpointSearchInstrument, nil)
	if err != nil {
		return FundamentalInstrument{}, err
	}
	q := req.URL.Query()
	q.Add("symbol", symbol)
	q.Add("projection", "fundamental")
	req.URL.RawQuery = q.Encode()
	resp, err := agent.Handler(req)
	if err != nil {
		return FundamentalInstrument{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return FundamentalInstrument{}, err
	}
	var instrument FundamentalInstrument
	split0 := strings.Split(string(body), "[{\"fundamental\":")[1]
	split := strings.Split(split0, "}")
	err = sonic.Unmarshal([]byte(fmt.Sprintf("%s}", strings.Join(split[:2], ""))), &instrument)
	if err != nil {
		return FundamentalInstrument{}, err
	}
	return instrument, nil
}

// PriceHistory returns a series of candles with price volume & datetime info per candlestick.
// It takes seven parameters:
// ticker = "AAPL", etc.;
// periodType = "day", "month", "year", "ytd" - default is "day";
// period = the number of periods to show;
// frequencyType = the type of frequency with which each candle is formed; valid fTypes by pType;
// "day": "minute" /
// "month": "daily", "weekly" /
// "year": "daily", "weekly", "monthly" /
// "ytd": "daily", "weekly";
// frequency = the number of the frequencyType included in each candle; valid freqs by fType
// "minute": 1,5,10,15,30 /
// "daily": 1 /
// "weekly": 1 /
// "monthly": 1;
// startDate =
// endDate =
func (agent *Agent) GetPriceHistory(symbol, periodType, period, frequencyType, frequency, startDate, endDate string) ([]Candle, error) {
	req, err := http.NewRequest("GET", endpointPriceHistory, nil)
	if err != nil {
		return []Candle{}, err
	}
	q := req.URL.Query()
	q.Add("symbol", symbol)
	q.Add("periodType", periodType)
	q.Add("period", period)
	q.Add("frequencyType", frequencyType)
	q.Add("frequency", frequency)
	q.Add("startDate", startDate)
	q.Add("endDate", endDate)
	req.URL.RawQuery = q.Encode()
	resp, err := agent.Handler(req)
	if err != nil {
		return []Candle{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []Candle{}, err
	}
	var candles []Candle
	err = sonic.Unmarshal([]byte(fmt.Sprintf("[%s]", strings.Split(strings.Split(string(body), "[")[1], "]")[0])), &candles)
	if err != nil {
		return []Candle{}, err
	}
	return candles, nil
}

// GetMovers returns information on the desired index's movers per your desired direction and change type(percent or value),
// It takes three params:
// index = "$DJI", "$SPX.X", or "$COMPX"
// direction = "up" or "down"
// change = "percent" or "value"
func (agent *Agent) GetMovers(index, direction, change string) ([]Screener, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(endpointMovers, index), nil)
	if err != nil {
		return []Screener{}, err
	}
	q := req.URL.Query()
	q.Add("direction", direction)
	q.Add("change", change)
	req.URL.RawQuery = q.Encode()
	resp, err := agent.Handler(req)
	if err != nil {
		return []Screener{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []Screener{}, err
	}
	var movers []Screener
	stringToParse := fmt.Sprintf("[%s]", strings.Split(string(body), "[")[1][:len(strings.Split(string(body), "[")[1])-2])
	err = sonic.Unmarshal([]byte(stringToParse), &movers)
	if err != nil {
		return []Screener{}, err
	}
	return movers, nil
}

// get all option chains for a ticker
func (agent *Agent) GetChains(symbol string) (Chain, error) {
	req, err := http.NewRequest("GET", endpointOptions, nil)
	if err != nil {
		return Chain{}, err
	}
	q := req.URL.Query()
	q.Add("symbol", symbol)
	req.URL.RawQuery = q.Encode()
	resp, err := agent.Handler(req)
	if err != nil {
		return Chain{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Chain{}, err
	}
	var chain Chain
	err = sonic.Unmarshal(body, &chain)
	if err != nil {
		return Chain{}, err
	}
	return chain, nil
}

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
func (agent *Agent) Single(symbol, contractType, strikeRange, strikeCount, toDate string) (Chain, error) {
	req, err := http.NewRequest("GET", endpointOptions, nil)
	if err != nil {
		return Chain{}, nil
	}
	q := req.URL.Query()
	q.Add("symbol", symbol)
	q.Add("contractType", contractType)
	q.Add("range", strikeRange)
	q.Add("strikeCount", strikeCount)
	q.Add("toDate", toDate)
	req.URL.RawQuery = q.Encode()
	resp, err := agent.Handler(req)
	if err != nil {
		return Chain{}, nil
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Chain{}, nil
	}
	var chain Chain
	// WIP
	err = sonic.Unmarshal(body, &chain)
	if err != nil {
		return Chain{}, nil
	}
	return chain, nil
}

/* Covered returns a string; containing covered option calls.
func Covered(ticker, contractType, strikeRange, strikeCount, toDate string) (string, error) {
	req, _ := http.NewRequest("GET", endpointOptions, nil)
	q := req.URL.Query()
	q.Add("strategy", "COVERED")
	q.Add("symbol", ticker)
	q.Add("contractType", contractType)
	q.Add("range", strikeRange)
	q.Add("strikeCount", strikeCount)
	q.Add("toDate", toDate)
	body, err := utils.Handler(req)
	utils.Check(err)

	return body, nil
}

// Butterfly returns a string; containing Butterfly spread option calls.
func Butterfly(ticker, contractType, strikeRange, strikeCount, toDate string) (string, error) {
	req, _ := http.NewRequest("GET", endpointOptions, nil)
	q := req.URL.Query()
	q.Add("strategy", "BUTTERFLY")
	q.Add("symbol", ticker)
	q.Add("contractType", contractType)
	q.Add("range", strikeRange)
	q.Add("strikeCount", strikeCount)
	q.Add("toDate", toDate)
	body, err := utils.Handler(req)
	utils.Check(err)

	return body, nil
}

// ANALYTICAL returns a string; allows you to control additional parameters for theoretical value calculations:
// It takes nine parameters:
func Analytical(ticker, contractType, strikeRange, strikeCount, toDate, volatility, underlyingPrice, interestRate, daysToExpiration string) (string, error) {
	req, _ := http.NewRequest("GET", endpointOptions, nil)
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
	body, err := utils.Handler(req)
	utils.Check(err)

	return body, nil
}

func Vertical() string {}
func Calendar() string {}
func Strangle() string {}
func Straddle() string {}
func Condor() string {}
func Diagonal() string {}
func Collar() string {}
func Roll() string {}*/
