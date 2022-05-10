package instrument

import (
	"fmt"
	"strings"
	"net/http"
	. "github.com/samjtro/go-tda/utils"
)

type SIMPLE struct {
	CUSIP		string
	TICKER		string
	DESCRIPTION	string
	EXCHANGE	string
	TYPE		string
}

type FUNDAMENTAL struct {
	TICKER			string
	/*CUSIP			string
	DESCRIPTION		string
	EXCHANGE		string
	TYPE			string*/
	HI52			string
	LO52			string
	DIV_YIELD		string
	DIV_AMOUNT		string
	PE_RATIO		string
	PEG_RATIO		string
	PB_RATIO		string
	PR_RATIO		string
	PCF_RATIO		string
	/*GROSS_MARGIN_TTM	string
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
	VOL_3MON		string*/
}

var endpoint_searchinstrument string = "https://api.tdameritrade.com/v1/instruments"
var endpoint_getinstrument string = "https://api.tdameritrade.com/v1/instruments/%s"//  		--> cusip

// Get returns a string; with simple fundamental information regarding the desired ticker,
// it takes one parameter:
// cusip = "037833100", etc.
func Get(ticker string) SIMPLE {
	req2,_ := http.NewRequest("GET",endpoint_searchinstrument,nil)
	q2 := req2.URL.Query()
	q2.Add("symbol",ticker)
	q2.Add("projection","fundamental")
	req2.URL.RawQuery = q2.Encode()
	body2 := Handler(req2)

	var cusip string

	split2 := strings.Split(body2,"\"")
	for i,x := range split2 {
		if(x == "cusip") { cusip = split2[i+2] }
	}

	url := fmt.Sprintf(endpoint_getinstrument,cusip)
	req,_ := http.NewRequest("GET",url,nil)
	body := Handler(req)

	var desc,exchange,Type string

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

// Search uses more detailed parameters to return a string; containing more detailed information of various types on your desired ticker,
// it takes two parameters:
// ticker = "AAPL", etc.
// projection = the type of search to perform: from td-ameritrade's website:
// symbol-search: Retrieve instrument data of a specific symbol or cusip 
// symbol-regex: Retrieve instrument data for all symbols matching regex. Example: symbol=XYZ.* will return all symbols beginning with XYZ 
// desc-search: Retrieve instrument data for instruments whose description contains the word supplied. Example: symbol=FakeCompany will return all instruments with FakeCompany in the description. 
// desc-regex: Search description with full regex support. Example: symbol=XYZ.[A-C] returns all instruments whose descriptions contain a word beginning with XYZ followed by a character A through C. 
// fundamental: Returns fundamental data for a single instrument specified by exact symbol.'
func Fundamental(ticker string) FUNDAMENTAL {
	req,_ := http.NewRequest("GET",endpoint_searchinstrument,nil)
	q := req.URL.Query()
	q.Add("symbol",ticker)
	q.Add("projection","fundamental")
	req.URL.RawQuery = q.Encode()
	body := Handler(req)

	//var cusip,desc,exchange,Type,hi52,lo52,div_yield,div_amount,pe_ratio,peg_ratio,pb_ratio,pr_ratio,pcf_ratio,gross_margin_ttm,gross_margin_mrq,net_profit_margin_ttm,net_profit_margin_mrq,operating_margin_ttm,operating_margin_mrq,return_on_equity,return_on_assets,return_on_investment,quick_ratio,current_ratio,interest_coverage,total_debt_to_capital,total_debt_to_equity,eps_ttm,eps_change_percent_ttm,eps_change_per_yr,eps_change_yr,rev_change_yr,rev_change_ttm,rev_change_in,shares_outstanding,market_cap_float,market_cap,book_value_per_share,beta,vol_1day,vol_10day,vol_3mon string

	var hi52,lo52,div_amount,div_yield,pe,peg,pb,pr,pcf string

	split := strings.Split(body,"\"")
	for i,x := range split {
		if(x == "high52") { hi52 = split[i+1]
		} else if(x == "low52") { lo52 = split[i+1]
		} else if(x == "dividendAmount") { div_amount = split[i+1]
		} else if(x == "dividendYield") { div_yield = split[i+1] 
		} else if(x == "peRatio") { pe = split[i+1]
		} else if(x == "pegRatio") { peg = split[i+1]
		} else if(x == "pbRatio") { pb = split[i+1]
		} else if(x == "prRatio") { pr = split[i+1]
		} else if(x == "pcfRatio") { pcf = split[i+1]
		}
	}

	hi52 = TrimFL(hi52)
	lo52 = TrimFL(lo52)
	div_yield = TrimFL(div_yield)
	div_amount = TrimFL(div_amount)
	pe = TrimFL(pe)
	peg = TrimFL(peg)
	pb = TrimFL(pb)
	pr = TrimFL(pr)
	pcf = TrimFL(pcf)

	return FUNDAMENTAL{
		TICKER:		ticker,
		HI52:		hi52,
		LO52:		lo52,
		DIV_YIELD:	div_yield,
		DIV_AMOUNT:	div_amount,
		PE_RATIO:	pe,
		PEG_RATIO:	peg,
		PB_RATIO:	pb,
		PR_RATIO:	pr,
		PCF_RATIO:	pcf,
	}
}

