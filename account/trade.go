package account

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	. "github.com/samjtro/go-tda/utils"
)

var endpoint_place = "https://api.tdameritrade.com/v1/accounts/%s/orders"
var endpoint_replace = "https://api.tdameritrade.com/v1/accounts/%s/orders/%s"
var endpoint_cancel = "https://api.tdameritrade.com/v1/accounts/%s/orders/%s"
var endpoint_get = "https://api.tdameritrade.com/v1/accounts/%s/orders/%s"

// function to Place an order with TD-Ameritrade
func Place(accountID, order string) string {
	bearer, err := GetBearerToken(accountID)

	if err != nil {
		log.Fatal(err)
	}

	req, _ := http.NewRequest("POST", fmt.Sprintf(endpoint_place, accountID), bytes.NewReader([]byte(order)))
	req.Header.Set("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	body, err := Handler(req)

	if err != nil {
		log.Fatal(err)
	}

	return body
}

// func Replace(accountID, bearerToken, orderID string) (string, error) {}
// func Get(accountID, bearerToken, orderID string) (string, error) {}
// func Cancel(accountID, bearerToken, orderID string) (string, error) {}

func Stock(accountID, instruction, ticker, quantity string) string {
	str := fmt.Sprintf(`
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
  }`, instruction, quantity, ticker)

	return Place(accountID, str)
}

func ConditionalOrder(accountID, price, instruction, quantity, ticker, price2, instruction2, quantity2, ticker2 string) string {
	str := fmt.Sprintf(`
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
}`, price, instruction, quantity, ticker, price2, instruction2, quantity2, ticker2)

	return Place(accountID, str)
}

func SingleOption(accountID, price, quantity, ticker string) string {
	str := fmt.Sprintf(`
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
}`, price, quantity, ticker)

	return Place(accountID, str)
}

func VerticalSpread(accountID, bearerToken string) string {
	str := fmt.Sprintf(`
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
}`)

	return str
}

func CustomSpread(accountID, bearerToken string) string {
	str := fmt.Sprintf(`
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
}`)

	return str
}
