package trader

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
