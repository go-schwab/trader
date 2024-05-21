package schwab 

var (
	// endpoint_quotelist
	endpoint_quote string = fmt.Sprintf(endpoint + "/%s/quotes") // Symbol
	endpoint_pricehistory string = fmt.Sprintf(endpoint + "/%s/pricehistory") // Symbol
)

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

// This is a Go implementation of the pandas "DataFrame" structure
// Slices of FRAMEs form DataFrames, which can then be used in analysis
type FRAME struct {
	Datetime string
	Volume   float64
	Open     float64
	Close    float64
	Hi       float64
	Lo       float64
}
