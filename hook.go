package main
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
var endpoint_movers string = "https://api.tdameritrade.com/v1/marketdata/%s/movers"// - 		--> index
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

func realTime(ticker string) string {
	url := fmt.Sprintf(endpoint_realtime,ticker)
	req,_ := http.NewRequest("GET",url,nil)
	body := handling(req)
	
	return body
}

func movers(index, direction, change string) string {
	url := fmt.Sprintf(endpoint_movers,index)
	req,_ := http.NewRequest("GET",url,nil)
	q := req.URL.Query()
	q.Add("direction",direction)
	q.Add("change",change)
	req.URL.RawQuery = q.Encode()
	body := handling(req)

	return body
}

func priceHistory() {

}

//func option() {}
//func instrument() {}
//func account() {}

func main() {
	req := movers("$DJI","up","percent")
	fmt.Println(req)
}

