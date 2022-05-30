package main

import (
	"fmt"
	"net/http"
	. "github.com/samjtro/go-tda/utils"
)

func main() { 
	Stock("","","","","")
}

var endpoint_place = "https://api.tdameritrade.com/v1/accounts/%s/orders"
var endpoint_replace = "https://api.tdameritrade.com/v1/accounts/%s/orders/%s"
var endpoint_cancel = "https://api.tdameritrade.com/v1/accounts/%s/orders/%s"
var endpoint_get = "https://api.tdameritrade.com/v1/accounts/%s/orders/%s"

func Get(accountID,bearerToken string) string {
	url := fmt.Sprintf(endpoint_get)
	req,_ := http.NewRequest(endpoint_replace,accountID,orderID)
	req.Authorization = bearerToken
}

func Stock(accountID,bearerToken,instruction,ticker,quantity string) string {
	url := fmt.Sprintf(endpoint_place,accountID)
	str := fmt.Sprintf("
{
  \"orderType\": \"MARKET\",
  \"session\": \"NORMAL\",
  \"duration\": \"DAY\",
  \"orderStrategyType\": \"SINGLE\",
  \"orderLegCollection\": [
    {
      \"instruction\": %s,
      \"quantity\": %s,
      \"instrument\": {
        \"symbol\": %s,
        \"assetType\": \"EQUITY\"
      }
    }
  ]
}",instruction,quantity,instrument,ticker)

	req,_ := http.NewRequest("POST",url,bytes.NewBuffer(str))
	req.Header.Set("Authorization",bearerToken)
	req.Header.Set("Content-Type","application/json; charset=UTF-8")
	body := Handler(req)

	return body

}

func ConditionalOrder(accountID,bearerToken string) string {
	url := fmt.Sprintf(endpoint_place,accountID)
	str := fmt.Sprintf("
{
  \"orderType\": \"LIMIT\",
  \"session\": \"NORMAL\",
  \"price\": \"%s\",
  \"duration\": \"DAY\",
  \"orderStrategyType\": \"TRIGGER\",
  \"orderLegCollection\": [
    {
      \"instruction\": \"%s\",
      \"quantity\": %s,
      \"instrument\": {
        \"symbol\": \"%s\",
        \"assetType\": \"EQUITY\"
      }
    }
  ],
  \"childOrderStrategies\": [
    {
      \"orderType\": \"LIMIT\",
      \"session\": \"NORMAL\",
      \"price\": \"%s\",
      \"duration\": \"DAY\",
      \"orderStrategyType\": \"SINGLE\",
      \"orderLegCollection\": [
        {
          \"instruction\": \"%s\",
          \"quantity\": %s,
          \"instrument\": {
            \"symbol\": \"%s\",
            \"assetType\": \"EQUITY\"
          }
        }
      ]
    }
  ]
}",price,instruction,quantity,ticker,price2,instruction2,quantity2,ticker2)

	req,_ := http.NewRequest("POST",url,bytes.NewBuffer(str))
	req.Header.Set("Authorization",bearerToken)
	req.Header.Set("Content-Type","application/json; charset=UTF-8")
	body := Handler(req)

	return body

}

func SingleOption(accountID,bearerToken string) string {
	url := fmt.Sprintf(endpoint_place,accountID)
	str := fmt.Sprintf(" 
{
  \"complexOrderStrategyType\": \"NONE\",
  \"orderType\": \"LIMIT\",
  \"session\": \"NORMAL\",
  \"price\": \"%s\",
  \"duration\": \"DAY\",
  \"orderStrategyType\": \"SINGLE\",
  \"orderLegCollection\": [
    {
      \"instruction\": \"BUY_TO_OPEN\",
      \"quantity\": %s,
      \"instrument\": {
        \"symbol\": \"%s\",
        \"assetType\": \"OPTION\"
    	}
    }
  ]
}")

	req,_ := http.NewRequest("POST",url,bytes.NewBuffer(str))
	req.Header.Set("Authorization",bearerToken)
	req.Header.Set("Content-Type","application/json; charset=UTF-8")
	body := Handler(req)

	return body

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

