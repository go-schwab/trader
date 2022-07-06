package instrument

var endpoint_searchinstrument string = "https://api.tdameritrade.com/v1/instruments"
var endpoint_getinstrument string = "https://api.tdameritrade.com/v1/instruments/%s" // cusip

// desc-regex: Search description with full regex support. Example: symbol=XYZ.[A-C] returns all instruments whose descriptions contain a word beginning with XYZ followed by a character A through C.
