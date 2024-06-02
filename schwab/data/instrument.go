package data

import (
	"net/http"
	"strconv"
	"strings"

	schwabutils "github.com/samjtro/go-trade/schwab/utils"
	utils "github.com/samjtro/go-trade/utils"
)

func SearchInstrumentFundamental(symbol string) (FundamentalInstrument, error) {
	var cusip, desc, exchange, Type string
	var hi52, lo52, divAmount, divYield, pe, peg, pb, pr, pcf, gmTTM, gmMRQ, npmTTM, npmMRQ, omTTM, omMRQ, roe, roa, roi, qRatio, cRatio, interestCoverage, debtCapital, debtEquity, epsTTM, epsPercentTTM, epsChangeYR, revChangeYR, revChangeTTM, revChangeIn, sharesOutstanding, marketCapFloat, marketCap, bookVPS, beta, vol1, vol10, vol3 float64
	req, err := http.NewRequest("GET", Endpoint_searchinstruments, nil)
	utils.Check(err)
	q := req.URL.Query()
	q.Add("symbol", symbol)
	q.Add("projection", "fundamental")
	req.URL.RawQuery = q.Encode()
	body, err := schwabutils.Handler(req)
	utils.Check(err)
	split := strings.Split(body, "\"")
	for i, x := range split {
		switch x {
		case "cusip":
			cusip = split[i+2]
		case "description":
			desc = split[i+2]
		case "exchange":
			exchange = split[i+2]
		case "assetType":
			Type = split[i+2]
		case "high52":
			hi52, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "low52":
			lo52, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "dividendAmount":
			divAmount, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "dividendYield":
			divYield, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "peRatio":
			pe, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "pegRatio":
			peg, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "pbRatio":
			pb, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "prRatio":
			pr, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "pcfRatio":
			pcf, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "grossMarginTTM":
			gmTTM, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "grossMarginMRQ":
			gmMRQ, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "netProfitMarginTTM":
			npmTTM, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "netProfitMarginMRQ":
			npmMRQ, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "operatingMarginTTM":
			omTTM, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "operatingMarginMRQ":
			omMRQ, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "returnOnEquity":
			roe, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "returnOnAssets":
			roa, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "returnOnInvestment":
			roi, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "quickRatio":
			qRatio, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "currentRatio":
			cRatio, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "interestCoverage":
			interestCoverage, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "totalDebtToCapital":
			debtCapital, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "totalDebtToEquity":
			debtEquity, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "epsTTM":
			epsTTM, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "epsChangePercentTTM":
			epsPercentTTM, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "epsChangeYear":
			epsChangeYR, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "revChangeYear":
			revChangeYR, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "revChangeTTM":
			revChangeTTM, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "revChangeIn":
			revChangeIn, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "sharesOutstanding":
			sharesOutstanding, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "marketCapFloat":
			marketCapFloat, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "marketCap":
			marketCap, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "bookValuePerShare":
			bookVPS, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "beta":
			beta, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "vol1DayAvg":
			vol1, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "vol10DayAvg":
			vol10, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "vol3MonthAvg":
			vol3, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(utils.TrimOneLast(split[i+1])), 64)
			utils.Check(err)
		}
	}

	return FundamentalInstrument{
		Symbol:                 symbol,
		CUSIP:                  cusip,
		DESCRIPTION:            desc,
		EXCHANGE:               exchange,
		TYPE:                   Type,
		HI52:                   hi52,
		LO52:                   lo52,
		DIV_YIELD:              divYield,
		DIV_AMOUNT:             divAmount,
		PE_RATIO:               pe,
		PEG_RATIO:              peg,
		PB_RATIO:               pb,
		PR_RATIO:               pr,
		PCF_RATIO:              pcf,
		GROSS_MARGIN_TTM:       gmTTM,
		GROSS_MARGIN_MRQ:       gmMRQ,
		NET_PROFIT_MARGIN_TTM:  npmTTM,
		NET_PROFIT_MARGIN_MRQ:  npmMRQ,
		OPERATING_MARGIN_TTM:   omTTM,
		OPERATING_MARGIN_MRQ:   omMRQ,
		RETURN_ON_EQUITY:       roe,
		RETURN_ON_ASSETS:       roa,
		RETURN_ON_INVESTMENT:   roi,
		QUICK_RATIO:            qRatio,
		CURRENT_RATIO:          cRatio,
		INTEREST_COVERAGE:      interestCoverage,
		TOTAL_DEBT_TO_CAPITAL:  debtCapital,
		TOTAL_DEBT_TO_EQUITY:   debtEquity,
		EPS_TTM:                epsTTM,
		EPS_CHANGE_PERCENT_TTM: epsPercentTTM,
		EPS_CHANGE_YR:          epsChangeYR,
		REV_CHANGE_YR:          revChangeYR,
		REV_CHANGE_TTM:         revChangeTTM,
		REV_CHANGE_IN:          revChangeIn,
		SHARES_OUTSTANDING:     sharesOutstanding,
		MARKET_CAP_FLOAT:       marketCapFloat,
		MARKET_CAP:             marketCap,
		BOOK_VALUE_PER_SHARE:   bookVPS,
		BETA:                   beta,
		VOL_1DAY:               vol1,
		VOL_10DAY:              vol10,
		VOL_3MON:               vol3,
	}, nil
}
