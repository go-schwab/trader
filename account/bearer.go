package account

import (
	"fmt"
	"net/http"
	. "github.com/samjtro/go-tda/utils"
)

var endpoint = "https://api.tdameritrade.com/v1/oauth2/token"

// GetBearerToken returns a string; containing the Bearer Token for your account
func GetBearerToken(code,client_id,redirect_url string) {
	
}

