package account

import (
	"fmt"
	"net/http"
	. "github.com/samjtro/go-tda/utils"
)

/*type MARGIN struct {
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
}*/

// Get returns a string; containing account information,
// it takes one param:
// fields = this command will only return balances, but you can add positions or orders, or both (formatted positions,orders)
// func Get(fields string) string {}


