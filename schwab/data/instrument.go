package data

// Search for an instrument by a specific symbol, return all fundamental asset information.
func SearchInstrument(symbol string) (INSTRUMENT, error) {
	/*req, _ := http.NewRequest("GET", Endpoint_searchinstruments, nil)
	q := req.URL.Query()
	q.Add("symbol", symbol)
	q.Add("projection", "fundamental")
	req.URL.RawQuery = q.Encode()
	body, err := schwabutils.Handler(req) // schwabutils "github.com/samjtro/go-trade/schwab/utils"

	if err != nil {
		return INSTRUMENT{}, err
	}

	split := strings.Split(body, "{")*/

	return INSTRUMENT{}, nil
}
