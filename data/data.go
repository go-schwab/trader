package data

var endpoint_realtime string = "https://api.tdameritrade.com/v1/marketdata/%s/quotes"           // symbol
var endpoint_pricehistory string = "https://api.tdameritrade.com/v1/marketdata/%s/pricehistory" // symbol

// RealTime's native struct; returns various indicators related to the asset
type QUOTE struct {
	Datetime   string
	Ticker     string
	Mark       string
	Volume     string
	Volatility string
	Bid        string
	Ask        string
	Last       string
	Open       string
	Close      string
	Hi         string
	Lo         string
	Hi52       string
	Lo52       string
	PE         string
}

// this is a Go implementation of the pandas DataFrame
// slices of FRAMEs form DataFrames, which can then be used in analysis
type FRAME struct {
	Datetime string
	Volume   string
	Open     string
	Close    string
	Hi       string
	Lo       string
}
