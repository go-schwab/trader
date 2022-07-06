package account

import (
	"fmt"
	//"strings"
	"net/http"

	. "github.com/samjtro/go-tda/utils"
)

var endpoint_watchlist = "https://api.tdameritrade.com/v1/accounts/%s/watchlists"

func CreateWatchlist(accountID, Bearer string) (string, error) {
	url := fmt.Sprintf(endpoint_watchlist, accountID)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", Bearer)
	body, err := Handler(req)

	if err != nil {
		return "", err
	}

	return body, nil
}

// func DeleteWatchlist() {}
// func GetWatchlist() {}
// func ReplaceWatchlist() {}
// func UpdateWatchlist() {}
