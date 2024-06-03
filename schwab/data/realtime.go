package data

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	schwabutils "github.com/samjtro/go-trade/schwab/utils"
	utils "github.com/samjtro/go-trade/utils"
)

// GetCandles returns a []Candle with the previous 7 candles.
// It takes one paramter:
// ticker = "AAPL", etc.
func GetCandles(ticker string) ([]Candle, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(Endpoint_quote, ticker), nil)
	utils.Check(err)
	body, err := schwabutils.Handler(req)
	utils.Check(err)
	var candles []Candle
	err = json.Unmarshal([]byte(body), &candles)
	utils.Check(err)
	return candles, nil
}

// Quote returns a Quote; containing a real time quote of the desired stock's performance with a number of different indicators (including volatility, volume, price, fundamentals & more).
// It takes one parameter:
// ticker = "AAPL", etc.
func GetQuotes(tickers string) (Quote, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(Endpoint_quotes, tickers), nil)
	utils.Check(err)
	body, err := schwabutils.Handler(req)
	utils.Check(err)
	// WIP
	var quote Quote
	err = json.Unmarshal([]byte(body), &quote)
	utils.Check(err)
	quote.Time = utils.Now(time.Now())
	return quote, err
}

// func GetQuotes() []QUOTE {}
