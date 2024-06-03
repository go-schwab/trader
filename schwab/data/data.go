package data

var (
	Endpoint string = "https://api.schwabapi.com/marketdata/v1"

	// Real Time
	Endpoint_quote  string = Endpoint + "/%s/quotes" // Symbol
	Endpoint_quotes string = Endpoint + "/quotes"

	// Price History
	Endpoint_pricehistory string = Endpoint + "/%s/pricehistory" // Symbol

	// Instruments
	Endpoint_searchinstruments string = Endpoint + "/instruments"

	// Movers
	Endpoint_movers string = Endpoint + "/movers/%s" // Index ID

	// Options
	Endpoint_option string = Endpoint + "/chains"
)

type Candle struct {
	Time   string
	Volume float64
	Open   float64
	Close  float64
	Hi     float64
	Lo     float64
}

type Quote struct {
	Time       string
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

type FundamentalInstrument struct {
	Symbol                  string
	Cusip                   string
	Description             string
	Exchange                string
	Type                    string
	Hi52                    float64
	Lo52                    float64
	DivYield                float64
	DivAmount               float64
	DividendDate            string
	PE                      float64
	PEG                     float64
	PB                      float64
	PR                      float64
	PCF                     float64
	GrossMarginTTM          float64
	NetProfitMarginTTM      float64
	OperatingMarginTTM      float64
	GrossMarginMRQ          float64
	NetProfitMarginMRQ      float64
	OperatingMarginMRQ      float64
	ROE                     float64
	ROA                     float64
	ROI                     float64
	QuickRatio              float64
	CurrentRatio            float64
	InterestCoverage        float64
	TotalDebtToCapital      float64
	LTDebtToEquity          float64
	TotalDebtToEquity       float64
	EPSTTM                  float64
	EPSChangePercentTTM     float64
	EPSChangeYear           float64
	EPSChange               float64
	RevenueChangeYear       float64
	RevenueChangeTTM        float64
	RevenueChangeIn         float64
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
	Vol1DayAverage          float64
	Vol10DayAverage         float64
	Vol3MonthAverage        float64
	Avg1DayVolume           float64
	Avg10DaysVolume         float64
	Avg3MonthVolume         float64
	DeclarationDate         string
	DividendFrequency       float64
	EPS                     float64
	DTNVolume               float64
	NextDividendPayDate     string
	NextDividendDate        string
	FundLeverageFactor      float64
}

type SimpleInstrument struct {
	Cusip       string
	Symbol      string
	Description string
	Exchange    string
	AssetType   string
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

// WIP: Options
type Underlying struct {
}

type Contract struct {
	TYPE                   string
	SYMBOL                 string
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
	IN_THE_MONEY           bool //bool
}
