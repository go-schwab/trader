package data

var endpoint_realtime string = "https://api.tdameritrade.com/v1/marketdata/%s/quotes"           // symbol
var endpoint_pricehistory string = "https://api.tdameritrade.com/v1/marketdata/%s/pricehistory" // symbol

// RealTime's native struct; returns various indicators related to the asset
type QUOTE struct {
	Datetime   string
	Ticker     string
	Mark       float64
	Volume     float64
	Volatility float64
	Bid        float64
	Ask        float64
	Last       float64
	Open       float64
	Close      float64
	Hi         float64
	Lo         float64
	Hi52       float64
	Lo52       float64
	PE         float64
}

// this is a Go implementation of the pandas DataFrame
// slices of FRAMEs form DataFrames, which can then be used in analysis
type FRAME struct {
	Datetime string
	Volume   float64
	Open     float64
	Close    float64
	Hi       float64
	Lo       float64
}
