package data

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	err = json.Unmarshal([]byte(body), &movers)
	utils.Check(err)
	return movers, nil
}
