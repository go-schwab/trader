package data

import (
	"fmt"
	"time"
	"strings"
	"net/http"
	. "github.com/samjtro/go-tda/utils"
)

// for use with RealTime
type QUOTE struct {
	DATETIME	string
	TICKER		string
	MARK		string
	VOLUME		string
	VOLATILITY	string
	BID		string
	ASK		string
	LAST		string
	OPEN		string
	CLOSE		string
	HI		string
	LO		string
	HI52		string
	LO52		string
	PE_RATIO	string
}

// for use with PriceHistory 
type FRAME struct {
	DATETIME	string
	VOLUME		string
	OPEN		string
	CLOSE		string
	HI		string
	LO		string
}

var endpoint_realtime string = "https://api.tdameritrade.com/v1/marketdata/%s/quotes"// 	 	--> symbol
var endpoint_pricehistory string = "https://api.tdameritrade.com/v1/marketdata/%s/pricehistory"// 	--> symbol

// RealTime returns a QUOTE; containing a real time quote of the desired stock's performance with a number of different indicators (including volatility, volume, price, fundamentals & more), 
// it takes one parameter:
// ticker = "AAPL", etc.
func RealTime(ticker string) QUOTE {
	dt := Now(time.Now())
	url := fmt.Sprintf(endpoint_realtime,ticker)
	req,_ := http.NewRequest("GET",url,nil)
	body := Handler(req)

	var bid,ask,last,open,hi,lo,closeP,mark,volume,volatility,hi52,lo52,pe string

	split := strings.Split(body,"\"")
	for i,x := range split {
		if(x == "bidPrice") { bid = split[i+1]
		} else if(x == "askPrice") { ask = split[i+1]
		} else if(x == "lastPrice") { last = split[i+1]
		} else if(x == "openPrice") { open = split[i+1]
		} else if(x == "highPrice") { hi = split[i+1] 
		} else if(x == "lowPrice") { lo = split[i+1] 
		} else if(x == "closePrice") { closeP = split[i+1] 
		} else if(x == "mark") { mark = split[i+1] 
		} else if(x == "totalVolume") { volume = split[i+1] 
		} else if(x == "volatility") { volatility = split[i+1] 
		} else if(x == "52WkHigh") { hi52 = split[i+1]
		} else if(x == "52WkLow") { lo52 = split[i+1]
		} else if(x == "peRatio") { pe = split[i+1]
		}
	}

	return QUOTE{
		DATETIME: 	dt,
		TICKER:   	ticker,
		MARK:		TrimFL(mark),
		VOLUME:		TrimFL(volume),
		VOLATILITY:	TrimFL(volatility),
		BID:	  	TrimFL(bid),
		ASK:	  	TrimFL(ask),
		LAST:     	TrimFL(last),
		OPEN:	  	TrimFL(open),
		CLOSE:		TrimFL(closeP),
		HI:	  	TrimFL(hi),
		LO:	  	TrimFL(lo),
		HI52:		TrimFL(hi52),
		LO52:		TrimFL(lo52),
		PE_RATIO:	TrimFL(pe),
	}
}

// PriceHistory returns a string; containing a series of candles with price volume & datetime info per candlestick,
// it takes five parameters:
// ticker = "AAPL", etc.
// periodType = "day", "month", "year", "ytd" - default is "day"
// period = the number of periods to show
// frequencyType = the type of frequency with which each candle is formed; valid fTypes by pType
// "day": "minute"
// "month": "daily", "weekly"
// "year": "daily", "weekly", "monthly"
// "ytd": "daily", "weekly"
// frequency = the number of the frequencyType included in each candle; valid freqs by fType
// "minute": 1,5,10,15,30
// "daily": 1
// "weekly": 1
// "monthly": 1
func PriceHistory(ticker,periodType,period,frequencyType,frequency string) []FRAME {
	url := fmt.Sprintf(endpoint_pricehistory,ticker)
	req,_ := http.NewRequest("GET",url,nil)
	q := req.URL.Query()
	q.Add("periodType",periodType)
	q.Add("period",period)
	q.Add("frequencyType",frequencyType)
	q.Add("frequency",frequency)
	req.URL.RawQuery = q.Encode()
	body := Handler(req)
	
	var df []FRAME
	var open,hi,lo,Close,volume,datetime string

	split := strings.Split(body,"{")
	split = split[2:len(split)]
	for _,x := range split {
		split2 := strings.Split(x,"\"")
		for i,x2 := range split2 {
			if(x2 == "open") { open = split2[i+1]
			} else if(x2 == "high") { hi = split2[i+1]
			} else if(x2 == "low") { lo = split2[i+1]
			} else if(x2 == "close") { Close = split2[i+1]
			} else if(x2 == "volume") { volume = split2[i+1]
			} else if(x2 == "datetime") { datetime = split2[i+1] }
		}

		f := FRAME{
			DATETIME:	TrimL(TrimFL(datetime)),
			VOLUME:		TrimFL(volume),
			OPEN:		TrimFL(open),
			CLOSE:		TrimFL(Close),
			HI:		TrimFL(hi),
			LO:		TrimFL(lo),
		}

		df = append(df,f)
	}

	return df
}

