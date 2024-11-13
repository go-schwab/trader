package trader

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/bytedance/sonic"
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
	Time   int `json:"datetime"`
	Volume int
	Open   float64
	Close  float64
	Hi     float64 `json:"high"`
	Lo     float64 `json:"low"`
}

type Quote struct {
	AssetMainType           string
	AssetSubType            string
	QuoteType               string
	RealTime                bool
	SSID                    int `json:"ssid"`
	Symbol                  string
	Hi52                    float64 `json:"52WeekHigh"`
	Lo52                    float64 `json:"52WeekLow"`
	AskMICId                string
	AskPrice                float64
	AskSize                 int
	AskTime                 int
	BidMICId                string
	BidPrice                float64
	BidSize                 int
	BidTime                 int
	Close                   float64 `json:"closePrice"`
	HiPrice                 float64 `json:"highPrice"`
	LastMICId               string
	LastPrice               float64
	LastSize                int
	LoPrice                 float64 `json:"lowPrice"`
	Mark                    float64 `json:"mark"`
	MarkChange              float64
	MarkPercentChange       float64
	NetChange               float64
	NetPercentChange        float64
	Open                    float64
	PostMarketChange        float64
	PostMarketPercentChange float64
	QuoteTime               int
	SecurityStatus          string
	TotalVolume             int
	TradeTime               int
}

type SimpleInstrument struct {
	Cusip       string
	Symbol      string
	Description string
	Exchange    string
	AssetType   string
}

// Change this to reflect ordering of schwab return
type FundamentalInstrument struct {
	Symbol                  string
	Cusip                   string
	Description             string
	Exchange                string
	AssetType               string
	Hi52                    float64 `json:"high52"`
	Lo52                    float64 `json:"low52"`
	DividendYield           float64
	DividendAmount          float64
	DividendDate            string
	PE                      float64 `json:"peRatio"`
	PEG                     float64 `json:"pegRatio"`
	PB                      float64 `json:"pbRatio"`
	PR                      float64 `json:"prRatio"`
	PCF                     float64 `json:"pcfRatio"`
	GrossMarginTTM          float64
	NetMarginTTM            float64
	OperatingMarginTTM      float64
	GrossMarginMRQ          float64
	NetProfitMarginMRQ      float64
	OperatingMarginMRQ      float64
	ROE                     float64 `json:"returnOnEquity"`
	ROA                     float64 `json:"returnOnAssets"`
	ROI                     float64 `json:"returnOnInvestment"`
	QuickRatio              float64
	CurrentRatio            float64
	InterestCoverage        float64
	TotalDebtToCapital      float64
	LtDebtToEquity          float64
	TotalDebtToEquity       float64
	EpsTTM                  float64
	EpsChangePercentTTM     float64
	EpsChangeYear           float64
	EpsChange               float64
	RevChangeYear           float64
	RevChangeTTM            float64
	RevChangeIn             float64
	SharesOutstanding       float64
	MarketCapFloat          float64
	MarketCap               float64
	BookValuePerShare       float64
	ShortIntToFloat         float64
	ShortIntDayToCover      float64
	DividendGrowthRate3Year float64
	DividendPayAmount       float64
	DividendPayDate         string
	Beta                    float64
	Vol1DayAvg              float64
	Vol10DayAvg             float64
	Vol3MonthAvg            float64
	Avg1DayVolume           int
	Avg10DaysVolume         int
	Avg3MonthVolume         int
	DeclarationDate         string
	DividendFreq            int
	Eps                     float64
	DtnVolume               int
	NextDividendPayDate     string
	NextDividendDate        string
	FundLeverageFactor      float64
}

type Screener struct {
	Symbol           string
	Description      string
	Volume           int
	LastPrice        float64
	NetChange        float64
	MarketShare      float64
	TotalVolume      int
	Trades           int
	NetPercentChange float64
}

/*

Legacy:
	STRIKE                 float64
	EXCHANGE               string
	EXPIRATION             float64
	DAYS2EXPIRATION        float64
	BID                    float64
	ASK                    float64
	LAST                   float64
	MARK                   float64
	BIDASK_SIZE            string
	VOLATILITY             float64
	DELTA                  float64
	GAMMA                  float64
	THETA                  float64
	VEGA                   float64
	RHO                    float64
	OPEN_INTEREST          float64
	TIME_VALUE             float64
	THEORETICAL_VALUE      float64
	THEORETICAL_VOLATILITY float64
	PERCENT_CHANGE         float64
	MARK_CHANGE            float64
	MARK_PERCENT_CHANGE    float64
	INTRINSIC_VALUE        float64
	IN_THE_MONEY           bool

*/

type Chain struct {
	Symbol           string
	Status           string
	Underlying       Underlying
	Strategy         string
	Interval         float64
	IsDelayed        bool
	IsIndex          bool
	DaysToExpiration float64
	InterestRate     float64
	UnderlyingPrice  float64
	Volatility       float64
	CallExpDateMap   map[string]map[string][]OptionContract
	PutExpDateMap    map[string]map[string][]OptionContract
}

type Underlying struct {
	Ask               float64
	AskSize           float64
	Bid               float64
	BidSize           float64
	Change            float64
	Close             float64
	Delayed           bool
	Description       string
	ExchangeName      string
	Hi52              float64 `json:"fiftyTwoWeekHigh"fiftyTwoWeekHigh"`
	Lo52              float64 `json:"fiftyTwoWeekHigh"fiftyTwoWeekLow"`
	HiPrice           float64 `json:"highPrice"`
	LoPrice           float64 `json:"lowPrice"`
	Mark              float64
	MarkPercentChange float64
	OpenPrice         float64
	PercentChange     float64
	QuoteTime         int64
	Symbol            string
	TotalVolume       int64
	TradeTime         int64
}

type OptionContract struct {
	PutCall                string
	Symbol                 string
	Description            string
	ExchangeName           string
	Bid                    float64
	Ask                    float64
	Last                   float64
	Mark                   float64
	BidSize                int64
	AskSize                int64
	LastSize               int64
	HighPrice              float64
	LowPrice               float64
	OpenPrice              float64
	ClosePrice             float64
	TotalVolume            int64
	QuoteTimeInLong        int64
	TradeTimeInLong        int64
	NetChange              float64
	Volatility             float64
	Delta                  float64
	Gamma                  float64
	Theta                  float64
	Vega                   float64
	RHO                    float64
	OpenInterest           float64
	TimeValue              float64
	InTheMoney             bool
	TheoreticalOptionValue float64
	TheoreticalVolatility  float64
	Mini                   bool
	NonStandard            bool
	OptionDeliverablesList []OptionDeliverables
	StrikePrice            float64
	ExpirationDate         string
	DaysToExpiration       int64
	ExpirationType         string
	LastTradingDay         int64
	Multiplier             float64
	SettlementType         string
	DeliverableNote        string
	PercentChange          float64
	MarkChange             float64
	MarkPercentChange      float64
	PennyPilot             bool
	IntrinsicValue         float64
	ExtrinsicValue         float64
	OptionRoot             string
	ExerciseType           string
	Hi52                   float64 `json:"high52Week"`
	Lo52                   float64 `json:"low52Week"`
}

type OptionDeliverables struct {
	Symbol           string
	AssetType        string
	DeliverableUnits float64
}

// WIP: func GetQuotes(symbols string) Quote, error) {}

// Quote returns a Quote; containing a real time quote of the desired stock's performance with a number of different indicators (including volatility, volume, price, fundamentals & more).
// It takes one parameter:
// ticker = "AAPL", etc.
func (agent *Agent) GetQuote(symbol string) (Quote, error) {
	req, err := http.NewRequest("GET", endpointQuotes, nil)
	if err != nil {
		return Quote{}, err
	}
	q := req.URL.Query()
	q.Add("symbols", symbol)
	q.Add("fields", "quote")
	req.URL.RawQuery = q.Encode()
	resp, err := agent.Handler(req)
	if err != nil {
		return Quote{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Quote{}, err
	}
	var quote Quote
	err = sonic.Unmarshal([]byte(strings.Join(strings.Split(strings.Split(string(body), fmt.Sprintf("\"%s\":", symbol))[1], "\"quote\":{"), "")[:len(strings.Join(strings.Split(strings.Split(string(body), fmt.Sprintf("\"%s\":", symbol))[1], "\"quote\":{"), ""))-2]), &quote)
	if err != nil {
		return Quote{}, err
	}
	return quote, err
}

// SearchInstrumentSimple returns instrument's simples.
// It takes one param:
func (agent *Agent) SearchInstrumentSimple(symbols string) (SimpleInstrument, error) {
	req, err := http.NewRequest("GET", endpointSearchInstrument, nil)
	if err != nil {
		return SimpleInstrument{}, err
	}
	q := req.URL.Query()
	q.Add("symbol", symbols)
	q.Add("projection", "symbol-search")
	req.URL.RawQuery = q.Encode()
	resp, err := agent.Handler(req)
	if err != nil {
		return SimpleInstrument{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return SimpleInstrument{}, err
	}
	var instrument SimpleInstrument
	err = sonic.Unmarshal([]byte(strings.Split(string(body), "[")[1][:len(strings.Split(string(body), "[")[1])-2]), &instrument)
	if err != nil {
		return SimpleInstrument{}, err
	}
	return instrument, nil
}

// SearchInstrumentFundamental returns instrument's fundamentals.
// It takes one param:
func (agent *Agent) SearchInstrumentFundamental(symbol string) (FundamentalInstrument, error) {
	req, err := http.NewRequest("GET", endpointSearchInstrument, nil)
	if err != nil {
		return FundamentalInstrument{}, err
	}
	q := req.URL.Query()
	q.Add("symbol", symbol)
	q.Add("projection", "fundamental")
	req.URL.RawQuery = q.Encode()
	resp, err := agent.Handler(req)
	if err != nil {
		return FundamentalInstrument{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return FundamentalInstrument{}, err
	}
	var instrument FundamentalInstrument
	split0 := strings.Split(string(body), "[{\"fundamental\":")[1]
	split := strings.Split(split0, "}")
	err = sonic.Unmarshal([]byte(fmt.Sprintf("%s}", strings.Join(split[:2], ""))), &instrument)
	if err != nil {
		return FundamentalInstrument{}, err
	}
	return instrument, nil
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
func (agent *Agent) GetPriceHistory(symbol, periodType, period, frequencyType, frequency, startDate, endDate string) ([]Candle, error) {
	req, err := http.NewRequest("GET", endpointPriceHistory, nil)
	if err != nil {
		return []Candle{}, err
	}
	q := req.URL.Query()
	q.Add("symbol", symbol)
	q.Add("periodType", periodType)
	q.Add("period", period)
	q.Add("frequencyType", frequencyType)
	q.Add("frequency", frequency)
	q.Add("startDate", startDate)
	q.Add("endDate", endDate)
	req.URL.RawQuery = q.Encode()
	resp, err := agent.Handler(req)
	if err != nil {
		return []Candle{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []Candle{}, err
	}
	var candles []Candle
	err = sonic.Unmarshal([]byte(fmt.Sprintf("[%s]", strings.Split(strings.Split(string(body), "[")[1], "]")[0])), &candles)
	if err != nil {
		return []Candle{}, err
	}
	return candles, nil
}

// GetMovers returns information on the desired index's movers per your desired direction and change type(percent or value),
// It takes three params:
// index = "$DJI", "$SPX.X", or "$COMPX"
// direction = "up" or "down"
// change = "percent" or "value"
func (agent *Agent) GetMovers(index, direction, change string) ([]Screener, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(endpointMovers, index), nil)
	if err != nil {
		return []Screener{}, err
	}
	q := req.URL.Query()
	q.Add("direction", direction)
	q.Add("change", change)
	req.URL.RawQuery = q.Encode()
	resp, err := agent.Handler(req)
	if err != nil {
		return []Screener{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []Screener{}, err
	}
	var movers []Screener
	stringToParse := fmt.Sprintf("[%s]", strings.Split(string(body), "[")[1][:len(strings.Split(string(body), "[")[1])-2])
	err = sonic.Unmarshal([]byte(stringToParse), &movers)
	if err != nil {
		return []Screener{}, err
	}
	return movers, nil
}

// get all option chains for a ticker
func (agent *Agent) GetChains(symbol string) (Chain, error) {
	req, err := http.NewRequest("GET", endpointOptions, nil)
	if err != nil {
		return Chain{}, err
	}
	q := req.URL.Query()
	q.Add("symbol", symbol)
	req.URL.RawQuery = q.Encode()
	resp, err := agent.Handler(req)
	if err != nil {
		return Chain{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Chain{}, err
	}
	var chain Chain
	err = sonic.Unmarshal(body, &chain)
	if err != nil {
		return Chain{}, err
	}
	return chain, nil
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
func (agent *Agent) Single(symbol, contractType, strikeRange, strikeCount, toDate string) (Chain, error) {
	req, err := http.NewRequest("GET", endpointOptions, nil)
	isErrNil(err)
	q := req.URL.Query()
	q.Add("symbol", symbol)
	q.Add("contractType", contractType)
	q.Add("range", strikeRange)
	q.Add("strikeCount", strikeCount)
	q.Add("toDate", toDate)
	req.URL.RawQuery = q.Encode()
	resp, err := agent.Handler(req)
	isErrNil(err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	isErrNil(err)
	var chain Chain
	// WIP
	err = sonic.Unmarshal(body, &chain)
	isErrNil(err)
	return chain, nil
}

/* Covered returns a string; containing covered option calls.
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

func Vertical() string {}
func Calendar() string {}
func Strangle() string {}
func Straddle() string {}
func Condor() string {}
func Diagonal() string {}
func Collar() string {}
func Roll() string {}*/
