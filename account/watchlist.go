package account

import (
	"fmt"
	"log"

	//"strings"
	"net/http"

	"github.com/samjtro/go-tda/utils"
)

var endpoint_watchlist = "https://api.tdameritrade.com/v1/accounts/%s/watchlists"

// function to create a watchlist
func CreateWatchlist(accountID, Bearer string) string {
	url := fmt.Sprintf(endpoint_watchlist, accountID)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", Bearer)
	body, err := utils.Handler(req)

	if err != nil {
		log.Fatal(err)
	}

	return body
}

// func DeleteWatchlist() {}
// func GetWatchlist() {}
// func ReplaceWatchlist() {}
// func UpdateWatchlist() {}
