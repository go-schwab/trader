package schwab

import (
	"fmt"
	"net/http"

	"github.com/samjtro/go-trade/utils"
)

// As of 2024, these TDA endpoints are still functional. Working on migrating to Schwab before v1.0.0.

type MARGIN struct {
	Type			string
	accountId		string
	roundTrips		string
	isDayTrader		bool
	positions		[]map(string,int)
	orderStrategies 	[]map(string,string)
	initialBalances 	[]map(string,int)
	currentBalances 	[]map(string,int)
	projectedBalances	[]map(string,int)
}

type CASH struct {
	Type			string
	accountId		string
	roundTrips		string
	isDayTrader		bool
	positions		[]map(string,int)
	orderStrategies 	[]map(string,string)
	initialBalances 	[]map(string,int)
	currentBalances 	[]map(string,int)
	projectedBalances	[]map(string,int)
}

var endpoint_account string = "https://api.tdameritrade.com/v1/accounts/%s" // accountID
var endpoint_place = "https://api.tdameritrade.com/v1/accounts/%s/orders"
var endpoint_watchlist = "https://api.tdameritrade.com/v1/accounts/%s/watchlists"
var endpoint_bearer = "https://api.tdameritrade.com/v1/oauth2/token"

// Get returns a string; containing account information,
// it takes three params:
// accountID = your accountID
// fields = this command will only return balances, but you can add positions or orders, or both - "positions,orders"
// Bearer = Bearer token for your account, generated from https://developer.tdameritrade.com/authentication/apis/post/token-0
func Get(accountID, fields, Bearer string) (string, error) {
	url := fmt.Sprintf(endpoint_account, accountID)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", Bearer)
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	body, err := utils.Handler(req)

	if err != nil {
		return "", err
	}

	return body, err
}
