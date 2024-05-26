package data

import "fmt"

var (
	Endpoint string = "https://api.schwabapi.com/marketdata/v1"

	// Real Time
	Endpoint_quote  string = Endpoint + "/%s/quotes" // Symbol
	Endpoint_quotes string = Endpoint + "/quotes"

	// Price History
	Endpoint_pricehistory string = fmt.Sprintf(Endpoint + "/%s/pricehistory") // Symbol

	// Instruments
	Endpoint_searchinstruments string = fmt.Sprintf(Endpoint + "/instruments")

	// Movers
	Endpoint_movers string = fmt.Sprintf(Endpoint + "/movers/%s") // Index ID

	// Options
	Endpoint_option string = fmt.Sprintf(Endpoint + "/chains")
)

// Real Time, Price History
type CANDLE struct {
	Datetime string
	Volume   float64
	Open     float64
	Close    float64
	Hi       float64
	Lo       float64
}

// Real Time
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

// Instruments
type INSTRUMENT struct {
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

// Movers
type MOVER struct {
	TICKER      string
	DESCRIPTION string
	LAST        float64
	VOLUME      float64
	DIRECTION   string
	CHANGE      float64
}

// Options: WIP
// type UNDERLYING struct {}

type CONTRACT struct {
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
