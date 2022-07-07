package instrument

import (
	"log"
	"net/http"
	"strings"

	. "github.com/samjtro/go-tda/utils"
)

// Returns a FUNDAMENTAL; containing information regarding both price history and fundamentals.
func Fundamental(ticker string) FUNDAMENTAL {
	req, _ := http.NewRequest("GET", endpoint_searchinstrument, nil)
	q := req.URL.Query()
	q.Add("symbol", ticker)
	q.Add("projection", "fundamental")
	req.URL.RawQuery = q.Encode()
	body, err := Handler(req)

	if err != nil {
		log.Fatal(err)
	}

	var cusip, desc, exchange, Type, hi52, lo52, divAmount, divYield, pe, peg, pb, pr, pcf, gmTTM, gmMRQ, npmTTM, npmMRQ, omTTM, omMRQ, roe, roa, roi, qRatio, cRatio, interestCoverage, debtCapital, debtEquity, epsTTM, epsPercentTTM, epsChangeYR, revChangeYR, revChangeTTM, revChangeIn, sharesOutstanding, marketCapFloat, marketCap, bookVPS, beta, vol1, vol10, vol3 string
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
			hi52 = split[i+1]
		case "low52":
			lo52 = split[i+1]
		case "dividendAmount":
			divAmount = split[i+1]
		case "dividendYield":
			divYield = split[i+1]
		case "peRatio":
			pe = split[i+1]
		case "pegRatio":
			peg = split[i+1]
		case "pbRatio":
			pb = split[i+1]
		case "prRatio":
			pr = split[i+1]
		case "pcfRatio":
			pcf = split[i+1]
		case "grossMarginTTM":
			gmTTM = split[i+1]
		case "grossMarginMRQ":
			gmMRQ = split[i+1]
		case "netProfitMarginTTM":
			npmTTM = split[i+1]
		case "netProfitMarginMRQ":
			npmMRQ = split[i+1]
		case "operatingMarginTTM":
			omTTM = split[i+1]
		case "operatingMarginMRQ":
			omMRQ = split[i+1]
		case "returnOnEquity":
			roe = split[i+1]
		case "returnOnAssets":
			roa = split[i+1]
		case "returnOnInvestment":
			roi = split[i+1]
		case "quickRatio":
			qRatio = split[i+1]
		case "currentRatio":
			cRatio = split[i+1]
		case "interestCoverage":
			interestCoverage = split[i+1]
		case "totalDebtToCapital":
			debtCapital = split[i+1]
		case "totalDebtToEquity":
			debtEquity = split[i+1]
		case "epsTTM":
			epsTTM = split[i+1]
		case "epsChangePercentTTM":
			epsPercentTTM = split[i+1]
		case "epsChangeYear":
			epsChangeYR = split[i+1]
		case "revChangeYear":
			revChangeYR = split[i+1]
		case "revChangeTTM":
			revChangeTTM = split[i+1]
		case "revChangeIn":
			revChangeIn = split[i+1]
		case "sharesOutstanding":
			sharesOutstanding = split[i+1]
		case "marketCapFloat":
			marketCapFloat = split[i+1]
		case "marketCap":
			marketCap = split[i+1]
		case "bookValuePerShare":
			bookVPS = split[i+1]
		case "beta":
			beta = split[i+1]
		case "vol1DayAvg":
			vol1 = split[i+1]
		case "vol10DayAvg":
			vol10 = split[i+1]
		case "vol3MonthAvg":
			vol3 = split[i+1]
		}
	}

	return FUNDAMENTAL{
		TICKER:                 ticker,
		CUSIP:                  cusip,
		DESCRIPTION:            desc,
		EXCHANGE:               exchange,
		TYPE:                   Type,
		HI52:                   TrimFL(hi52),
		LO52:                   TrimFL(lo52),
		DIV_YIELD:              TrimFL(divYield),
		DIV_AMOUNT:             TrimFL(divAmount),
		PE_RATIO:               TrimFL(pe),
		PEG_RATIO:              TrimFL(peg),
		PB_RATIO:               TrimFL(pb),
		PR_RATIO:               TrimFL(pr),
		PCF_RATIO:              TrimFL(pcf),
		GROSS_MARGIN_TTM:       TrimFL(gmTTM),
		GROSS_MARGIN_MRQ:       TrimFL(gmMRQ),
		NET_PROFIT_MARGIN_TTM:  TrimFL(npmTTM),
		NET_PROFIT_MARGIN_MRQ:  TrimFL(npmMRQ),
		OPERATING_MARGIN_TTM:   TrimFL(omTTM),
		OPERATING_MARGIN_MRQ:   TrimFL(omMRQ),
		RETURN_ON_EQUITY:       TrimFL(roe),
		RETURN_ON_ASSETS:       TrimFL(roa),
		RETURN_ON_INVESTMENT:   TrimFL(roi),
		QUICK_RATIO:            TrimFL(qRatio),
		CURRENT_RATIO:          TrimFL(cRatio),
		INTEREST_COVERAGE:      TrimFL(interestCoverage),
		TOTAL_DEBT_TO_CAPITAL:  TrimFL(debtCapital),
		TOTAL_DEBT_TO_EQUITY:   TrimFL(debtEquity),
		EPS_TTM:                TrimFL(epsTTM),
		EPS_CHANGE_PERCENT_TTM: TrimFL(epsPercentTTM),
		EPS_CHANGE_YR:          TrimFL(epsChangeYR),
		REV_CHANGE_YR:          TrimFL(revChangeYR),
		REV_CHANGE_TTM:         TrimFL(revChangeTTM),
		REV_CHANGE_IN:          TrimFL(revChangeIn),
		SHARES_OUTSTANDING:     TrimFL(sharesOutstanding),
		MARKET_CAP_FLOAT:       TrimFL(marketCapFloat),
		MARKET_CAP:             TrimFL(marketCap),
		BOOK_VALUE_PER_SHARE:   TrimFL(bookVPS),
		BETA:                   TrimFL(beta),
		VOL_1DAY:               TrimFL(vol1),
		VOL_10DAY:              TrimFL(vol10),
		VOL_3MON:               TrimFL(TrimL(vol3)),
	}
}
