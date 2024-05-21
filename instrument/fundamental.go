package instrument

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/samjtro/go-tda/utils"
)

// Returns a FUNDAMENTAL; containing information regarding both price history and fundamentals.
func Fundamental(ticker string) (FUNDAMENTAL, error) {
	req, _ := http.NewRequest("GET", endpoint_searchinstrument, nil)
	q := req.URL.Query()
	q.Add("symbol", ticker)
	q.Add("projection", "fundamental")
	req.URL.RawQuery = q.Encode()
	body, err := utils.Handler(req)

	if err != nil {
		return FUNDAMENTAL{}, err
	}

	var cusip, desc, exchange, Type string
	var hi52, lo52, divAmount, divYield, pe, peg, pb, pr, pcf, gmTTM, gmMRQ, npmTTM, npmMRQ, omTTM, omMRQ, roe, roa, roi, qRatio, cRatio, interestCoverage, debtCapital, debtEquity, epsTTM, epsPercentTTM, epsChangeYR, revChangeYR, revChangeTTM, revChangeIn, sharesOutstanding, marketCapFloat, marketCap, bookVPS, beta, vol1, vol10, vol3 float64
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
			hi521 := utils.TrimFL(split[i+1])

			hi52, err = strconv.ParseFloat(hi521, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "low52":
			lo521 := utils.TrimFL(split[i+1])

			lo52, err = strconv.ParseFloat(lo521, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "dividendAmount":
			divAmount1 := utils.TrimFL(split[i+1])

			divAmount, err = strconv.ParseFloat(divAmount1, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "dividendYield":
			divYield1 := utils.TrimFL(split[i+1])

			divYield, err = strconv.ParseFloat(divYield1, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "peRatio":
			pe1 := utils.TrimFL(split[i+1])

			pe, err = strconv.ParseFloat(pe1, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "pegRatio":
			peg1 := utils.TrimFL(split[i+1])

			peg, err = strconv.ParseFloat(peg1, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "pbRatio":
			pb1 := utils.TrimFL(split[i+1])

			pb, err = strconv.ParseFloat(pb1, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "prRatio":
			pr1 := utils.TrimFL(split[i+1])

			pr, err = strconv.ParseFloat(pr1, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "pcfRatio":
			pcf1 := utils.TrimFL(split[i+1])

			pcf, err = strconv.ParseFloat(pcf1, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "grossMarginTTM":
			gmTTM1 := utils.TrimFL(split[i+1])

			gmTTM, err = strconv.ParseFloat(gmTTM1, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "grossMarginMRQ":
			gmMRQ1 := utils.TrimFL(split[i+1])

			gmMRQ, err = strconv.ParseFloat(gmMRQ1, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "netProfitMarginTTM":
			npmTTM1 := utils.TrimFL(split[i+1])

			npmTTM, err = strconv.ParseFloat(npmTTM1, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "netProfitMarginMRQ":
			npmMRQ1 := utils.TrimFL(split[i+1])

			npmMRQ, err = strconv.ParseFloat(npmMRQ1, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "operatingMarginTTM":
			omTTM1 := utils.TrimFL(split[i+1])

			omTTM, err = strconv.ParseFloat(omTTM1, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "operatingMarginMRQ":
			omMRQ1 := utils.TrimFL(split[i+1])

			omMRQ, err = strconv.ParseFloat(omMRQ1, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "returnOnEquity":
			roe1 := utils.TrimFL(split[i+1])

			roe, err = strconv.ParseFloat(roe1, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "returnOnAssets":
			roa1 := utils.TrimFL(split[i+1])

			roa, err = strconv.ParseFloat(roa1, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "returnOnInvestment":
			roi1 := utils.TrimFL(split[i+1])

			roi, err = strconv.ParseFloat(roi1, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "quickRatio":
			qRatio1 := utils.TrimFL(split[i+1])

			qRatio, err = strconv.ParseFloat(qRatio1, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "currentRatio":
			cRatio1 := utils.TrimFL(split[i+1])

			cRatio, err = strconv.ParseFloat(cRatio1, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "interestCoverage":
			interestCoverage1 := utils.TrimFL(split[i+1])

			interestCoverage, err = strconv.ParseFloat(interestCoverage1, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "totalDebtToCapital":
			debtCapital1 := utils.TrimFL(split[i+1])

			debtCapital, err = strconv.ParseFloat(debtCapital1, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "totalDebtToEquity":
			debtEquity1 := utils.TrimFL(split[i+1])

			debtEquity, err = strconv.ParseFloat(debtEquity1, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "epsTTM":
			epsTTM1 := utils.TrimFL(split[i+1])

			epsTTM, err = strconv.ParseFloat(epsTTM1, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "epsChangePercentTTM":
			epsPercentTTM1 := utils.TrimFL(split[i+1])

			epsPercentTTM, err = strconv.ParseFloat(epsPercentTTM1, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "epsChangeYear":
			epsChangeYR1 := utils.TrimFL(split[i+1])

			epsChangeYR, err = strconv.ParseFloat(epsChangeYR1, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "revChangeYear":
			revChangeYR1 := utils.TrimFL(split[i+1])

			revChangeYR, err = strconv.ParseFloat(revChangeYR1, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "revChangeTTM":
			revChangeTTM1 := utils.TrimFL(split[i+1])

			revChangeTTM, err = strconv.ParseFloat(revChangeTTM1, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "revChangeIn":
			revChangeIn1 := utils.TrimFL(split[i+1])

			revChangeIn, err = strconv.ParseFloat(revChangeIn1, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "sharesOutstanding":
			sharesOutstanding1 := utils.TrimFL(split[i+1])

			sharesOutstanding, err = strconv.ParseFloat(sharesOutstanding1, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "marketCapFloat":
			marketCapFloat1 := utils.TrimFL(split[i+1])

			marketCapFloat, err = strconv.ParseFloat(marketCapFloat1, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "marketCap":
			marketCap1 := utils.TrimFL(split[i+1])

			marketCap, err = strconv.ParseFloat(marketCap1, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "bookValuePerShare":
			bookVPS1 := utils.TrimFL(split[i+1])

			bookVPS, err = strconv.ParseFloat(bookVPS1, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "beta":
			beta1 := utils.TrimFL(split[i+1])

			beta, err = strconv.ParseFloat(beta1, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "vol1DayAvg":
			vol11 := utils.TrimFL(split[i+1])

			vol1, err = strconv.ParseFloat(vol11, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "vol10DayAvg":
			vol101 := utils.TrimFL(split[i+1])

			vol10, err = strconv.ParseFloat(vol101, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		case "vol3MonthAvg":
			vol31 := utils.TrimFL(utils.TrimL(split[i+1]))

			vol3, err = strconv.ParseFloat(vol31, 64)

			if err != nil {
				log.Fatalf(err.Error())
			}
		}
	}

	return FUNDAMENTAL{
		TICKER:                 ticker,
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
