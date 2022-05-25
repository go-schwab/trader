package instrument 

import (
	"fmt"
	"strings"
	"net/http"
	. "github.com/samjtro/go-tda/utils"
)

// for use with the Get function
type SIMPLE struct {
	CUSIP		string
	TICKER		string
	DESCRIPTION	string
	EXCHANGE	string
	TYPE		string
}

// for use with the Fundamental function
type FUNDAMENTAL struct {
	TICKER			string
	CUSIP			string
	DESCRIPTION		string
	EXCHANGE		string
	TYPE			string
	HI52			string
	LO52			string
	DIV_YIELD		string
	DIV_AMOUNT		string
	PE_RATIO		string
	PEG_RATIO		string
	PB_RATIO		string
	PR_RATIO		string
	PCF_RATIO		string
	GROSS_MARGIN_TTM	string
	GROSS_MARGIN_MRQ	string
	NET_PROFIT_MARGIN_TTM	string
	NET_PROFIT_MARGIN_MRQ	string
	OPERATING_MARGIN_TTM	string
	OPERATING_MARGIN_MRQ	string
	RETURN_ON_EQUITY	string
	RETURN_ON_ASSETS	string
	RETURN_ON_INVESTMENT	string
	QUICK_RATIO		string
	CURRENT_RATIO		string
	INTEREST_COVERAGE	string
	TOTAL_DEBT_TO_CAPITAL	string
	TOTAL_DEBT_TO_EQUITY	string
	EPS_TTM			string
	EPS_CHANGE_PERCENT_TTM	string
	EPS_CHANGE_YR		string
	REV_CHANGE_YR		string
	REV_CHANGE_TTM		string
	REV_CHANGE_IN		string
	SHARES_OUTSTANDING	string
	MARKET_CAP_FLOAT	string
	MARKET_CAP		string
	BOOK_VALUE_PER_SHARE	string
	BETA			string
	VOL_1DAY		string
	VOL_10DAY		string
	VOL_3MON		string
}

var endpoint_searchinstrument string = "https://api.tdameritrade.com/v1/instruments"
var endpoint_getinstrument string = "https://api.tdameritrade.com/v1/instruments/%s"// cusip

// Get returns a SIMPLE; with simple fundamental information regarding the desired ticker,
// it takes one parameter:
// cusip = "037833100", etc.
func Get(ticker string) SIMPLE {
	req2,_ := http.NewRequest("GET",endpoint_searchinstrument,nil)
	q2 := req2.URL.Query()
	q2.Add("symbol",ticker)
	q2.Add("projection","fundamental")
	req2.URL.RawQuery = q2.Encode()
	body2 := Handler(req2)

	split2 := strings.Split(body2,"\"")
	for i,x := range split2 {
		if(x == "cusip") { cusip = split2[i+2] }
	}

	url := fmt.Sprintf(endpoint_getinstrument,cusip)
	req,_ := http.NewRequest("GET",url,nil)
	body := Handler(req)

	var cusip,desc,exchange,Type string

	split := strings.Split(body,"\"")
	for i,x := range split {
		if(x == "description") { desc = split[i+2]
		} else if(x == "exchange") { exchange = split[i+2]
		} else if(x == "assetType") { Type = split[i+2] }
	}

	return SIMPLE{
		CUSIP:		cusip,
		TICKER:		ticker,
		DESCRIPTION:	desc,
		EXCHANGE:	exchange,
		TYPE:		Type,
	}
}

// returns a FUNDAMENTAL; containing information regarding both price and underlying information and history
// Returns fundamental data for a single instrument specified by ticker
func Fundamental(ticker string) FUNDAMENTAL {
	req,_ := http.NewRequest("GET",endpoint_searchinstrument,nil)
	q := req.URL.Query()
	q.Add("symbol",ticker)
	q.Add("projection","fundamental")
	req.URL.RawQuery = q.Encode()
	body := Handler(req)

	var cusip,desc,exchange,Type,hi52,lo52,divAmount,divYield,pe,peg,pb,pr,pcf,gmTTM,gmMRQ,npmTTM,npmMRQ,omTTM,omMRQ,roe,roa,roi,qRatio,cRatio,interestCoverage,debtCapital,debtEquity,epsTTM,epsPercentTTM,epsChangeYR,revChangeYR,revChangeTTM,revChangeIn,sharesOutstanding,marketCapFloat,marketCap,bookVPS,beta,vol1,vol10,vol3 string

	split := strings.Split(body,"\"")
	for i,x := range split {
		if(x == "cusip") { cusip = split[i+2]
		} else if(x == "description") { desc = split[i+2]
		} else if(x == "exchange") { exchange = split[i+2]
		} else if(x == "assetType") { Type = split[i+2]
		} else if(x == "high52") { hi52 = split[i+1]
		} else if(x == "low52") { lo52 = split[i+1]
		} else if(x == "dividendAmount") { divAmount = split[i+1]
		} else if(x == "dividendYield") { divYield = split[i+1] 
		} else if(x == "peRatio") { pe = split[i+1]
		} else if(x == "pegRatio") { peg = split[i+1]
		} else if(x == "pbRatio") { pb = split[i+1]
		} else if(x == "prRatio") { pr = split[i+1]
		} else if(x == "pcfRatio") { pcf = split[i+1]
		} else if(x == "grossMarginTTM") { gmTTM = split[i+1]
		} else if(x == "grossMarginMRQ") { gmMRQ = split[i+1]
		} else if(x == "netProfitMarginTTM") { npmTTM = split[i+1]
		} else if(x == "netProfitMarginMRQ") { npmMRQ = split[i+1]
		} else if(x == "operatingMarginTTM") { omTTM = split[i+1]
		} else if(x == "operatingMarginMRQ") { omMRQ = split[i+1]
		} else if(x == "returnOnEquity") { roe = split[i+1]
		} else if(x == "returnOnAssets") { roa = split[i+1]
		} else if(x == "returnOnInvestment") { roi = split[i+1]
		} else if(x == "quickRatio") { qRatio = split[i+1]
		} else if(x == "currentRatio") { cRatio = split[i+1]
		} else if(x == "interestCoverage") { interestCoverage = split[i+1]
		} else if(x == "totalDebtToCapital") { debtCapital = split[i+1]
		} else if(x == "totalDebtToEquity") { debtEquity = split[i+1]
		} else if(x == "epsTTM") { epsTTM = split[i+1]
		} else if(x == "epsChangePercentTTM") { epsPercentTTM = split[i+1]
		} else if(x == "epsChangeYear") { epsChangeYR = split[i+1]
		} else if(x == "revChangeYear") { revChangeYR = split[i+1]
		} else if(x == "revChangeTTM") { revChangeTTM = split[i+1]
		} else if(x == "revChangeIn") { revChangeIn = split[i+1]
		} else if(x == "sharesOutstanding") { sharesOutstanding = split[i+1]
		} else if(x == "marketCapFloat") { marketCapFloat = split[i+1]
		} else if(x == "marketCap") { marketCap = split[i+1]
		} else if(x == "bookValuePerShare") { bookVPS = split[i+1]
		} else if(x == "beta") { beta = split[i+1]
		} else if(x == "vol1DayAvg") { vol1 = split[i+1]
		} else if(x == "vol10DayAvg") { vol10 = split[i+1]
		} else if(x == "vol3MonthAvg") { vol3 = split[i+1]
		}
	}

	return FUNDAMENTAL{
		TICKER:			ticker,
		CUSIP:			cusip,
		DESCRIPTION:		desc,
		EXCHANGE:		exchange,
		TYPE:			Type,
		HI52:			TrimFL(hi52),
		LO52:			TrimFL(lo52),
		DIV_YIELD:		TrimFL(divYield),
		DIV_AMOUNT:		TrimFL(divAmount),
		PE_RATIO:		TrimFL(pe),
		PEG_RATIO:		TrimFL(peg),
		PB_RATIO:		TrimFL(pb),
		PR_RATIO:		TrimFL(pr),
		PCF_RATIO:		TrimFL(pcf),
		GROSS_MARGIN_TTM:	TrimFL(gmTTM),
		GROSS_MARGIN_MRQ:	TrimFL(gmMRQ),
		NET_PROFIT_MARGIN_TTM:  TrimFL(npmTTM),
		NET_PROFIT_MARGIN_MRQ:  TrimFL(npmMRQ),
		OPERATING_MARGIN_TTM:   TrimFL(omTTM),
		OPERATING_MARGIN_MRQ:   TrimFL(omMRQ),
		RETURN_ON_EQUITY:	TrimFL(roe),
		RETURN_ON_ASSETS:	TrimFL(roa),
		RETURN_ON_INVESTMENT:	TrimFL(roi),
		QUICK_RATIO:		TrimFL(qRatio),
		CURRENT_RATIO:		TrimFL(cRatio),
		INTEREST_COVERAGE:	TrimFL(interestCoverage),
		TOTAL_DEBT_TO_CAPITAL:	TrimFL(debtCapital),
		TOTAL_DEBT_TO_EQUITY:	TrimFL(debtEquity),
		EPS_TTM:		TrimFL(epsTTM),
		EPS_CHANGE_PERCENT_TTM:	TrimFL(epsPercentTTM),
		EPS_CHANGE_YR:		TrimFL(epsChangeYR),
		REV_CHANGE_YR:		TrimFL(revChangeYR),
		REV_CHANGE_TTM:		TrimFL(revChangeTTM),
		REV_CHANGE_IN:		TrimFL(revChangeIn),
		SHARES_OUTSTANDING:	TrimFL(sharesOutstanding),
		MARKET_CAP_FLOAT:	TrimFL(marketCapFloat),
		MARKET_CAP:		TrimFL(marketCap),
		BOOK_VALUE_PER_SHARE:	TrimFL(bookVPS),
		BETA:			TrimFL(beta),
		VOL_1DAY:		TrimFL(vol1),
		VOL_10DAY:		TrimFL(vol10),
		VOL_3MON:		TrimFL(TrimL(vol3)),
	}
}

// desc-regex: Search description with full regex support. Example: symbol=XYZ.[A-C] returns all instruments whose descriptions contain a word beginning with XYZ followed by a character A through C.

