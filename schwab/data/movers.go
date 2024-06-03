package data

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	req, err := http.NewRequest("GET", fmt.Sprintf(Endpoint_movers, index), nil)
	utils.Check(err)
	q := req.URL.Query()
	q.Add("direction", direction)
	q.Add("change", change)
	req.URL.RawQuery = q.Encode()
	body, err := schwabutils.Handler(req)
	utils.Check(err)
	var movers []Screener
	stringToParse := fmt.Sprintf("[%s]", strings.Split(body, "[")[1][:len(strings.Split(body, "[")[1])-2])
	err = json.Unmarshal([]byte(stringToParse), &movers)
	utils.Check(err)
	return movers, nil
}
