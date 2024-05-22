package data 

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/samjtro/go-trade/schwab"
	"github.com/samjtro/go-trade/schwab/utils"
)

var (
	endpoint_quote string = fmt.Sprintf(schwab.Endpoint + "/%s/quotes") // Symbol
	// endpoint_quotes string
)

// RealTime's native struct; returns various indicators related to the asset
type QUOTE struct {
	Datetime   string
	Ticker     string
	Mark       float64
	Volume     float64
	Volatility float64
	Bid        float64
	Ask        float64
	Last       float64
	Open       float64
	Close      float64
	Hi         float64
	Lo         float64
	Hi52       float64
	Lo52       float64
	PE         float64
}

// Quote returns a QUOTE; containing a real time quote of the desired stock's performance with a number of different indicators (including volatility, volume, price, fundamentals & more).
// It takes one parameter:
// ticker = "AAPL", etc.
func GetQuote(ticker string) (QUOTE, error) {
	dt := utils.Now(time.Now())
	url := fmt.Sprintf(endpoint_quote, ticker)
	req, _ := http.NewRequest("GET", url, nil)
	body, err := utils.Handler(req)

	if err != nil {
		return QUOTE{}, err
	}

	var bid, ask, last, open, hi, lo, closeP, mark, volume, volatility, hi52, lo52, pe float64
	split := strings.Split(body, "\"")

	for i, x := range split {
		switch x {
		case "bidPrice":
			bid1 := utils.TrimFL(split[i+1])

			bid, err = strconv.ParseFloat(bid1, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "askPrice":
			ask1 := utils.TrimFL(split[i+1])

			ask, err = strconv.ParseFloat(ask1, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "lastPrice":
			last1 := utils.TrimFL(split[i+1])

			last, err = strconv.ParseFloat(last1, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "openPrice":
			open1 := utils.TrimFL(split[i+1])

			open, err = strconv.ParseFloat(open1, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "highPrice":
			hi1 := utils.TrimFL(split[i+1])

			hi, err = strconv.ParseFloat(hi1, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "lowPrice":
			lo1 := utils.TrimFL(split[i+1])

			lo, err = strconv.ParseFloat(lo1, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "closePrice":
			closeP1 := utils.TrimFL(split[i+1])

			closeP, err = strconv.ParseFloat(closeP1, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "mark":
			mark1 := utils.TrimFL(split[i+1])

			mark, err = strconv.ParseFloat(mark1, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "totalVolume":
			volume1 := utils.TrimFL(split[i+1])

			volume, err = strconv.ParseFloat(volume1, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "volatility":
			volatility1 := utils.TrimFL(split[i+1])

			volatility, err = strconv.ParseFloat(volatility1, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "52WkHigh":
			hi521 := utils.TrimFL(split[i+1])

			hi52, err = strconv.ParseFloat(hi521, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "52WkLow":
			lo521 := utils.TrimFL(split[i+1])

			lo52, err = strconv.ParseFloat(lo521, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "peRatio":
			pe1 := utils.TrimFL(split[i+1])

			pe, err = strconv.ParseFloat(pe1, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		}
	}

	return QUOTE{
		Datetime:   dt,
		Ticker:     ticker,
		Mark:       mark,
		Volume:     volume,
		Volatility: volatility,
		Bid:        bid,
		Ask:        ask,
		Last:       last,
		Open:       open,
		Close:      closeP,
		Hi:         hi,
		Lo:         lo,
		Hi52:       hi52,
		Lo52:       lo52,
		PE:         pe,
	}, nil
}

// func GetQuotes() []QUOTE {}
