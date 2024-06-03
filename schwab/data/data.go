package data

var (
	Endpoint string = "https://api.schwabapi.com/marketdata/v1"

	// Real Time
	Endpoint_quote  string = Endpoint + "/%s/quotes" // Symbol
	Endpoint_quotes string = Endpoint + "/quotes"

	// Price History
	Endpoint_priceHistory string = Endpoint + "/%s/pricehistory" // Symbol

	// Instruments
	Endpoint_searchInstruments string = Endpoint + "/instruments"
	Endpoint_searchInstrument  string = Endpoint_searchInstruments + "/%s" // Cusip

	// Movers
	Endpoint_movers string = Endpoint + "/movers/%s" // Index ID

	// Options
	Endpoint_option string = Endpoint + "/chains"
)

type Candle struct {
	Time   string  `json:"datetime"`
	Volume float64 `json:"volume"`
	Open   float64 `json:"open"`
	Close  float64 `json:"close"`
	Hi     float64 `json:"high"`
	Lo     float64 `json:"low"`
}

// WIP
type Quote struct {
	Time       string  `json:""`
	Ticker     string  `json:""`
	Mark       float64 `json:""`
	Volume     float64 `json:""`
	Volatility float64 `json:""`
	Bid        float64 `json:""`
	Ask        float64 `json:""`
	Last       float64 `json:""`
	Open       float64 `json:""`
	Close      float64 `json:""`
	Hi         float64 `json:""`
	Lo         float64 `json:""`
	Hi52       float64 `json:""`
	Lo52       float64 `json:""`
	PE         float64 `json:""`
}

type FundamentalInstrument struct {
	Symbol                  string  `json:"synbol"`
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

type SimpleInstrument struct {
	Cusip       string `json:"cusip"`
	Symbol      string `json:"symbol"`
	Description string `json:"description"`
	Exchange    string `json:"exchange"`
	AssetType   string `json:"assetType"`
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
type Underlying struct{}

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
