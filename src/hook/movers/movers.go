package movers

import (
	"fmt"
	"net/http"
	"github.com/samjtro/go-tda/handler"
)

// movers takes three parameters:
// index = "$DJI", "$SPX.X", or "$COMPX"
// direction = "up" or "down"
// change = "percent" or "value"
func movers(index,direction,change string) string {
	url := fmt.Sprintf(endpoint_movers,index)
	req,_ := http.NewRequest("GET",url,nil)
	q := req.URL.Query()
	q.Add("direction",direction)
	q.Add("change",change)
	req.URL.RawQuery = q.Encode()
	body := handler(req)

	return body
}

