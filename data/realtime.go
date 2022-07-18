package data

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/samjtro/go-tda/utils"
)

// RealTime returns a QUOTE; containing a real time quote of the desired stock's performance with a number of different indicators (including volatility, volume, price, fundamentals & more).
// It takes one parameter:
// ticker = "AAPL", etc.
func RealTime(ticker string) (QUOTE, error) {
	dt := utils.Now(time.Now())
	url := fmt.Sprintf(endpoint_realtime, ticker)
	req, _ := http.NewRequest("GET", url, nil)
	body, err := utils.Handler(req)

	if err != nil {
		return QUOTE{}, err
	}

	var bid, ask, last, open, hi, lo, closeP, mark, volume, volatility, hi52, lo52, pe string
	split := strings.Split(body, "\"")

	for i, x := range split {
		switch x {
		case "bidPrice":
			bid = split[i+1]
		case "askPrice":
			ask = split[i+1]
		case "lastPrice":
			last = split[i+1]
		case "openPrice":
			open = split[i+1]
		case "highPrice":
			hi = split[i+1]
		case "lowPrice":
			lo = split[i+1]
		case "closePrice":
			closeP = split[i+1]
		case "mark":
			mark = split[i+1]
		case "totalVolume":
			volume = split[i+1]
		case "volatility":
			volatility = split[i+1]
		case "52WkHigh":
			hi52 = split[i+1]
		case "52WkLow":
			lo52 = split[i+1]
		case "peRatio":
			pe = split[i+1]
		}
	}

	return QUOTE{
		Datetime:   dt,
		Ticker:     ticker,
		Mark:       utils.TrimFL(mark),
		Volume:     utils.TrimFL(volume),
		Volatility: utils.TrimFL(volatility),
		Bid:        utils.TrimFL(bid),
		Ask:        utils.TrimFL(ask),
		Last:       utils.TrimFL(last),
		Open:       utils.TrimFL(open),
		Close:      utils.TrimFL(closeP),
		Hi:         utils.TrimFL(hi),
		Lo:         utils.TrimFL(lo),
		Hi52:       utils.TrimFL(hi52),
		Lo52:       utils.TrimFL(lo52),
		PE:         utils.TrimFL(pe),
	}, nil
}
