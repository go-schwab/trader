package data

import (
	"fmt"
	"time"
	"strings"
	"net/http"
	. "github.com/samjtro/go-tda/utils"
)

// custom struct for use with RealTime
// contains a number of important real time & historical financial indicators
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
	//FiftyTwoWkHI	string
	//FiftyTwoWkLO	string
	PE		string
}

// for use with PriceHistory 
type FRAME struct {
	DATETIME	string
	OPEN		string
	HIGH		string
	LOW		string
	CLOSE		string
	VOLUME		string
}

var endpoint_realtime string = "https://api.tdameritrade.com/v1/marketdata/%s/quotes"// 	 	--> symbol
var endpoint_pricehistory string = "https://api.tdameritrade.com/v1/marketdata/%s/pricehistory"// 	--> symbol

func trimFirstLastChar(s string) string {
	str := s[:len(s)-1]
	return str[1:len(str)]
}

// RealTime returns a QUOTE; containing a real time quote of the desired stock's performance with a number of different indicators (including volatility, volume, price, fundamentals & more), 
// it takes one parameter:
// ticker = "AAPL", etc.
func RealTime(ticker string) QUOTE {
	dt := Now(time.Now())
	url := fmt.Sprintf(endpoint_realtime,ticker)
	req,_ := http.NewRequest("GET",url,nil)
	body := Handler(req)

	var bid,ask,last,open,hi,lo,closeP,mark,volume,volatility,pe string

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
		//} else if(x == "FTWkHigh") { FTwhi = split[i+1]
		//} else if(x == "FTWkLow") { FTwlo = split[i+1]
		} else if(x == "peRatio") { pe = split[i+1]
		}
	}

	bid = trimFirstLastChar(bid)
	ask = trimFirstLastChar(ask)
	last = trimFirstLastChar(last)
	open = trimFirstLastChar(open)
	hi = trimFirstLastChar(hi)
	lo = trimFirstLastChar(lo)
	closeP = trimFirstLastChar(closeP)
	mark = trimFirstLastChar(mark)
	volume = trimFirstLastChar(volume)
	volatility = trimFirstLastChar(volatility)
	//FTwhi = trimFirstLastChar(FTwhi)
	//FTwlo = trimFirstLastChar(FTwlo)
	pe = trimFirstLastChar(pe)	

	return QUOTE{
		DATETIME: 	dt,
		TICKER:   	ticker,
		MARK:		mark,
		VOLUME:		volume,
		VOLATILITY:	volatility,
		BID:	  	bid,
		ASK:	  	ask,
		LAST:     	last,
		OPEN:	  	open,
		CLOSE:		closeP,
		HI:	  	hi,
		LO:	  	lo,
		//FiftyTwoWkHI:	FTwhi,
		//FiftyTwoWkLO:	FTwlo,
		PE:		pe,
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
func PriceHistory(ticker,periodType,period,frequencyType,frequency string) string {
	url := fmt.Sprintf(endpoint_pricehistory,ticker)
	req,_ := http.NewRequest("GET",url,nil)
	q := req.URL.Query()
	q.Add("periodType",periodType)
	q.Add("period",period)
	q.Add("frequencyType",frequencyType)
	q.Add("frequency",frequency)
	req.URL.RawQuery = q.Encode()
	body := Handler(req)
	// var df = []FRAME

	return body
}
