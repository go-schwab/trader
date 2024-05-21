package schwab 

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/samjtro/go-tda/utils"
)

//var endpoint_replace = "https://api.tdameritrade.com/v1/accounts/%s/orders/%s"
//var endpoint_cancel = "https://api.tdameritrade.com/v1/accounts/%s/orders/%s"
//var endpoint_get = "https://api.tdameritrade.com/v1/accounts/%s/orders/%s"

// function to Place an order with TD-Ameritrade
func Place(accountID, order string) (string, error) {
	bearer, err := GetBearerToken(accountID)

	if err != nil {
		return "", err
	}

	req, _ := http.NewRequest("POST", fmt.Sprintf(endpoint_place, accountID), bytes.NewReader([]byte(order)))
	req.Header.Set("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	body, err := utils.Handler(req)

	if err != nil {
		return "", err
	}

	return body, nil
}

// func Replace(accountID, bearerToken, orderID string) (string, error) {}
// func Get(accountID, bearerToken, orderID string) (string, error) {}
// func Cancel(accountID, bearerToken, orderID string) (string, error) {}

func Stock(accountID, instruction, ticker, quantity string) (string, error) {
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

	placeOrder, err := Place(accountID, str)

	if err != nil {
		return "", err
	}

	return placeOrder, nil
}

func ConditionalOrder(accountID, price, instruction, quantity, ticker, price2, instruction2, quantity2, ticker2 string) (string, error) {
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

	placeOrder, err := Place(accountID, str)

	if err != nil {
		return "", err
	}

	return placeOrder, nil
}

func SingleOption(accountID, price, quantity, ticker string) (string, error) {
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

	placeOrder, err := Place(accountID, str)

	if err != nil {
		return "", err
	}

	return placeOrder, nil
}

/* WIP
func VerticalSpread(accountID, bearerToken string) (string, error) {
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

        placeOrder, err := Place(accountID, str)

        if err != nil {
        return "", err
        }

        return placeOrder, nil
}

func CustomSpread(accountID, bearerToken string) (string, error) {
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

        placeOrder, err := Place(accountID, str)

        if err != nil {
        return "", err
        }

        return placeOrder, nil
}*/
