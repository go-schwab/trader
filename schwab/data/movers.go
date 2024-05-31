package data

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	schwabutils "github.com/samjtro/go-trade/schwab/utils"
	utils "github.com/samjtro/go-trade/utils"
)

// Get returns a string; containing information on the desired index's movers per your desired direction and change type(percent or value),
// it takes three parameters:
// index = "$DJI", "$SPX.X", or "$COMPX"
// direction = "up" or "down"
// change = "percent" or "value"
func GetMovers(index, direction, change string) ([]MOVER, error) {
	url := fmt.Sprintf(Endpoint_movers, index)
	req, _ := http.NewRequest("GET", url, nil)
	q := req.URL.Query()
	q.Add("direction", direction)
	q.Add("change", change)
	req.URL.RawQuery = q.Encode()
	body, err := schwabutils.Handler(req)

	if err != nil {
		return []MOVER{}, err
	}

	var movers []MOVER
	split0 := strings.Split(body, "[")
	split := strings.Split(split0[1], "{")

	// each mover
	for i, x := range split {
		split1 := strings.Split(x, ",")
		fmt.Println(split1)
		for _, x1 := range split1 {
			split2 := strings.Split(x1, ":")
			var mov MOVER
			fmt.Println(split2)
			switch utils.TrimOneFirstOneLast(split2[0]) {
			case "description":
				mov.Description = utils.TrimOneFirstOneLast(split2[1])
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
				mov.Symbol = utils.TrimOneFirstOneLast(split[i+1])
			}
			movers = append(movers, mov)
		}
	}

	return movers, nil
}
