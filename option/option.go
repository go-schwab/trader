package option

// Anatomy of the TDA option return:
//{"symbol":"AAPL","status":"SUCCESS","underlying":null,"strategy":"SINGLE","interval":0.0,"isDelayed":true,"isIndex":false,"interestRate":0.1,"underlyingPrice":152.625,"volatility":29.0,"daysToExpiration":0.0,"numberOfContracts":270,"putExpDateMap":{},"callExpDateMap":{"2022-05-13:2":{
//"145.0":[{"putCall":"CALL","symbol":"AAPL_051322C145","description":"AAPL May 13 2022 145 Call (Weekly)","exchangeName":"OPR","bid":9.5,"ask":10.95,"last":9.8,"mark":10.23,"bidSize":50,"askSize":51,"bidAskSize":"50X51","lastSize":0,"highPrice":0.0,"lowPrice":0.0,"openPrice":0.0,"closePrice":9.87,"totalVolume":0,"tradeDate":null,"tradeTimeInLong":1652212791092,"quoteTimeInLong":1652212798993,"netChange":-0.07,"volatility":49.752,"delta":0.903,"gamma":0.022,"theta":-0.181,"vega":0.027,"rho":0.013,"openInterest":907,"timeValue":0.29,"theoreticalOptionValue":9.874,"theoreticalVolatility":29.0,"optionDeliverablesList":null,"strikePrice":145.0,"expirationDate":1652472000000,"daysToExpiration":2,"expirationType":"S","lastTradingDay":1652486400000,"multiplier":100.0,"settlementType":" ","deliverableNote":"","isIndexOption":null,"percentChange":-0.75,"markChange":0.35,"markPercentChange":3.55,"intrinsicValue":9.51,"inTheMoney":true,"mini":false,"pennyPilot":true,"nonStandard":false}],
//"146.0":[{"putCall":"CALL","symbol":"AAPL_051322C146","description":"AAPL May 13 2022 146 Call(Weekly)","exchangeName":"OPR","bid":8.85,"ask":9.75,"last":10.35,"mark":9.3,"bidSize":21,"askSize":50,"bidAskSize":"21X50","lastSize":0,"highPrice":0.0,"lowPrice":0.0,"openPrice":0.0,"closePrice":8.98,"totalVolume":0,"tradeDate":null,"tradeTimeInLong":1652209972509,"quoteTimeInLong":1652212799988,"netChange":1.37,"volatility":49.456,"delta":0.878,"gamma":0.026,"theta":-0.212,"vega":0.031,"rho":0.013,"openInterest":320,"timeValue":1.84,"theoreticalOptionValue":8.985,"theoreticalVolatility":29.0,"optionDeliverablesList":null,"strikePrice":146.0,"expirationDate":1652472000000,"daysToExpiration":2,"expirationType":"S","lastTradingDay":1652486400000,"multiplier":100.0,"settlementType":" ","deliverableNote":"","isIndexOption":null,"percentChange":15.2,"markChange":0.32,"markPercentChange":3.52,"intrinsicValue":8.51,"inTheMoney":true,"mini":false,"pennyPilot":true,"nonStandard":false}],

var endpoint_option string = "https://api.tdameritrade.com/v1/marketdata/chains"

//type UNDERLYING struct {}

type CONTRACT struct {
	TYPE                   string
	SYMBOL                 string
	STRIKE                 string
	EXCHANGE               string
	EXPIRATION             string
	DAYS2EXPIRATION        string
	BID                    string
	ASK                    string
	LAST                   string
	MARK                   string
	BIDASK_SIZE            string
	VOLATILITY             string
	DELTA                  string
	GAMMA                  string
	THETA                  string
	VEGA                   string
	RHO                    string
	OPEN_INTEREST          string
	TIME_VALUE             string
	THEORETICAL_VALUE      string
	THEORETICAL_VOLATILITY string
	PERCENT_CHANGE         string
	MARK_CHANGE            string
	MARK_PERCENT_CHANGE    string
	INTRINSIC_VALUE        string
	IN_THE_MONEY           string //bool
}

// func Vertical() string {}
// func Calendar() string {}
// func Strangle() string {}
// func Straddle() string {}
// func Condor() string {}
// func Diagonal() string {}
// func Collar() string {}
// func Roll() string {}
