package hook

import (
	"fmt"
	"os"
	"bufio"
	"net/http"
	"io/ioutil"
)

var endpoint_realtime string = "https://api.tdameritrade.com/v1/marketdata/%s/quotes"// 	 	--> symbol
var endpoint_pricehistory string = "https://api.tdameritrade.com/v1/marketdata/%s/pricehistory"// 	--> symbol
var endpoint_option string = "https://api.tdameritrade.com/v1/marketdata/chains"
var endpoint_searchinstrument string = "https://api.tdameritrade.com/v1/instruments"
var endpoint_getinstrument string = "https://api.tdameritrade.com/v1/instruments/%s"//  		--> cusip
var endpoint_movers string = "https://api.tdameritrade.com/v1/marketdata/%s/movers"// 	 		--> index
var endpoint_account string = "https://api.tdameritrade.com/v1/accounts/%s"//				--> accountID

// handling takes a *http.Request object
// it then adds the key found in .APIKEY to the parameter list
// finally, it executes the request and returns the response as a string
func handling(req *http.Request) string {
	file,_ := os.Open(".APIKEY")
	s := bufio.NewScanner(file)
	var APIKEY string;
	for s.Scan() {
		APIKEY += s.Text()
	}

	q := req.URL.Query()
	q.Add("apikey",APIKEY)
	req.URL.RawQuery = q.Encode()

	client := http.Client{}
	resp,_ := client.Do(req)
	body,_ := ioutil.ReadAll(resp.Body)

	resp.Body.Close()
	return string(body)
}

// realTime takes one parameter:
// ticker = "AAPL", etc.
func realTime(ticker string) string {
	url := fmt.Sprintf(endpoint_realtime,ticker)
	req,_ := http.NewRequest("GET",url,nil)
	body := handling(req)
	
	return body
}

// movers takes three parameters:
// index = "$DJI", "$SPX.X", or "$COMPX"
// direction = "up" or "down"
// change = "percent" or "value"
func movers(index,direction,change string) string {
	url := fmt.Sprintf(endpoint_movers,index)
	req,_ := http.NewRequest("GET",url,nil)
	q := req.URL.Query()
	q.Add("direction",direction)
	q.Add("change",change)
	req.URL.RawQuery = q.Encode()
	body := handling(req)

	return body
}

// priceHistory takes five parameters:
// ticker = "AAPL", etc.
// periodType = "day", "month", "year", "ytd" - default is "day"
// period = the number of periods to show
// frequencyType = the type of frequency with which each candle is formed; valid fTypes by pType
// "day": "minute"
// "month": "daily", "weekly"
// "year": "daily", "weekly", "monthly"
// "ytd": "daily", "weekly"
// frequency = the number of the frequencyType included in each candle; valid freqs by fType
// "minute": 1,5,10,15,30
// "daily": 1
// "weekly": 1
// "monthly": 1
func priceHistory(ticker,periodType,period,frequencyType,frequency string) string {
	url := fmt.Sprintf(endpoint_pricehistory,ticker)
	req,_ := http.NewRequest("GET",url,nil)
	q := req.URL.Query()
	q.Add("periodType",periodType)
	q.Add("period",period)
	q.Add("frequencyType",frequencyType)
	q.Add("frequency",frequency)
	req.URL.RawQuery = q.Encode()
	body := handling(req)

	return body
}

// simpleOption takes four parameters:
// ticker = "AAPL", etc.
// contractType = "CALL", "PUT", "ALL"
// strikeCount = number of strikes to return above and below the at-the-money price
// includeQuotes = "TRUE", "FALSE"
func simpleOption(ticker,contractType,strikeCount,includeQuotes string) string {
	req,_ := http.NewRequest("GET",endpoint_option,nil)
	q := req.URL.Query()
	q.Add("symbol",ticker)
	q.Add("contractType",contractType)
	q.Add("strikeCount",strikeCount)
	q.Add("includeQuotes",includeQuotes)
	req.URL.RawQuery = q.Encode()
	body := handling(req)

	return body
}

// getInstrument takes one parameter:
// cusip = "037833100", etc.
func getInstrument(cusip string) string {
	url := fmt.Sprintf(endpoint_getinstrument,cusip)
	req,_ := http.NewRequest("GET",url,nil)
	body := handling(req)

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
	body := handling(req)

	return body
}

//func strategyOption() {}
//func account() {}

func main() {
	fmt.Println(realTime("AAPL"))
}
