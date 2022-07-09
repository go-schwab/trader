package data

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	. "github.com/samjtro/go-tda/utils"
)

// PriceHistory returns a []FRAME; containing a series of candles with price volume & datetime info per candlestick.
// It takes five parameters:
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
	split = split[2:]

	for _, x := range split {
		split2 := strings.Split(x, "\"")

		for i, x2 := range split2 {
			switch x2 {
			case "open":
				open = split2[i+1]
			case "high":
				hi = split2[i+1]
			case "low":
				lo = split2[i+1]
			case "close":
				Close = split2[i+1]
			case "volume":
				volume = split2[i+1]
			case "datetime":
				datetime = split2[i+1]
			}
		}

		f := FRAME{
			Datetime: TrimL(TrimFL(datetime)),
			Volume:   TrimFL(volume),
			Open:     TrimFL(open),
			Close:    TrimFL(Close),
			Hi:       TrimFL(hi),
			Lo:       TrimFL(lo),
		}

		df = append(df, f)
	}

	return df
}
