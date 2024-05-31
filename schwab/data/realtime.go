package data

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	schwabutils "github.com/samjtro/go-trade/schwab/utils"
	utils "github.com/samjtro/go-trade/utils"
)

// Quote returns a []FRAME with the previous 7 candles.
// It takes one paramter:
// ticker = "AAPL", etc.
func GetCandles(ticker string) ([]CANDLE, error) {
	url := fmt.Sprintf(Endpoint_quote, ticker)
	req, _ := http.NewRequest("GET", url, nil)
	body, err := schwabutils.Handler(req)

	if err != nil {
		log.Fatalf(err.Error())
	}

	var open, hi, lo, Close, volume float64
	var datetime string
	var candles []CANDLE
	split := strings.Split(body, "},")

	for _, x1 := range split {
		for i2, x2 := range strings.Split(x1, "\"") {
			switch x2 {
			case "open":
				open, err = strconv.ParseFloat(utils.TrimFL(split[i2+1]), 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "high":
				hi, err = strconv.ParseFloat(utils.TrimFL(split[i2+1]), 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "low":
				lo, err = strconv.ParseFloat(utils.TrimFL(split[i2+1]), 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "close":
				Close, err = strconv.ParseFloat(utils.TrimFL(split[i2+1]), 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "volume":
				volume, err = strconv.ParseFloat(utils.TrimFL(split[i2+1]), 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "datetime":
				datetime = utils.UnixToLocal(utils.TrimFL(split[i2+1]))
			}
		}

		candles = append(candles, CANDLE{
			Open:     open,
			Hi:       hi,
			Lo:       lo,
			Close:    Close,
			Volume:   volume,
			Datetime: datetime,
		})
	}

	return candles, nil
}

// Quote returns a QUOTE; containing a real time quote of the desired stock's performance with a number of different indicators (including volatility, volume, price, fundamentals & more).
// It takes one parameter:
// ticker = "AAPL", etc.
func GetQuote(tickers string) (QUOTE, error) {
	dt := utils.Now(time.Now())
	url := fmt.Sprintf(Endpoint_quotes, tickers)
	req, _ := http.NewRequest("GET", url, nil)
	body, err := schwabutils.Handler(req)

	if err != nil {
		return QUOTE{}, err
	}

	var bid, ask, last, open, hi, lo, closeP, mark, volume, volatility, hi52, lo52, pe float64
	split := strings.Split(body, "\"")

	// Check for multiple tickers, iterate through them if need be
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
		Ticker:     tickers,
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
