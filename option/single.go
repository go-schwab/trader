package option

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/samjtro/go-tda/utils"
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
func Single(ticker, contractType, strikeRange, strikeCount, toDate string) ([]CONTRACT, error) {
	req, _ := http.NewRequest("GET", endpoint_option, nil)
	q := req.URL.Query()
	q.Add("symbol", ticker)
	q.Add("contractType", contractType)
	q.Add("range", strikeRange)
	q.Add("strikeCount", strikeCount)
	q.Add("toDate", toDate)
	req.URL.RawQuery = q.Encode()
	body, err := utils.Handler(req)

	if err != nil {
		return []CONTRACT{}, err
	}

	var chain []CONTRACT
	var Type, symbol, exchange, bidAskSize string
	var strikePrice, exp, d2e, bid, ask, last, mark, volatility, delta, gamma, theta, vega, rho, openInterest, timeValue, theoreticalValue, theoreticalVolatility, percentChange, markChange, markPercentChange, intrinsicValue float64
	var inTheMoney bool
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
				strikePrice1 := utils.TrimFL(split2[i+1])

				strikePrice, err = strconv.ParseFloat(strikePrice1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "expirationDate":
				exp1 := utils.TrimFL(split2[i+1])

				exp, err = strconv.ParseFloat(exp1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "daysToExpiration":
				d2e1 := utils.TrimFL(split2[i+1])

				d2e, err = strconv.ParseFloat(d2e1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "bid":
				bid1 := utils.TrimFL(split2[i+1])

				bid, err = strconv.ParseFloat(bid1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "ask":
				ask1 := utils.TrimFL(split2[i+1])

				ask, err = strconv.ParseFloat(ask1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "last":
				last1 := utils.TrimFL(split2[i+1])

				last, err = strconv.ParseFloat(last1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "mark":
				mark1 := utils.TrimFL(split2[i+1])

				mark, err = strconv.ParseFloat(mark1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "bidAskSize":
				bidAskSize = split2[i+2]
			case "volatility":
				volatility1 := utils.TrimFL(split2[i+1])

				volatility, err = strconv.ParseFloat(volatility1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "delta":
				delta1 := utils.TrimFL(split2[i+1])

				delta, err = strconv.ParseFloat(delta1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "gamma":
				gamma1 := utils.TrimFL(split2[i+1])

				gamma, err = strconv.ParseFloat(gamma1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "theta":
				theta1 := utils.TrimFL(split2[i+1])

				theta, err = strconv.ParseFloat(theta1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "vega":
				vega1 := utils.TrimFL(split2[i+1])

				vega, err = strconv.ParseFloat(vega1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "rho":
				rho1 := utils.TrimFL(split2[i+1])

				rho, err = strconv.ParseFloat(rho1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "openInterest":
				openInterest1 := utils.TrimFL(split2[i+1])

				openInterest, err = strconv.ParseFloat(openInterest1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "timeValue":
				timeValue1 := utils.TrimFL(split2[i+1])

				timeValue, err = strconv.ParseFloat(timeValue1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "theoreticalOptionValue":
				theoreticalValue1 := utils.TrimFL(split2[i+1])

				theoreticalValue, err = strconv.ParseFloat(theoreticalValue1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "theoreticalVolatility":
				theoreticalVolatility1 := utils.TrimFL(split2[i+1])

				theoreticalVolatility, err = strconv.ParseFloat(theoreticalVolatility1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "percentChange":
				percentChange1 := utils.TrimFL(split2[i+1])

				percentChange, err = strconv.ParseFloat(percentChange1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "markChange":
				markChange1 := utils.TrimFL(split2[i+1])

				markChange, err = strconv.ParseFloat(markChange1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "markPercentChange":
				markPercentChange1 := utils.TrimFL(split2[i+1])

				markPercentChange, err = strconv.ParseFloat(markPercentChange1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "intrinsicValue":
				intrinsicValue1 := utils.TrimFL(split2[i+1])

				intrinsicValue, err = strconv.ParseFloat(intrinsicValue1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "inTheMoney":
				inTheMoney, err = strconv.ParseBool(utils.TrimFL(split2[i+1]))

				if err != nil {
					log.Fatalf(err.Error())
				}
			}
		}

		contract := CONTRACT{
			TYPE:                   Type,
			SYMBOL:                 symbol,
			STRIKE:                 strikePrice,
			EXCHANGE:               exchange,
			EXPIRATION:             exp,
			DAYS2EXPIRATION:        d2e,
			BID:                    bid,
			ASK:                    ask,
			LAST:                   last,
			MARK:                   mark,
			BIDASK_SIZE:            bidAskSize,
			VOLATILITY:             volatility,
			DELTA:                  delta,
			GAMMA:                  gamma,
			THETA:                  theta,
			VEGA:                   vega,
			RHO:                    rho,
			OPEN_INTEREST:          openInterest,
			TIME_VALUE:             timeValue,
			THEORETICAL_VALUE:      theoreticalValue,
			THEORETICAL_VOLATILITY: theoreticalVolatility,
			PERCENT_CHANGE:         percentChange,
			MARK_CHANGE:            markChange,
			MARK_PERCENT_CHANGE:    markPercentChange,
			INTRINSIC_VALUE:        intrinsicValue,
			IN_THE_MONEY:           inTheMoney,
		}

		chain = append(chain, contract)
	}

	return chain, nil
}
