package data

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	utils "github.com/samjtro/go-schwab-traderapi/utils"
)

var (
	endpoint string = "https://api.schwabapi.com/marketdata/v1"

	// Real Time
	endpointQuote  string = endpoint + "/%s/quotes" // Symbol
	endpointQuotes string = endpoint + "/quotes"

	// Price History
	endpointPriceHistory string = endpoint + "/pricehistory"

	// Instruments
	endpointSearchInstrument string = endpoint + "/instruments"

	// Movers
	endpointMovers string = endpoint + "/movers/%s" // Index ID

	// Options
	endpointOptions string = endpoint + "/chains"
)

type Candle struct {
	Time   int     `json:"datetime"`
	Volume float64 `json:"volume"`
	Open   float64 `json:"open"`
	Close  float64 `json:"close"`
	Hi     float64 `json:"high"`
	Lo     float64 `json:"low"`
}

type Quote struct {
	AssetMainType           string  `json:"assetMainType"`
	AssetSubType            string  `json:"assetSubType"`
	QuoteType               string  `json:"quoteType"`
	RealTime                bool    `json:"realTime"`
	SSID                    int     `json:"ssid"`
	Symbol                  string  `json:"symbol"`
	Hi52                    float64 `json:"52WeekHigh"`
	Lo52                    float64 `json:"52WeekLow"`
	AskMICID                string  `json:"askMICId"`
	Ask                     float64 `json:"askPrice"`
	AskSize                 int     `json:"askSize"`
	AskTime                 int     `json:"askTime"`
	BidMICID                string  `json:"bidMICId"`
	Bid                     float64 `json:"bidPrice"`
	BidSize                 int     `json:"bidSize"`
	BidTime                 int     `json:"bidTime"`
	Close                   float64 `json:"closePrice"`
	Hi                      float64 `json:"highPrice"`
	LastMICID               string  `json:"lastMICId"`
	LastPrice               float64 `json:"lastPrice"`
	LastSize                int     `json:"lastSize"`
	Lo                      float64 `json:"lowPrice"`
	Mark                    float64 `json:"mark"`
	MarkChange              float64 `json:"markChange"`
	MarkPercentChange       float64 `json:"markPercentChange"`
	NetChange               float64 `json:"netChange"`
	NetPercentChange        float64 `json:"netPercentChange"`
	Open                    float64 `json:"open"`
	PostMarketChange        float64 `json:"postMarketChange"`
	PostMarketPercentChange float64 `json:"postMarketPercentChange"`
	QuoteTime               int     `json:"quoteTime"`
	SecurityStatus          string  `json:"securityStatus"`
	TotalVolume             int     `json:"totalVolume"`
	TradeTime               int     `json:"tradeTime"`
}

type SimpleInstrument struct {
	Cusip       string `json:"cusip"`
	Symbol      string `json:"symbol"`
	Description string `json:"description"`
	Exchange    string `json:"exchange"`
	AssetType   string `json:"assetType"`
}

// Change this to reflect ordering of schwab return
type FundamentalInstrument struct {
	Symbol                  string  `json:"symbol"`
	Cusip                   string  `json:"cusip"`
	Description             string  `json:"description"`
	Exchange                string  `json:"exchange"`
	Type                    string  `json:"assetType"`
	Hi52                    float64 `json:"high52"`
	Lo52                    float64 `json:"low52"`
	DividendYield           float64 `json:"dividendYield"`
	DividendAmount          float64 `json:"dividendAmount"`
	DividendDate            string  `json:"dividendDate"`
	PE                      float64 `json:"peRatio"`
	PEG                     float64 `json:"pegRatio"`
	PB                      float64 `json:"pbRatio"`
	PR                      float64 `json:"prRatio"`
	PCF                     float64 `json:"pcfRatio"`
	GrossMarginTTM          float64 `json:"grossMarginTTM"`
	NetProfitMarginTTM      float64 `json:"netMarginTTM"`
	OperatingMarginTTM      float64 `json:"operatingMarginTTM"`
	GrossMarginMRQ          float64 `json:"grossMarginMRQ"`
	NetProfitMarginMRQ      float64 `json:"netMarginMRQ"`
	OperatingMarginMRQ      float64 `json:"operatingMarginMRQ"`
	ROE                     float64 `json:"returnOnEquity"`
	ROA                     float64 `json:"returnOnAssets"`
	ROI                     float64 `json:"returnOnInvestment"`
	QuickRatio              float64 `json:"quickRatio"`
	CurrentRatio            float64 `json:"currentRatio"`
	InterestCoverage        float64 `json:"interestCoverage"`
	TotalDebtToCapital      float64 `json:"totalDebtToCapital"`
	LTDebtToEquity          float64 `json:"ltDebtToEquity"`
	TotalDebtToEquity       float64 `json:"totalDebtToEquity"`
	EPSTTM                  float64 `json:"epsTTM"`
	EPSChangePercentTTM     float64 `json:"epsChangePercentTTM"`
	EPSChangeYear           float64 `json:"epsChangeYear"`
	EPSChange               float64 `json:"epsChange"`
	RevenueChangeYear       float64 `json:"revChangeYear"`
	RevenueChangeTTM        float64 `json:"revChangeTTM"`
	RevenueChangeIn         float64 `json:"revChangeIn"`
	SharesOutstanding       float64 `json:"sharesOutstanding"`
	MarketCapFloat          float64 `json:"marketCapFloat"`
	MarketCap               float64 `json:"marketCap"`
	BookValuePerShare       float64 `json:"bookValuePerShare"`
	ShortIntToFloat         float64 `json:"shortIntToFloat"`
	ShortIntDayToCover      float64 `json:"shortIntDayToCover"`
	DividendGrowthRate3Year float64 `json:"dividendGrowthRate3Year"`
	DividendPayAmount       float64 `json:"dividendPayAmount"`
	DividendPayDate         string  `json:"dividendPayDate"`
	Beta                    float64 `json:"beta"`
	Vol1DayAverage          float64 `json:"vol1DayAvg"`
	Vol10DayAverage         float64 `json:"vol10DayAvg"`
	Vol3MonthAverage        float64 `json:"vol3MonthAvg"`
	Avg1DayVolume           float64 `json:"avg1DayVolume"`
	Avg10DaysVolume         float64 `json:"avg10DaysVolume"`
	Avg3MonthVolume         float64 `json:"avg3MonthVolume"`
	DeclarationDate         string  `json:"declarationDate"`
	DividendFrequency       float64 `json:"dividendFreq"`
	EPS                     float64 `json:"eps"`
	DTNVolume               float64 `json:"dtnVolume"`
	NextDividendPayDate     string  `json:"nextDividendPayDate"`
	NextDividendDate        string  `json:"nextDividendDate"`
	FundLeverageFactor      float64 `json:"fundLeverageFactor"`
}

type Screener struct {
	Symbol           string  `json:"symbol"`
	Description      string  `json:"description"`
	Volume           float64 `json:"volume"`
	LastPrice        float64 `json:"lastPrice"`
	NetChange        float64 `json:"netChange"`
	MarketShare      float64 `json:"marketShare"`
	TotalVolume      float64 `json:"totalVolume"`
	Trades           float64 `json:"trades"`
	NetPercentChange float64 `json:"netPercentChange"`
}

// WIP: type Underlying struct{}
// WIP:
type Contract struct {
	TYPE                   string  `json:""`
	SYMBOL                 string  `json:""`
	STRIKE                 float64 `json:""`
	EXCHANGE               string  `json:""`
	EXPIRATION             float64 `json:""`
	DAYS2EXPIRATION        float64 `json:""`
	BID                    float64 `json:""`
	ASK                    float64 `json:""`
	LAST                   float64 `json:""`
	MARK                   float64 `json:""`
	BIDASK_SIZE            string  `json:""`
	VOLATILITY             float64 `json:""`
	DELTA                  float64 `json:""`
	GAMMA                  float64 `json:""`
	THETA                  float64 `json:""`
	VEGA                   float64 `json:""`
	RHO                    float64 `json:""`
	OPEN_INTEREST          float64 `json:""`
	TIME_VALUE             float64 `json:""`
	THEORETICAL_VALUE      float64 `json:""`
	THEORETICAL_VOLATILITY float64 `json:""`
	PERCENT_CHANGE         float64 `json:""`
	MARK_CHANGE            float64 `json:""`
	MARK_PERCENT_CHANGE    float64 `json:""`
	INTRINSIC_VALUE        float64 `json:""`
	IN_THE_MONEY           bool    `json:""`
}

// GetCandles returns a []Candle with the previous 7 candles.
// It takes one paramter:
// ticker = "AAPL", etc.
func GetCandles(ticker string) ([]Candle, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(endpointQuote, ticker), nil)
	utils.Check(err)
	body, err := utils.Handler(req)
	utils.Check(err)
	var candles []Candle
	fmt.Println(body)
	/*err = json.Unmarshal([]byte(body), &candles)
	utils.Check(err)*/
	return candles, nil
}

// PriceHistory returns a series of candles with price volume & datetime info per candlestick.
// It takes seven parameters:
// ticker = "AAPL", etc.;
// periodType = "day", "month", "year", "ytd" - default is "day";
// period = the number of periods to show;
// frequencyType = the type of frequency with which each candle is formed; valid fTypes by pType;
// "day": "minute" /
// "month": "daily", "weekly" /
// "year": "daily", "weekly", "monthly" /
// "ytd": "daily", "weekly";
// frequency = the number of the frequencyType included in each candle; valid freqs by fType
// "minute": 1,5,10,15,30 /
// "daily": 1 /
// "weekly": 1 /
// "monthly": 1;
// startDate =
// endDate =
func GetPriceHistory(symbol, periodType, period, frequencyType, frequency, startDate, endDate string) ([]Candle, error) {
	req, err := http.NewRequest("GET", endpointPriceHistory, nil)
	utils.Check(err)
	q := req.URL.Query()
	q.Add("symbol", symbol)
	q.Add("periodType", periodType)
	q.Add("period", period)
	q.Add("frequencyType", frequencyType)
	q.Add("frequency", frequency)
	q.Add("startDate", startDate)
	q.Add("endDate", endDate)
	req.URL.RawQuery = q.Encode()
	body, err := utils.Handler(req)
	utils.Check(err)
	var candles []Candle
	err = json.Unmarshal([]byte(strings.Split(strings.Split(body, "[")[1], "]")[0]), &candles)
	utils.Check(err)
	return candles, nil
}

// WIP: func GetQuotes(symbols string) Quote, error) {}

// Quote returns a Quote; containing a real time quote of the desired stock's performance with a number of different indicators (including volatility, volume, price, fundamentals & more).
// It takes one parameter:
// ticker = "AAPL", etc.
func GetQuote(symbol string) (Quote, error) {
	req, err := http.NewRequest("GET", endpointQuotes, nil)
	utils.Check(err)
	q := req.URL.Query()
	q.Add("symbols", symbol)
	q.Add("fields", "quote")
	req.URL.RawQuery = q.Encode()
	body, err := utils.Handler(req)
	utils.Check(err)
	var quote Quote
	err = json.Unmarshal([]byte(strings.Join(strings.Split(strings.Split(body, fmt.Sprintf("\"%s\":", symbol))[1], "\"quote\":{"), "")[:len(strings.Join(strings.Split(strings.Split(body, fmt.Sprintf("\"%s\":", symbol))[1], "\"quote\":{"), ""))-2]), &quote)
	utils.Check(err)
	return quote, err
}

// SearchInstrumentSimple returns instrument's simples.
// It takes one param:
func SearchInstrumentSimple(symbols string) (SimpleInstrument, error) {
	req, err := http.NewRequest("GET", endpointSearchInstrument, nil)
	utils.Check(err)
	q := req.URL.Query()
	q.Add("symbol", symbols)
	q.Add("projection", "symbol-search")
	req.URL.RawQuery = q.Encode()
	body, err := utils.Handler(req)
	utils.Check(err)
	var instrument SimpleInstrument
	err = json.Unmarshal([]byte(strings.Split(body, "[")[1][:len(strings.Split(body, "[")[1])-2]), &instrument)
	utils.Check(err)
	return instrument, nil
}

// SearchInstrumentFundamental returns instrument's fundamentals.
// It takes one param:
func SearchInstrumentFundamental(symbol string) (FundamentalInstrument, error) {
	req, err := http.NewRequest("GET", endpointSearchInstrument, nil)
	utils.Check(err)
	q := req.URL.Query()
	q.Add("symbol", symbol)
	q.Add("projection", "fundamental")
	req.URL.RawQuery = q.Encode()
	body, err := utils.Handler(req)
	utils.Check(err)
	var instrument FundamentalInstrument
	split0 := strings.Split(body, "[{\"fundamental\":")[1]
	split := strings.Split(split0, "}")
	err = json.Unmarshal([]byte(fmt.Sprintf("%s}", strings.Join(split[:2], ""))), &instrument)
	utils.Check(err)
	return instrument, nil
}

// GetMovers returns information on the desired index's movers per your desired direction and change type(percent or value),
// It takes three params:
// index = "$DJI", "$SPX.X", or "$COMPX"
// direction = "up" or "down"
// change = "percent" or "value"
func GetMovers(index, direction, change string) ([]Screener, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(endpointMovers, index), nil)
	utils.Check(err)
	q := req.URL.Query()
	q.Add("direction", direction)
	q.Add("change", change)
	req.URL.RawQuery = q.Encode()
	body, err := utils.Handler(req)
	utils.Check(err)
	var movers []Screener
	stringToParse := fmt.Sprintf("[%s]", strings.Split(body, "[")[1][:len(strings.Split(body, "[")[1])-2])
	err = json.Unmarshal([]byte(stringToParse), &movers)
	utils.Check(err)
	return movers, nil
}

// Single returns a []CONTRACT; containing a SINGLE option chain of your desired strike, type, etc.,
// it takes four parameters:
// ticker = "AAPL", etc.
// contractType = "CALL", "PUT", "ALL";
// strikeRange = returns option chains for a given range:
// ITM = in da money
// NTM = near da money
// OTM = out da money
// SAK = strikes above market
// SBK = strikes below market
// SNK = strikes near market
// ALL* = default, all strikes;
// strikeCount = The number of strikes to return above and below the at-the-money price;
// toDate = Only return expirations before this date. Valid ISO-8601 formats are: yyyy-MM-dd and yyyy-MM-dd'T'HH:mm:ssz.
// Lets examine a sample call of Single: Single("AAPL","CALL","ALL","5","2022-07-01").
// This returns 5 AAPL CALL contracts both above and below the at the money price, with no preference as to the status of the contract ("ALL"), expiring before 2022-07-01
func Single(ticker, contractType, strikeRange, strikeCount, toDate string) ([]Contract, error) {
	req, err := http.NewRequest("GET", endpointOptions, nil)
	utils.Check(err)
	q := req.URL.Query()
	q.Add("symbol", ticker)
	q.Add("contractType", contractType)
	q.Add("range", strikeRange)
	q.Add("strikeCount", strikeCount)
	q.Add("toDate", toDate)
	req.URL.RawQuery = q.Encode()
	body, err := utils.Handler(req)
	utils.Check(err)
	var chain []Contract
	// WIP
	err = json.Unmarshal([]byte(body), &chain)
	utils.Check(err)
	return chain, nil
}

// Covered returns a string; containing covered option calls.
// Not functional ATM.
func Covered(ticker, contractType, strikeRange, strikeCount, toDate string) (string, error) {
	req, _ := http.NewRequest("GET", endpointOptions, nil)
	q := req.URL.Query()
	q.Add("strategy", "COVERED")
	q.Add("symbol", ticker)
	q.Add("contractType", contractType)
	q.Add("range", strikeRange)
	q.Add("strikeCount", strikeCount)
	q.Add("toDate", toDate)
	body, err := utils.Handler(req)
	utils.Check(err)

	return body, nil
}

// Butterfly returns a string; containing Butterfly spread option calls.
// Not functional ATM.
func Butterfly(ticker, contractType, strikeRange, strikeCount, toDate string) (string, error) {
	req, _ := http.NewRequest("GET", endpointOptions, nil)
	q := req.URL.Query()
	q.Add("strategy", "BUTTERFLY")
	q.Add("symbol", ticker)
	q.Add("contractType", contractType)
	q.Add("range", strikeRange)
	q.Add("strikeCount", strikeCount)
	q.Add("toDate", toDate)
	body, err := utils.Handler(req)
	utils.Check(err)

	return body, nil
}

// ANALYTICAL returns a string; allows you to control additional parameters for theoretical value calculations:
// It takes nine parameters:
// Not functional ATM.
func Analytical(ticker, contractType, strikeRange, strikeCount, toDate, volatility, underlyingPrice, interestRate, daysToExpiration string) (string, error) {
	req, _ := http.NewRequest("GET", endpointOptions, nil)
	q := req.URL.Query()
	q.Add("strategy", "ANALYTICAL")
	q.Add("symbol", ticker)
	q.Add("contractType", contractType)
	q.Add("range", strikeRange)
	q.Add("strikeCount", strikeCount)
	q.Add("toDate", toDate)
	q.Add("volatility", volatility)
	q.Add("underlyingPrice", underlyingPrice)
	q.Add("interestRate", interestRate)
	q.Add("daysToExpiration", underlyingPrice)
	req.URL.RawQuery = q.Encode()
	body, err := utils.Handler(req)
	utils.Check(err)

	return body, nil
}

// func Vertical() string {}
// func Calendar() string {}
// func Strangle() string {}
// func Straddle() string {}
// func Condor() string {}
// func Diagonal() string {}
// func Collar() string {}
// func Roll() string {}
