package account

import (
	"fmt"
	"net/http"
	. "github.com/samjtro/go-tda/utils"
)

func main() { Stock() }

var endpoint_place = "https://api.tdameritrade.com/v1/accounts/%s/orders"
var endpoint_replace = "https://api.tdameritrade.com/v1/accounts/%s/orders/%s"
var endpoint_cancel = "https://api.tdameritrade.com/v1/accounts/%s/orders/%s"
var endpoint_get = "https://api.tdameritrade.com/v1/accounts/%s/orders/%s"

func StockOrder(accountID,bearerToken,orderType,session,duration,strategy,orderLeg,instruction,quantity,symbol,assetType string) string {
	url := fmt.Sprintf(endpoint_place,accountID)
	req,_ := http.NewRequest("POST",url,nil)
	req.Authorization = bearerToken

	var str = fmt.Sprintf("
{
  \"orderType\": %s,
  \"session\": %s,
  \"duration\": %s,
  \"orderStrategyType\": %s,
  \"orderLegCollection\": [
    {
      \"instruction\": %s,
      \"quantity\": %s,
      \"instrument\": {
        \"symbol\": %s,
        \"assetType\": %s
      }
    }
  ]
}",orderType,session,duration,strategy,orderLeg,instruction,quantity,instrument,symbol,assetType)

	return str
}

func SingleOption(accountID,bearerToken string) string {
	url := fmt.Sprintf(endpoint_single)
}

func MultipleOption(accountID,bearerToken string) string {

}

func GetOrder(accountID,bearerToken string) string {
	url := fmt.Sprintf(endpoint_get)
	req,_ := http.NewRequest(endpoint_replace,accountID,orderID)
	req.Authorization = bearerToken

}
