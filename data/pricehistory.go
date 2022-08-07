package data

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/samjtro/go-tda/utils"
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
func PriceHistory(ticker, periodType, period, frequencyType, frequency string) ([]FRAME, error) {
	url := fmt.Sprintf(endpoint_pricehistory, ticker)
	req, _ := http.NewRequest("GET", url, nil)
	q := req.URL.Query()
	q.Add("periodType", periodType)
	q.Add("period", period)
	q.Add("frequencyType", frequencyType)
	q.Add("frequency", frequency)
	req.URL.RawQuery = q.Encode()
	body, err := utils.Handler(req)

	if err != nil {
		return []FRAME{}, err
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
				open = utils.TrimFL(split2[i+1])
			case "high":
				hi = utils.TrimFL(split2[i+1])
			case "low":
				lo = utils.TrimFL(split2[i+1])
			case "close":
				Close = utils.TrimFL(split2[i+1])
			case "volume":
				volume = utils.TrimFL(split2[i+1])
			case "datetime":
				datetime = utils.TrimFL(split2[i+1])
			}
		}

		volume, err := strconv.ParseFloat(volume, 64)

		if err != nil {
			log.Fatalf(err.Error())
		}

		open, err := strconv.ParseFloat(open, 64)

		if err != nil {
			log.Fatalf(err.Error())
		}

		Close, err := strconv.ParseFloat(Close, 64)

		if err != nil {
			log.Fatalf(err.Error())
		}

		hi, err := strconv.ParseFloat(hi, 64)

		if err != nil {
			log.Fatalf(err.Error())
		}

		lo, err := strconv.ParseFloat(lo, 64)

		if err != nil {
			log.Fatalf(err.Error())
		}

		f := FRAME{
			Datetime: utils.TrimL(datetime),
			Volume:   volume,
			Open:     open,
			Close:    Close,
			Hi:       hi,
			Lo:       lo,
		}

		df = append(df, f)
	}

	return df, nil
}
