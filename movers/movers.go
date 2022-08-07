package movers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/samjtro/go-tda/utils"
)

type MOVER struct {
	TICKER      string
	DESCRIPTION string
	LAST        float64
	VOLUME      float64
	DIRECTION   string
	CHANGE      float64
}

var (
	endpoint_movers string = "https://api.tdameritrade.com/v1/marketdata/%s/movers" // index
)

// Get returns a string; containing information on the desired index's movers per your desired direction and change type(percent or value),
// it takes three parameters:
// index = "$DJI", "$SPX.X", or "$COMPX"
// direction = "up" or "down"
// change = "percent" or "value"
func Get(index, direction, change string) ([]MOVER, error) {
	url := fmt.Sprintf(endpoint_movers, index)
	req, _ := http.NewRequest("GET", url, nil)
	q := req.URL.Query()
	q.Add("direction", direction)
	q.Add("change", change)
	req.URL.RawQuery = q.Encode()
	body, err := utils.Handler(req)

	if err != nil {
		return []MOVER{}, err
	}

	var movers []MOVER
	var ticker, desc, dir string
	var chang, last, volume float64
	split := strings.Split(body, "}")

	for _, x := range split {
		split2 := strings.Split(x, "\"")

		for i, x := range split2 {
			switch x {
			case "change":
				chang1 := utils.TrimFL(split2[i+1])

				chang, err = strconv.ParseFloat(chang1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "description":
				desc = split2[i+2]
			case "direction":
				dir = split2[i+2]
			case "last":
				last1 := utils.TrimFL(split2[i+1])

				last, err = strconv.ParseFloat(last1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "symbol":
				ticker = split2[i+2]
			case "totalVolume":
				volume1 := utils.TrimF(split2[i+1])

				volume, err = strconv.ParseFloat(volume1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			}
		}

		mov := MOVER{
			TICKER:      ticker,
			DESCRIPTION: desc,
			LAST:        last,
			CHANGE:      chang,
			DIRECTION:   dir,
			VOLUME:      volume,
		}

		movers = append(movers, mov)
	}

	return movers, nil
}
