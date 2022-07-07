package data

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	. "github.com/samjtro/go-tda/utils"
)

// RealTime returns a QUOTE; containing a real time quote of the desired stock's performance with a number of different indicators (including volatility, volume, price, fundamentals & more).
// It takes one parameter:
// ticker = "AAPL", etc.
func RealTime(ticker string) QUOTE {
	dt := Now(time.Now())
	url := fmt.Sprintf(endpoint_realtime, ticker)
	req, _ := http.NewRequest("GET", url, nil)
	body, err := Handler(req)

	if err != nil {
		log.Fatal(err)
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
		DATETIME:   dt,
		TICKER:     ticker,
		MARK:       TrimFL(mark),
		VOLUME:     TrimFL(volume),
		VOLATILITY: TrimFL(volatility),
		BID:        TrimFL(bid),
		ASK:        TrimFL(ask),
		LAST:       TrimFL(last),
		OPEN:       TrimFL(open),
		CLOSE:      TrimFL(closeP),
		HI:         TrimFL(hi),
		LO:         TrimFL(lo),
		HI52:       TrimFL(hi52),
		LO52:       TrimFL(lo52),
		PE_RATIO:   TrimFL(pe),
	}
}
