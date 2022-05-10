package instrument

import (
	"fmt"
	"net/http"
	"github.com/samjtro/go-tda/utils"
)

type INSTRUMENT struct {
	CUSIP		string
}

var endpoint_searchinstrument string = "https://api.tdameritrade.com/v1/instruments"
var endpoint_getinstrument string = "https://api.tdameritrade.com/v1/instruments/%s"//  		--> cusip

// Get returns a string; with information regarding the desired CUSIP,
// it takes one parameter:
// cusip = "037833100", etc.
func Get(cusip string) string {
	url := fmt.Sprintf(endpoint_getinstrument,cusip)
	req,_ := http.NewRequest("GET",url,nil)
	body := utils.Handler(req)

	return body
}

// Search uses more detailed parameters to return a string; containing more detailed information of various types on your desired ticker,
// it takes two parameters:
// ticker = "AAPL", etc.
// projection = the type of search to perform: from td-ameritrade's website:
// symbol-search: Retrieve instrument data of a specific symbol or cusip 
// symbol-regex: Retrieve instrument data for all symbols matching regex. Example: symbol=XYZ.* will return all symbols beginning with XYZ 
// desc-search: Retrieve instrument data for instruments whose description contains the word supplied. Example: symbol=FakeCompany will return all instruments with FakeCompany in the description. 
// desc-regex: Search description with full regex support. Example: symbol=XYZ.[A-C] returns all instruments whose descriptions contain a word beginning with XYZ followed by a character A through C. 
// fundamental: Returns fundamental data for a single instrument specified by exact symbol.'
func Search(ticker string, projection string) string {
	req,_ := http.NewRequest("GET",endpoint_searchinstrument,nil)
	q := req.URL.Query()
	q.Add("symbol",ticker)
	q.Add("projection",projection)
	req.URL.RawQuery = q.Encode()
	body := utils.Handler(req)

	return body
}

