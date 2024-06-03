package data

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	schwabutils "github.com/samjtro/go-trade/schwab/utils"
	utils "github.com/samjtro/go-trade/utils"
)

// GetMovers returns information on the desired index's movers per your desired direction and change type(percent or value),
// It takes three params:
// index = "$DJI", "$SPX.X", or "$COMPX"
// direction = "up" or "down"
// change = "percent" or "value"
func GetMovers(index, direction, change string) ([]Screener, error) {
	// Craft, send request
	url := fmt.Sprintf(Endpoint_movers, index)
	req, _ := http.NewRequest("GET", url, nil)
	q := req.URL.Query()
	q.Add("direction", direction)
	q.Add("change", change)
	req.URL.RawQuery = q.Encode()
	body, err := schwabutils.Handler(req)
	utils.Check(err)
	// Parse return
	var movers []Screener
	split0 := strings.Split(body, "[")
	split := strings.Split(split0[1], "}")
	lengthToCheckAgainst := len(split)
	lengthToCheck := 1
	for _, x := range split {
		split1 := strings.Split(x, ",")
		var mov Screener
		for _, x1 := range split1 {
			split2 := strings.Split(x1, ":")
			if (split2[0] == "{\"description\"") || (split2[0] == ",{\"description\"") {
				mov.Description = utils.TrimOneFirstOneLast(split2[1])
			}
			if len(split2[0]) > 2 {
				switch utils.TrimOneFirstOneLast(split2[0]) {
				default:
				case "volume":
					mov.Volume, err = strconv.ParseFloat(split2[1], 64)
					utils.Check(err)
				case "lastPrice":
					mov.LastPrice, err = strconv.ParseFloat(split2[1], 64)
					utils.Check(err)
				case "netChange":
					mov.NetChange, err = strconv.ParseFloat(split2[1], 64)
					utils.Check(err)
				case "marketShare":
					mov.MarketShare, err = strconv.ParseFloat(split2[1], 64)
					utils.Check(err)
				case "totalVolume":
					mov.TotalVolume, err = strconv.ParseFloat(split2[1], 64)
					utils.Check(err)
				case "trades":
					mov.Trades, err = strconv.ParseFloat(split2[1], 64)
					utils.Check(err)
				case "netPercentChange":
					mov.NetPercentChange, err = strconv.ParseFloat(split2[1], 64)
					utils.Check(err)
				case "symbol":
					if lengthToCheck < lengthToCheckAgainst {
						mov.Symbol = utils.TrimOneFirstOneLast(split2[1])
					} else {
						mov.Symbol = utils.TrimOneFirstTwoLast(split2[1])
					}
				}
			}
		}
		movers = append(movers, mov)
		lengthToCheck++
	}

	return movers, nil
}
