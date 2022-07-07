package option

import (
	"log"
	"net/http"
	"strings"

	. "github.com/samjtro/go-tda/utils"
)

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
