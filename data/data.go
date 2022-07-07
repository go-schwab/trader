package data

var endpoint_realtime string = "https://api.tdameritrade.com/v1/marketdata/%s/quotes"           // symbol
var endpoint_pricehistory string = "https://api.tdameritrade.com/v1/marketdata/%s/pricehistory" // symbol

// RealTime's native struct; returns various indicators related to the asset
type QUOTE struct {
	DATETIME   string
	TICKER     string
	MARK       string
	VOLUME     string
	VOLATILITY string
	BID        string
	ASK        string
	LAST       string
	OPEN       string
	CLOSE      string
	HI         string
	LO         string
	HI52       string
	LO52       string
	PE_RATIO   string
}

// this is a Go implementation of the pandas DataFrame
// slices of FRAMEs form DataFrames, which can then be used in analysis
type FRAME struct {
	DATETIME string
	VOLUME   string
	OPEN     string
	CLOSE    string
	HI       string
	LO       string
}
