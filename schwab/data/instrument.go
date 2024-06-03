package data

import (
	"net/http"
	"strconv"
	"strings"

	schwabutils "github.com/samjtro/go-trade/schwab/utils"
	utils "github.com/samjtro/go-trade/utils"
)

// SearchInstrumentSimple returns instrument's simples.
// It takes on param:
func SearchInstrumentSimple(symbol string) {

}

// SearchInstrumentFundamental returns instrument's fundamentals.
// It takes one param:
func SearchInstrumentFundamental(symbol string) (FundamentalInstrument, error) {
	// Craft, send request
	instrument := FundamentalInstrument{Symbol: symbol}
	req, err := http.NewRequest("GET", Endpoint_searchinstruments, nil)
	utils.Check(err)
	q := req.URL.Query()
	q.Add("symbol", symbol)
	q.Add("projection", "fundamental")
	req.URL.RawQuery = q.Encode()
	body, err := schwabutils.Handler(req)
	// Parse return
	utils.Check(err)
	split := strings.Split(body, "\"")
	for i, x := range split {
		switch x {
		case "cusip":
			instrument.Cusip = split[i+2]
		case "description":
			instrument.Description = split[i+2]
		case "exchange":
			instrument.Exchange = split[i+2]
		case "assetType":
			instrument.Type = split[i+2]
		case "high52":
			instrument.Hi52, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "low52":
			instrument.Lo52, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "dividendAmount":
			instrument.DivAmount, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "dividendYield":
			instrument.DivYield, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "dividendDate":
			instrument.DividendDate = utils.TrimOneFirstOneLast(split[i+2])
		case "peRatio":
			instrument.PE, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "pegRatio":
			instrument.PEG, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "pbRatio":
			instrument.PB, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "prRatio":
			instrument.PR, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "pcfRatio":
			instrument.PCF, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "grossMarginTTM":
			instrument.GrossMarginTTM, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "grossMarginMRQ":
			instrument.GrossMarginMRQ, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "netProfitMarginTTM":
			instrument.NetProfitMarginTTM, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "netProfitMarginMRQ":
			instrument.NetProfitMarginMRQ, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "operatingMarginTTM":
			instrument.OperatingMarginTTM, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "operatingMarginMRQ":
			instrument.OperatingMarginMRQ, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "returnOnEquity":
			instrument.ROE, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "returnOnAssets":
			instrument.ROA, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "returnOnInvestment":
			instrument.ROI, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "quickRatio":
			instrument.QuickRatio, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "currentRatio":
			instrument.CurrentRatio, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "interestCoverage":
			instrument.InterestCoverage, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "totalDebtToCapital":
			instrument.TotalDebtToCapital, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "ltDebtToEquity":
			instrument.LTDebtToEquity, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "totalDebtToEquity":
			instrument.TotalDebtToEquity, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "epsTTM":
			instrument.EPSTTM, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "epsChangePercentTTM":
			instrument.EPSChangePercentTTM, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "epsChangeYear":
			instrument.EPSChangeYear, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "epsChange":
			instrument.EPSChange, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "revChangeYear":
			instrument.RevenueChangeYear, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "revChangeTTM":
			instrument.RevenueChangeTTM, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "revChangeIn":
			instrument.RevenueChangeIn, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "sharesOutstanding":
			instrument.SharesOutstanding, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "marketCapFloat":
			instrument.MarketCapFloat, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "marketCap":
			instrument.MarketCap, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "bookValuePerShare":
			instrument.BookValuePerShare, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "shortIntToFloat":
			instrument.ShortIntToFloat, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "shortIntDayToCover":
			instrument.ShortIntDayToCover, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "divGrowthRate3Year":
			instrument.DividendGrowthRate3Year, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "dividendPayAmount":
			instrument.DividendPayAmount, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "dividendPayDate":
			instrument.DividendPayDate = utils.TrimOneFirstOneLast(split[i+2])
		case "beta":
			instrument.Beta, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "vol1DayAvg":
			instrument.Vol1DayAverage, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "vol10DayAvg":
			instrument.Vol10DayAverage, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "vol3MonthAvg":
			instrument.Vol3MonthAverage, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(utils.TrimOneLast(split[i+1])), 64)
			utils.Check(err)
		case "avg10DaysVolume":
			instrument.Avg10DaysVolume, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "avg1DayVolume":
			instrument.Avg1DayVolume, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "avg3MonthVolume":
			instrument.Avg3MonthVolume, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "declarationDate":
			instrument.DeclarationDate = utils.TrimOneFirstOneLast(split[i+2])
		case "dividendFreq":
			instrument.DividendFrequency, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "eps":
			instrument.EPS, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "dtnVolume":
			instrument.DTNVolume, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		case "nextDividendPayDate":
			instrument.NextDividendPayDate = utils.TrimOneFirstOneLast(split[i+2])
		case "nextDividendDate":
			instrument.NextDividendDate = utils.TrimOneFirstOneLast(split[i+2])
		case "fundLeverageFactor":
			instrument.FundLeverageFactor, err = strconv.ParseFloat(utils.TrimOneFirstOneLast(split[i+1]), 64)
			utils.Check(err)
		}
	}
	return instrument, nil
}
