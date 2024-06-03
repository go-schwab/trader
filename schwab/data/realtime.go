package data

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	schwabutils "github.com/samjtro/go-trade/schwab/utils"
	utils "github.com/samjtro/go-trade/utils"
)

// GetCandles returns a []Candle with the previous 7 candles.
// It takes one paramter:
// ticker = "AAPL", etc.
func GetCandles(ticker string) ([]Candle, error) {
	var candles []Candle
	// Craft, send request
	url := fmt.Sprintf(Endpoint_quote, ticker)
	req, _ := http.NewRequest("GET", url, nil)
	body, err := schwabutils.Handler(req)
	utils.Check(err)
	// Parse return
	split := strings.Split(body, "},")
	for _, x1 := range split {
		var candle Candle
		for i2, x2 := range strings.Split(x1, "\"") {
			switch x2 {
			case "open":
				candle.Open, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i2+1]), 64)
				utils.Check(err)
			case "high":
				candle.Hi, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i2+1]), 64)
				utils.Check(err)
			case "low":
				candle.Lo, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i2+1]), 64)
				utils.Check(err)
			case "close":
				candle.Close, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i2+1]), 64)
				utils.Check(err)
			case "volume":
				candle.Volume, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i2+1]), 64)
				utils.Check(err)
			case "datetime":
				candle.Time = utils.UnixToLocal(utils.TrimOneFirstOneLast(split[i2+1]))
			}
		}
		candles = append(candles, candle)
	}
	return candles, nil
}

// Quote returns a Quote; containing a real time quote of the desired stock's performance with a number of different indicators (including volatility, volume, price, fundamentals & more).
// It takes one parameter:
// ticker = "AAPL", etc.
func GetQuotes(tickers string) (Quote, error) {
	var quote Quote
	// Craft, send request
	quote.Time = utils.Now(time.Now())
	url := fmt.Sprintf(Endpoint_quotes, tickers)
	req, err := http.NewRequest("GET", url, nil)
	utils.Check(err)
	body, err := schwabutils.Handler(req)
	utils.Check(err)
	// Parse return
	split := strings.Split(body, "\"")
	// WIP: Split, iterate thru tickers
	for i, x := range split {
		switch x {
		case "bidPrice":
			quote.Bid, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "askPrice":
			quote.Ask, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "lastPrice":
			quote.Last, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "openPrice":
			quote.Open, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "highPrice":
			quote.Hi, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "lowPrice":
			quote.Lo, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "closePrice":
			quote.Close, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "mark":
			quote.Mark, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "totalVolume":
			quote.Volume, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "volatility":
			quote.Volatility, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "52WkHigh":
			quote.Hi52, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "52WkLow":
			quote.Lo52, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "peRatio":
			quote.PE, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		}
	}
	return quote, nil
}

// func GetQuotes() []QUOTE {}
