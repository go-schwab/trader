package instruments

package (
	"fmt"
	"net/http"
)

// getInstrument takes one parameter:
// cusip = "037833100", etc.
func getInstrument(cusip string) string {
	url := fmt.Sprintf(endpoint_getinstrument,cusip)
	req,_ := http.NewRequest("GET",url,nil)
	body := handler(req)

	return body
}

// searchInstrument takes two parameters:
// ticker = "AAPL", etc.
// projection = the type of search to perform: from td-ameritrade's website:
// symbol-search: Retrieve instrument data of a specific symbol or cusip 
// symbol-regex: Retrieve instrument data for all symbols matching regex. Example: symbol=XYZ.* will return all symbols beginning with XYZ 
// desc-search: Retrieve instrument data for instruments whose description contains the word supplied. Example: symbol=FakeCompany will return all instruments with FakeCompany in the description. 
// desc-regex: Search description with full regex support. Example: symbol=XYZ.[A-C] returns all instruments whose descriptions contain a word beginning with XYZ followed by a character A through C. 
// fundamental: Returns fundamental data for a single instrument specified by exact symbol.'
func searchInstrument(ticker string, projection string) string {
	req,_ := http.NewRequest("GET",endpoint_searchinstrument,nil)
	q := req.URL.Query()
	q.Add("symbol",ticker)
	q.Add("projection",projection)
	req.URL.RawQuery = q.Encode()
	body := handler(req)

	return body
}

