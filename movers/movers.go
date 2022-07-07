package movers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	. "github.com/samjtro/go-tda/utils"
)

type MOVER struct {
	TICKER      string
	DESCRIPTION string
	LAST        string
	CHANGE      string
	DIRECTION   string
	VOLUME      string
}

var endpoint_movers string = "https://api.tdameritrade.com/v1/marketdata/%s/movers" // index

// Get returns a string; containing information on the desired index's movers per your desired direction and change type(percent or value),
// it takes three parameters:
// index = "$DJI", "$SPX.X", or "$COMPX"
// direction = "up" or "down"
// change = "percent" or "value"
func Get(index, direction, change string) []MOVER {
	url := fmt.Sprintf(endpoint_movers, index)
	req, _ := http.NewRequest("GET", url, nil)
	q := req.URL.Query()
	q.Add("direction", direction)
	q.Add("change", change)
	req.URL.RawQuery = q.Encode()
	body, err := Handler(req)

	if err != nil {
		log.Fatal(err)
	}

	var movers []MOVER
	var chang, desc, dir, last, ticker, volume string
	split := strings.Split(body, "}")

	for _, x := range split {
		split2 := strings.Split(x, "\"")

		for i, x := range split2 {
			if x == "change" {
				chang = split2[i+1]
			} else if x == "description" {
				desc = split2[i+2]
			} else if x == "direction" {
				dir = split2[i+2]
			} else if x == "last" {
				last = split2[i+1]
			} else if x == "symbol" {
				ticker = split2[i+2]
			} else if x == "totalVolume" {
				volume = split2[i+1]
			}
		}

		mov := MOVER{
			TICKER:      ticker,
			DESCRIPTION: desc,
			LAST:        TrimFL(last),
			CHANGE:      TrimFL(chang),
			DIRECTION:   dir,
			VOLUME:      TrimF(volume),
		}

		movers = append(movers, mov)
	}

	return movers
}
