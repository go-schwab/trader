package data

var endpoint_realtime string = "https://api.tdameritrade.com/v1/marketdata/%s/quotes"           // symbol
var endpoint_pricehistory string = "https://api.tdameritrade.com/v1/marketdata/%s/pricehistory" // symbol

// for use with RealTime
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

// for use with PriceHistory
type FRAME struct {
	DATETIME string
	VOLUME   string
	OPEN     string
	CLOSE    string
	HI       string
	LO       string
}
