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

type CANDLE struct {
	Time   string
	Volume float64
	Open   float64
	Close  float64
	Hi     float64
	Lo     float64
}

type QUOTE struct {
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

type INSTRUMENT struct {
	Cusip       string
	Symbol      string
	Description string
	Exchange    string
	AssetType   string
}

type MOVER struct {
	Symbol           string
	Description      string
	Volume           float64
	LastPrice        float64
	NetChange        float64
	MarketShare      float64
	TotalVolume      float64
	Trades           float64
	NetPercentChange float64
}

// WIP: Options
type UNDERLYING struct{}

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
