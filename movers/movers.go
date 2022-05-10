package movers

import (
	"fmt"
	"net/http"
	. "github.com/samjtro/go-tda/utils"
)
 
var endpoint_movers string = "https://api.tdameritrade.com/v1/marketdata/%s/movers"// 	 		--> index

// Get returns a string; containing information on the desired index's movers per your desired direction and change type(percent or value), 
// it takes three parameters:
// index = "$DJI", "$SPX.X", or "$COMPX"
// direction = "up" or "down"
// change = "percent" or "value"
func Get(index,direction,change string) string {
	url := fmt.Sprintf(endpoint_movers,index)
	req,_ := http.NewRequest("GET",url,nil)
	q := req.URL.Query()
	q.Add("direction",direction)
	q.Add("change",change)
	req.URL.RawQuery = q.Encode()
	body := Handler(req)

	return body
}

