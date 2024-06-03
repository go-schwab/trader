package data

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	schwabutils "github.com/samjtro/go-trade/schwab/utils"
	utils "github.com/samjtro/go-trade/utils"
)

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
func GetPriceHistory(ticker, periodType, period, frequencyType, frequency, startDate, endDate string) ([]Candle, error) {
	var candles []Candle
	url := fmt.Sprintf(Endpoint_priceHistory, ticker)
	req, _ := http.NewRequest("GET", url, nil)
	q := req.URL.Query()
	q.Add("periodType", periodType)
	q.Add("period", period)
	q.Add("frequencyType", frequencyType)
	q.Add("frequency", frequency)
	q.Add("startDate", startDate)
	q.Add("endDate", endDate)
	req.URL.RawQuery = q.Encode()
	body, err := schwabutils.Handler(req)
	utils.Check(err)
	split := strings.Split(body, "{")
	split = split[2:]
	for _, x := range split {
		var candle Candle
		split2 := strings.Split(x, "\"")
		for i, x2 := range split2 {
			switch x2 {
			case "open":
				candle.Open, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split2[i+1]), 64)
				utils.Check(err)
			case "high":
				candle.Hi, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split2[i+1]), 64)
				utils.Check(err)
			case "low":
				candle.Lo, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split2[i+1]), 64)
				utils.Check(err)
			case "close":
				candle.Close, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split2[i+1]), 64)
				utils.Check(err)
			case "volume":
				candle.Volume, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split2[i+1]), 64)
				utils.Check(err)
			case "datetime":
				candle.Time = utils.TrimOneLast(utils.TrimOneFirstOneLast(split2[i+1]))
			}
		}
		candles = append(candles, candle)
	}

	return candles, nil
}
