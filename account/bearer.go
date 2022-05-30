package main

import (
	"fmt"
	"net/http"
	. "github.com/samjtro/go-tda/utils"
)

func main() { fmt.Println(GetBearerToken("","")) }

var endpoint = "https://api.tdameritrade.com/v1/oauth2/token"

// GetBearerToken returns a string; containing the Bearer Token for your account
func GetBearerToken(client_id,redirect_url string) string {
	req,_ := http.NewRequest("GET",endpoint,nil)
	q := req.URL.Query()
	q.Add("grant_type","authorization_code")
	q.Add("client_id",client_id)
	req.URL.RawQuery = q.Encode()
	resp := Handler(req)
	body := fmt.Sprintf("Bearer <%s>",resp)

	return body
}

