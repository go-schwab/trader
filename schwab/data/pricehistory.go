package data

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	req, err := http.NewRequest("GET", fmt.Sprintf(Endpoint_priceHistory, ticker), nil)
	utils.Check(err)
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
	var candles []Candle
	fmt.Println(body)
	err = json.Unmarshal([]byte(body), &candles)
	utils.Check(err)
	return candles, nil
}
