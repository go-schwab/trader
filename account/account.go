package account

import (
	"fmt"
	"net/http"
)

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

//func Get(string fields) string {}


