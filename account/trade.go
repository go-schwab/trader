package account

import (
	"fmt"
	"net/http"
	. "github.com/samjtro/go-tda/utils"
)

var endpoint_place = "https://api.tdameritrade.com/v1/accounts/%s/orders"
var endpoint_replace = "https://api.tdameritrade.com/v1/accounts/%s/orders/%s"
var endpoint_cancel = "https://api.tdameritrade.com/v1/accounts/%s/orders/%s"
var endpoint_get = "https://api.tdameritrade.com/v1/accounts/%s/orders/%s"

func GetOrder(accountID,bearerToken string) string {
	url := fmt.Sprintf(endpoint_get)
	req,_ := http.NewRequest(endpoint_replace,accountID,orderID)
	req.Authorization = bearerToken
}

func StockOrder(accountID,bearerToken,orderType,session,duration,strategy,orderLeg,instruction,quantity,symbol,assetType string) string {
	url := fmt.Sprintf(endpoint_place,accountID)
	req,_ := http.NewRequest("POST",url,nil)
	req.Authorization = bearerToken

	str := fmt.Sprintf("
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

func ConditionalOrder(accountID,bearerToken string) string {
	url := fmt.Sprintf(endpoint_place,accountID)
	req,_ := http.NewRequest("POST",url,nil)
	req.Authorization = bearerToken

	str := fmt.Sprintf("
{
  \"orderType\": \"LIMIT\",
  \"session\": \"NORMAL\",
  \"price\": \"34.97\",
  \"duration\": \"DAY\",
  \"orderStrategyType\": \"TRIGGER\",
  \"orderLegCollection\": [
    {
      \"instruction\": \"BUY\",
      \"quantity\": 10,
      \"instrument\": {
        \"symbol\": \"XYZ\",
        \"assetType\": \"EQUITY\"
      }
    }
  ],
  \"childOrderStrategies\": [
    {
      \"orderType": "LIMIT\",
      \"session": "NORMAL\",
      \"price": "42.03\",
      \"duration": "DAY\",
      \"orderStrategyType\": \"SINGLE\",
      \"orderLegCollection\": [
        {
          \"instruction\": \"SELL\",
          \"quantity\": 10,
          \"instrument\": {
            \"symbol\": \"XYZ\",
            \"assetType\": \"EQUITY\"
          }
        }
      ]
    }
  ]
}")
}

func SingleOption(accountID,bearerToken string) string {
	url := fmt.Sprintf(endpoint_single)
	req,_ := http.NewRequest("POST",url,nil)
	req.Authorization = bearerToken


	str := fmt.Sprintf(" 
{
  \"complexOrderStrategyType\": \"NONE\",
  \"orderType\": \"LIMIT\",
  \"session\": \"NORMAL\",
  \"price\": \"6.45\",
  \"duration\": \"DAY\",
  \"orderStrategyType\": \"SINGLE\",
  \"orderLegCollection\": [
    {
      \"instruction\": \"BUY_TO_OPEN\",
      \"quantity\": 10,
      \"instrument\": {
        \"symbol\": \"XYZ_032015C49\",
        \"assetType\": \"OPTION\"
    	}
    }
  ]
}")
}

func VerticalSpread(accountID,bearerToken string) string {
	str := fmt.Sprintf("
{
  \"orderType\": \"NET_DEBIT\",
  \"session\": \"NORMAL\",
  \"price\": \"1.20\",
  \"duration\": \"DAY\",
  \"orderStrategyType\": \"SINGLE\",
  \"orderLegCollection\": [
    {
      \"instruction\": \"BUY_TO_OPEN\",
      \"quantity\": 10,
      \"instrument\": {
        \"symbol\": \"XYZ_011516C40\",
        \"assetType\": \"OPTION\"
      }
    },
    {
      \"instruction\": \"SELL_TO_OPEN\",
      \"quantity\": 10,
      \"instrument\": {
        \"symbol\": \"XYZ_011516C42.5\",
        \"assetType\": \"OPTION\"
      }
    }
  ]
}")
}

func CustomSpread(accountID,bearerToken string) string {
	str := fmt.Sprintf("
{
 \"orderStrategyType\": \"SINGLE\",
  \"orderType\": \"MARKET\",
  \"orderLegCollection\": [
    {
      \"instrument\": {
        \"assetType\": \"OPTION\",
        \"symbol\": \"XYZ_011819P45\"
    },
      \"instruction\": \"SELL_TO_OPEN\",
      \"quantity\": 1
    },
    {
      \"instrument\": {
        \"assetType\": \"OPTION\",
        \"symbol\": \"XYZ_011720P43\"
      },
      \"instruction\": \"BUY_TO_OPEN\",
      \"quantity\": 2
    }
  ],
  \"complexOrderStrategyType\": \"CUSTOM\",
  \"duration\": \"DAY\",
  \"session\": \"NORMAL\"
}")
}

