package schwab 

import (
	"fmt"
	"net/http"

	"github.com/samjtro/go-trade/utils"
)

// GetBearerToken returns a string; containing the Bearer Token for your account
func GetBearerToken(accountID string) (string, error) {
	req, _ := http.NewRequest("GET", endpoint_bearer, nil)
	q := req.URL.Query()
	q.Add("grant_type", "authorization_code")
	q.Add("client_id", accountID)
	req.URL.RawQuery = q.Encode()
	resp, err := utils.Handler(req)

	if err != nil {
		return "", err
	}

	body := fmt.Sprintf("Bearer <%s>", resp)

	return body, nil
}
