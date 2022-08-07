package instrument

var endpoint_searchinstrument string = "https://api.tdameritrade.com/v1/instruments"
var endpoint_getinstrument string = "https://api.tdameritrade.com/v1/instruments/%s" // cusip

// Returns simple information regarding the assets
type SIMPLE struct {
	CUSIP       string
	TICKER      string
	DESCRIPTION string
	EXCHANGE    string
	TYPE        string
}

// Returns all fundamentals of the given asset
type FUNDAMENTAL struct {
	TICKER                 string
	CUSIP                  string
	DESCRIPTION            string
	EXCHANGE               string
	TYPE                   string
	HI52                   float64
	LO52                   float64
	DIV_YIELD              float64
	DIV_AMOUNT             float64
	PE_RATIO               float64
	PEG_RATIO              float64
	PB_RATIO               float64
	PR_RATIO               float64
	PCF_RATIO              float64
	GROSS_MARGIN_TTM       float64
	GROSS_MARGIN_MRQ       float64
	NET_PROFIT_MARGIN_TTM  float64
	NET_PROFIT_MARGIN_MRQ  float64
	OPERATING_MARGIN_TTM   float64
	OPERATING_MARGIN_MRQ   float64
	RETURN_ON_EQUITY       float64
	RETURN_ON_ASSETS       float64
	RETURN_ON_INVESTMENT   float64
	QUICK_RATIO            float64
	CURRENT_RATIO          float64
	INTEREST_COVERAGE      float64
	TOTAL_DEBT_TO_CAPITAL  float64
	TOTAL_DEBT_TO_EQUITY   float64
	EPS_TTM                float64
	EPS_CHANGE_PERCENT_TTM float64
	EPS_CHANGE_YR          float64
	REV_CHANGE_YR          float64
	REV_CHANGE_TTM         float64
	REV_CHANGE_IN          float64
	SHARES_OUTSTANDING     float64
	MARKET_CAP_FLOAT       float64
	MARKET_CAP             float64
	BOOK_VALUE_PER_SHARE   float64
	BETA                   float64
	VOL_1DAY               float64
	VOL_10DAY              float64
	VOL_3MON               float64
}

// desc-regex: Search description with full regex support. Example: symbol=XYZ.[A-C] returns all instruments whose descriptions contain a word beginning with XYZ followed by a character A through C.
