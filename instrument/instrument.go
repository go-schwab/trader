package instrument

var endpoint_searchinstrument string = "https://api.tdameritrade.com/v1/instruments"
var endpoint_getinstrument string = "https://api.tdameritrade.com/v1/instruments/%s" // cusip

// for use with the Simple function
type SIMPLE struct {
	CUSIP       string
	TICKER      string
	DESCRIPTION string
	EXCHANGE    string
	TYPE        string
}

// for use with the Fundamental function
type FUNDAMENTAL struct {
	TICKER                 string
	CUSIP                  string
	DESCRIPTION            string
	EXCHANGE               string
	TYPE                   string
	HI52                   string
	LO52                   string
	DIV_YIELD              string
	DIV_AMOUNT             string
	PE_RATIO               string
	PEG_RATIO              string
	PB_RATIO               string
	PR_RATIO               string
	PCF_RATIO              string
	GROSS_MARGIN_TTM       string
	GROSS_MARGIN_MRQ       string
	NET_PROFIT_MARGIN_TTM  string
	NET_PROFIT_MARGIN_MRQ  string
	OPERATING_MARGIN_TTM   string
	OPERATING_MARGIN_MRQ   string
	RETURN_ON_EQUITY       string
	RETURN_ON_ASSETS       string
	RETURN_ON_INVESTMENT   string
	QUICK_RATIO            string
	CURRENT_RATIO          string
	INTEREST_COVERAGE      string
	TOTAL_DEBT_TO_CAPITAL  string
	TOTAL_DEBT_TO_EQUITY   string
	EPS_TTM                string
	EPS_CHANGE_PERCENT_TTM string
	EPS_CHANGE_YR          string
	REV_CHANGE_YR          string
	REV_CHANGE_TTM         string
	REV_CHANGE_IN          string
	SHARES_OUTSTANDING     string
	MARKET_CAP_FLOAT       string
	MARKET_CAP             string
	BOOK_VALUE_PER_SHARE   string
	BETA                   string
	VOL_1DAY               string
	VOL_10DAY              string
	VOL_3MON               string
}

// desc-regex: Search description with full regex support. Example: symbol=XYZ.[A-C] returns all instruments whose descriptions contain a word beginning with XYZ followed by a character A through C.
