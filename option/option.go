package option

import (
	"log"
	"net/http"
	"strings"

	. "github.com/samjtro/go-tda/utils"
)

// Anatomy of the TDA option return:
//{"symbol":"AAPL","status":"SUCCESS","underlying":null,"strategy":"SINGLE","interval":0.0,"isDelayed":true,"isIndex":false,"interestRate":0.1,"underlyingPrice":152.625,"volatility":29.0,"daysToExpiration":0.0,"numberOfContracts":270,"putExpDateMap":{},"callExpDateMap":{"2022-05-13:2":{
//"145.0":[{"putCall":"CALL","symbol":"AAPL_051322C145","description":"AAPL May 13 2022 145 Call (Weekly)","exchangeName":"OPR","bid":9.5,"ask":10.95,"last":9.8,"mark":10.23,"bidSize":50,"askSize":51,"bidAskSize":"50X51","lastSize":0,"highPrice":0.0,"lowPrice":0.0,"openPrice":0.0,"closePrice":9.87,"totalVolume":0,"tradeDate":null,"tradeTimeInLong":1652212791092,"quoteTimeInLong":1652212798993,"netChange":-0.07,"volatility":49.752,"delta":0.903,"gamma":0.022,"theta":-0.181,"vega":0.027,"rho":0.013,"openInterest":907,"timeValue":0.29,"theoreticalOptionValue":9.874,"theoreticalVolatility":29.0,"optionDeliverablesList":null,"strikePrice":145.0,"expirationDate":1652472000000,"daysToExpiration":2,"expirationType":"S","lastTradingDay":1652486400000,"multiplier":100.0,"settlementType":" ","deliverableNote":"","isIndexOption":null,"percentChange":-0.75,"markChange":0.35,"markPercentChange":3.55,"intrinsicValue":9.51,"inTheMoney":true,"mini":false,"pennyPilot":true,"nonStandard":false}],
//"146.0":[{"putCall":"CALL","symbol":"AAPL_051322C146","description":"AAPL May 13 2022 146 Call(Weekly)","exchangeName":"OPR","bid":8.85,"ask":9.75,"last":10.35,"mark":9.3,"bidSize":21,"askSize":50,"bidAskSize":"21X50","lastSize":0,"highPrice":0.0,"lowPrice":0.0,"openPrice":0.0,"closePrice":8.98,"totalVolume":0,"tradeDate":null,"tradeTimeInLong":1652209972509,"quoteTimeInLong":1652212799988,"netChange":1.37,"volatility":49.456,"delta":0.878,"gamma":0.026,"theta":-0.212,"vega":0.031,"rho":0.013,"openInterest":320,"timeValue":1.84,"theoreticalOptionValue":8.985,"theoreticalVolatility":29.0,"optionDeliverablesList":null,"strikePrice":146.0,"expirationDate":1652472000000,"daysToExpiration":2,"expirationType":"S","lastTradingDay":1652486400000,"multiplier":100.0,"settlementType":" ","deliverableNote":"","isIndexOption":null,"percentChange":15.2,"markChange":0.32,"markPercentChange":3.52,"intrinsicValue":8.51,"inTheMoney":true,"mini":false,"pennyPilot":true,"nonStandard":false}],

var endpoint_option string = "https://api.tdameritrade.com/v1/marketdata/chains"

//type UNDERLYING struct {}

type CONTRACT struct {
	TYPE                   string
	SYMBOL                 string
	STRIKE                 string
	EXCHANGE               string
	EXPIRATION             string
	DAYS2EXPIRATION        string
	BID                    string
	ASK                    string
	LAST                   string
	MARK                   string
	BIDASK_SIZE            string
	VOLATILITY             string
	DELTA                  string
	GAMMA                  string
	THETA                  string
	VEGA                   string
	RHO                    string
	OPEN_INTEREST          string
	TIME_VALUE             string
	THEORETICAL_VALUE      string
	THEORETICAL_VOLATILITY string
	PERCENT_CHANGE         string
	MARK_CHANGE            string
	MARK_PERCENT_CHANGE    string
	INTRINSIC_VALUE        string
	IN_THE_MONEY           string //bool
}

// Single returns a []CONTRACT; containing a SINGLE option chain of your desired strike, type, etc.,
// it takes four parameters:
// ticker = "AAPL", etc.
// contractType = "CALL", "PUT", "ALL"
// strikeRange = returns option chains for a given range:
// ITM = in da money
// NTM = near da money
// OTM = out da money
// SAK = strikes above market
// SBK = strikes below market
// SNK = strikes near market
// ALL* = default, all strikes
// strikeCount = The number of strikes to return above and below the at-the-money price
// toDate = Only return expirations before this date. Valid ISO-8601 formats are: yyyy-MM-dd and yyyy-MM-dd'T'HH:mm:ssz.
// lets examine a sample call of Single: Single("AAPL","CALL","ALL","5","2022-07-01")
// this returns 5 AAPL CALL contracts both above and below the at the money price, with no preference as to the
// status of the contract ("ALL"), expiring before 2022-07-01
func Single(ticker, contractType, strikeRange, strikeCount, toDate string) []CONTRACT {
	req, _ := http.NewRequest("GET", endpoint_option, nil)
	q := req.URL.Query()
	q.Add("symbol", ticker)
	q.Add("contractType", contractType)
	q.Add("range", strikeRange)
	q.Add("strikeCount", strikeCount)
	q.Add("toDate", toDate)
	req.URL.RawQuery = q.Encode()
	body, err := Handler(req)

	if err != nil {
		log.Fatal(err)
	}

	var chain []CONTRACT
	var Type, symbol, exchange, strikePrice, exp, d2e, bid, ask, last, mark, bidAskSize, volatility, delta, gamma, theta, vega, rho, openInterest, timeValue, theoreticalValue, theoreticalVolatility, percentChange, markChange, markPercentChange, intrinsicValue, inTheMoney string
	split := strings.Split(body, "}],")

	for _, x := range split {
		split2 := strings.Split(x, "\"")

		for i, x := range split2 {
			switch x {
			case "putCall":
				Type = split2[i+2]
			case "symbol":
				symbol = split2[i+2]
			case "exchangeName":
				exchange = split2[i+2]
			case "strikePrice":
				strikePrice = split2[i+1]
			case "expirationDate":
				exp = split2[i+1]
			case "daysToExpiration":
				d2e = split2[i+1]
			case "bid":
				bid = split2[i+1]
			case "ask":
				ask = split2[i+1]
			case "last":
				last = split2[i+1]
			case "mark":
				mark = split2[i+1]
			case "bidAskSize":
				bidAskSize = split2[i+2]
			case "volatility":
				volatility = split2[i+1]
			case "delta":
				delta = split2[i+1]
			case "gamma":
				gamma = split2[i+1]
			case "theta":
				theta = split2[i+1]
			case "vega":
				vega = split2[i+1]
			case "rho":
				rho = split2[i+1]
			case "openInterest":
				openInterest = split2[i+1]
			case "timeValue":
				timeValue = split2[i+1]
			case "theoreticalOptionValue":
				theoreticalValue = split2[i+1]
			case "theoreticalVolatility":
				theoreticalVolatility = split2[i+1]
			case "percentChange":
				percentChange = split2[i+1]
			case "markChange":
				markChange = split2[i+1]
			case "markPercentChange":
				markPercentChange = split2[i+1]
			case "intrinsicValue":
				intrinsicValue = split2[i+1]
			case "inTheMoney":
				inTheMoney = split2[i+1]
			}
		}

		contract := CONTRACT{
			TYPE:                   Type,
			SYMBOL:                 symbol,
			STRIKE:                 TrimFL(strikePrice),
			EXCHANGE:               exchange,
			EXPIRATION:             TrimFL(exp),
			DAYS2EXPIRATION:        TrimFL(d2e),
			BID:                    TrimFL(bid),
			ASK:                    TrimFL(ask),
			LAST:                   TrimFL(last),
			MARK:                   TrimFL(mark),
			BIDASK_SIZE:            bidAskSize,
			VOLATILITY:             TrimFL(volatility),
			DELTA:                  TrimFL(delta),
			GAMMA:                  TrimFL(gamma),
			THETA:                  TrimFL(theta),
			VEGA:                   TrimFL(vega),
			RHO:                    TrimFL(rho),
			OPEN_INTEREST:          TrimFL(openInterest),
			TIME_VALUE:             TrimFL(timeValue),
			THEORETICAL_VALUE:      TrimFL(theoreticalValue),
			THEORETICAL_VOLATILITY: TrimFL(theoreticalVolatility),
			PERCENT_CHANGE:         TrimFL(percentChange),
			MARK_CHANGE:            TrimFL(markChange),
			MARK_PERCENT_CHANGE:    TrimFL(markPercentChange),
			INTRINSIC_VALUE:        TrimFL(intrinsicValue),
			IN_THE_MONEY:           TrimFL(inTheMoney),
		}

		chain = append(chain, contract)
	}

	return chain
}

// ANALYTICAL returns a string; allows you to control additional parameters for theoretical value calculations;
// it takes nine parameters
func Analytical(ticker, contractType, strikeRange, strikeCount, toDate, volatility, underlyingPrice, interestRate, daysToExpiration string) string {
	req, _ := http.NewRequest("GET", endpoint_option, nil)
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
	body, err := Handler(req)

	if err != nil {
		log.Fatal(err)
	}

	return body
}

// Covered returns a string; containing covered option calls
func Covered(ticker, contractType, strikeRange, strikeCount, toDate string) string {
	req, _ := http.NewRequest("GET", endpoint_option, nil)
	q := req.URL.Query()
	q.Add("strategy", "COVERED")
	q.Add("symbol", ticker)
	q.Add("contractType", contractType)
	q.Add("range", strikeRange)
	q.Add("strikeCount", strikeCount)
	q.Add("toDate", toDate)
	body, err := Handler(req)

	if err != nil {
		log.Fatal(err)
	}

	return body
}

// Butterfly returns a string; containing Butterfly spread option calls
func Butterfly(ticker, contractType, strikeRange, strikeCount, toDate string) string {
	req, _ := http.NewRequest("GET", endpoint_option, nil)
	q := req.URL.Query()
	q.Add("strategy", "BUTTERFLY")
	q.Add("symbol", ticker)
	q.Add("contractType", contractType)
	q.Add("range", strikeRange)
	q.Add("strikeCount", strikeCount)
	q.Add("toDate", toDate)
	body, err := Handler(req)

	if err != nil {
		log.Fatal(err)
	}

	return body
}

// func Vertical() string {}
// func Calendar() string {}
// func Strangle() string {}
// func Straddle() string {}
// func Condor() string {}
// func Diagonal() string {}
// func Collar() string {}
// func Roll() string {}
