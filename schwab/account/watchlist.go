package account 

import (
	"fmt"

	//"strings"
	"net/http"

	"github.com/samjtro/go-trade/schwab/utils"
)

// function to create a watchlist
func CreateWatchlist(accountID, Bearer string) (string, error) {
	url := fmt.Sprintf(endpoint_watchlist, accountID)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", Bearer)
	body, err := utils.Handler(req)

	if err != nil {
		return "", err
	}

	return body, nil
}

// func DeleteWatchlist() {}
// func GetWatchlist() {}
// func ReplaceWatchlist() {}
// func UpdateWatchlist() {}
