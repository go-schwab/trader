package data

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	. "github.com/samjtro/go-tda/utils"
)

// PriceHistory returns a []FRAME; containing a series of candles with price volume & datetime info per candlestick,
// it takes five parameters:
// ticker = "AAPL", etc.
// periodType = "day", "month", "year", "ytd" - default is "day"
// period = the number of periods to show
// frequencyType = the type of frequency with which each candle is formed; valid fTypes by pType
// "day": "minute"
// "month": "daily", "weekly"
// "year": "daily", "weekly", "monthly"
// "ytd": "daily", "weekly"
// frequency = the number of the frequencyType included in each candle; valid freqs by fType
// "minute": 1,5,10,15,30
// "daily": 1
// "weekly": 1
// "monthly": 1
func PriceHistory(ticker, periodType, period, frequencyType, frequency string) []FRAME {
	url := fmt.Sprintf(endpoint_pricehistory, ticker)
	req, _ := http.NewRequest("GET", url, nil)
	q := req.URL.Query()
	q.Add("periodType", periodType)
	q.Add("period", period)
	q.Add("frequencyType", frequencyType)
	q.Add("frequency", frequency)
	req.URL.RawQuery = q.Encode()
	body, err := Handler(req)

	if err != nil {
		log.Fatal(err)
	}

	var df []FRAME
	var open, hi, lo, Close, volume, datetime string

	split := strings.Split(body, "{")
	split = split[2:len(split)]

	for _, x := range split {
		split2 := strings.Split(x, "\"")
		for i, x2 := range split2 {
			if x2 == "open" {
				open = split2[i+1]
			} else if x2 == "high" {
				hi = split2[i+1]
			} else if x2 == "low" {
				lo = split2[i+1]
			} else if x2 == "close" {
				Close = split2[i+1]
			} else if x2 == "volume" {
				volume = split2[i+1]
			} else if x2 == "datetime" {
				datetime = split2[i+1]
			}
		}

		f := FRAME{
			DATETIME: TrimL(TrimFL(datetime)),
			VOLUME:   TrimFL(volume),
			OPEN:     TrimFL(open),
			CLOSE:    TrimFL(Close),
			HI:       TrimFL(hi),
			LO:       TrimFL(lo),
		}

		df = append(df, f)
	}

	return df
}
